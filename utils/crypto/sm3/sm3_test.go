/*
 * Package sm3 implements the Chinese SM3 Digest Algorithm,
 * according to "go/src/crypto/sha256"
 * author: weizhang <d5c5ceb0@gmail.com>
 * fork by: borisding
 */

package sm3

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

type sm3Test struct {
	out string
	in  string
}

var golden = []sm3Test{
	{"1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b", ""},
	{"66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0", "abc"},
	{"8367760325bd95ebb27d1259e721d12714591ce89fd5a22853e70499353090c8", "天地玄黃 宇宙洪荒 日月盈昃 辰宿列張 てんち　げんこう"},
	{"12da74524112ed95c72f78ea02b4e10d5b7b3b2ead9225d93f0cfec9b3e2ad97", `{"name":"example","meta":{"id":1},"tags":["GML","XML"]}`},
	{"2bb6c53ad20eaf2552425f44e72d96d1b61e63310a1a30f4e5406a103619177d", "He who has a shady past knows that nice guys finish last."},
}

// nolint: errcheck
func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		s := fmt.Sprintf("%x", Sum([]byte(g.in)))
		if s != g.out {
			t.Fatalf("Sum function: sm3(%s) = %s want %s", g.in, s, g.out)
		}
		c := New()
		for j := 0; j < 3; j++ {
			if j < 2 {
				io.WriteString(c, g.in)
			} else {
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.Sum(nil)
				io.WriteString(c, g.in[len(g.in)/2:])
			}
			s := fmt.Sprintf("%x", c.Sum(nil))
			if s != g.out {
				t.Fatalf("sm3[%d](%s) = %s want %s", j, g.in, s, g.out)
			}
			c.Reset()
		}
	}
}

func TestSize(t *testing.T) {
	c := New()
	if got := c.Size(); got != Size {
		t.Errorf("Size = %d; want %d", got, Size)
	}
}

func TestBlockSize(t *testing.T) {
	c := New()
	if got := c.BlockSize(); got != BlockSize {
		t.Errorf("BlockSize = %d want %d", got, BlockSize)
	}
}

// Tests that blockGeneric (pure Go) and block (in assembly for some architectures) match.
// nolint: errcheck
func TestBlockGeneric(t *testing.T) {
	gen, asm := New().(*digest), New().(*digest)
	buf := make([]byte, BlockSize*20) // arbitrary factor
	rand.Read(buf)
	blockGeneric(gen, buf)
	block(asm, buf)
	if *gen != *asm {
		t.Error("block and blockGeneric resulted in different states")
	}
}

var bench = New()
var buf = make([]byte, 8192)

// nolint: errcheck
func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192)
}

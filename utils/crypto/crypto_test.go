package crypto

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSM3Hash(t *testing.T) {
	msg := "处江泽之远则忧其民"
	msgHash := "759edcbf9211a630d1968e77a907a3b7585bc5bfd88a43e7d4d99cb9f03b6c62"
	if hash := SM3Hash(msg); hash != msgHash {
		t.Fatalf("SM3Hash(\"%v\") expect is %v , but got %v", msg, msgHash, hash)
	}
}

func TestGenerateKeypair(t *testing.T) {
	kp := GenerateKeypair()
	if kp.PublicKey == nil || kp.PrivateKey == nil {
		t.Fatalf("GenerateKeypair failure, PubKey is %v, PrivKey is %v", kp.PublicKey, kp.PrivateKey)
	}
}

func TestBase58Encode(t *testing.T) {
	buf := Base58Encode(nil, big.NewInt(1234567890))
	if !reflect.DeepEqual(buf, []byte("aSozam")) {
		t.Fatalf("Base58Encode failure, expect is aSozam, but got %v", buf)
	}
}

func TestBase58Decode(t *testing.T) {
	n, err := Base58Decode([]byte("aSozam"))
	if err != nil || n.String() != "1237104346" {
		t.Fatalf("Base58Decode failure, expect is 1237104346, but got %v. or error is %v", n.String(), err)
		return
	}
}

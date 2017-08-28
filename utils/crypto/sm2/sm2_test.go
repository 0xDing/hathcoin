package sm2

import (
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/borisding1994/hathcoin/utils/crypto/sm3"
)

func testKeyGeneration(t *testing.T, c elliptic.Curve, tag string) {
	priv, err := GenerateKey(c, rand.Reader)
	if err != nil {
		t.Errorf("%s: error: %s", tag, err)
		return
	}
	if !c.IsOnCurve(priv.PublicKey.X, priv.PublicKey.Y) {
		t.Errorf("%s: public key invalid: %s", tag, err)
	}
}

func TestKeyGeneration(t *testing.T) {
	testKeyGeneration(t, P256Sm2(), "sm2 p256")
	if testing.Short() {
		return
	}
}

func testSignAndVerify(t *testing.T, c elliptic.Curve, tag string) {
	priv, _ := GenerateKey(c, rand.Reader)

	msg := []byte("testing")
	dig := sm3.Sum(msg)
	hashed := dig[:]
	r, s, err := Sign(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	if !Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("%s: Verify failed", tag)
	}

	msg[0] ^= 0xff
	dig = sm3.Sum(msg)
	hashed = dig[:]
	if Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("%s: Verify always works!", tag)
	}
}

func TestSignAndVerify(t *testing.T) {
	testSignAndVerify(t, P256Sm2(), "sm2 p256")
	if testing.Short() {
		return
	}

	for i := 0; i < 20; i++ {
		testSignAndVerify(t, P256Sm2(), "sm2 p256")
	}
}

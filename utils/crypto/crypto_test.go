package crypto

import "testing"

func TestSM3Hash(t *testing.T) {
	msg := "处江泽之远则忧其民"
	msgHash := "759edcbf9211a630d1968e77a907a3b7585bc5bfd88a43e7d4d99cb9f03b6c62"
	if hash := SM3Hash(msg); hash != msgHash {
		t.Fatalf("SM3Hash(\"%v\") expect is %v , but got %v", msg, msgHash, hash)
	}
}

package baseha

import (
	"bytes"
	"testing"
)

func TestEncodingAndDecoding(t *testing.T) {
	src := []byte("千金散尽还复来，叫你闷声发大财 赵钱孙李一起膜蛤1234!!@@##Helloasc..》10231 werk 百壶且试开怀抱，熟悉西方那一套")
	t.Logf("Original:\t%v", src)

	encodedData := STD.EncodeToString(src)
	t.Logf("Encoded:\t%s", encodedData)

	encodedLen := STD.EncodedLen(len(src))
	decodedLen := STD.DecodedLen(encodedData)
	if encodedLen != decodedLen {
		t.Errorf("No Mached Length. encoded len: %d, decoded len: %d", encodedLen, decodedLen)
		t.FailNow()
	}

	decodedData := STD.Decode(encodedData)
	t.Logf("Decoded:\t%v", decodedData)

	if !bytes.Equal(src, decodedData) {
		t.Errorf("No Matched Original Data.")
		t.FailNow()
	}
}

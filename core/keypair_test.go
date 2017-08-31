package core

import (
	"os"
	"testing"

	"github.com/borisding1994/hathcoin/config"
)

func TestLoadKeypair(t *testing.T) {
	dir := config.GetString("data_dir")
	os.MkdirAll(dir, 0777)
	LoadKeypair()
	t.Logf(string(currentPeer.Keypair.PublicKey))
	os.RemoveAll(dir)
}

package core

import (
	"encoding/gob"
	"errors"
	"os"
	"path"

	"github.com/borisding1994/hathcoin/config"
	"github.com/borisding1994/hathcoin/utils"
	"github.com/borisding1994/hathcoin/utils/crypto"
)

const (
	FileName   = "keypair.dat"
	OpenErrMsg = "Failed to open or create keypair data file. "
)

func LoadKeypair() {
	dir := config.GetString("data_dir")
	keypair, _ := openKeypair(dir)
	if keypair == nil {
		utils.Logger.Info("Generating keypair...")
		keypair = crypto.GenerateKeypair()
		err := saveKeypair(dir, keypair)
		if err != nil {
			utils.Logger.Fatal("Failed to save keypair data file. ", err)
		}
	}
	Peer.Keypair = keypair
}

func openKeypair(dir string) (*crypto.Keypair, error) {
	file, err := os.OpenFile(path.Join(dir, FileName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.Logger.Fatal(OpenErrMsg, err)
	}
	defer file.Close()

	k := &crypto.Keypair{}
	gob.NewDecoder(file).Decode(k)

	if k == nil || k.PublicKey == nil || k.PrivateKey == nil {
		return nil, nil
	}

	return k, file.Close()
}

func saveKeypair(dir string, keypair *crypto.Keypair) error {
	if keypair != nil {
		file, err := os.OpenFile(path.Join(dir, FileName), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			utils.Logger.Fatal(OpenErrMsg, err)
		}
		err = gob.NewEncoder(file).Encode(keypair)
		if err != nil {
			utils.Logger.Fatal("Failed to encode raw keypair. ", err)
		}
		return file.Close()

	}
	return errors.New("No keypair provided to save")
}

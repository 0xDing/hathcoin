package utils

import (
	"os"

	"github.com/borisding1994/hathcoin/config"
	"github.com/sirupsen/logrus"
)

// Logger is new logrus logger instantiate.
var Logger = logrus.New()

func init() {
	Logger.Formatter = &logrus.JSONFormatter{}
	if config.GetString("log_driver") == "file" {
		// nolint: gas
		file, err := os.OpenFile(config.GetString("log_file"), os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			Logger.Out = file
		} else {
			Logger.Warn("Failed to log to file, using default stderr. ", err)
		}
	} else {
		Logger.Out = os.Stdout
	}
}

package test

import (
	"github.com/dipper-mizar/eyes"
	"testing"
)

func TestFile(t *testing.T) {
	logger := eyes.GetLogger()
	logger.SetConfig(eyes.INFO, true, "Asia/Shanghai",
		eyes.WithFile("class3", "./"),
		)
	logger.Info("name")
	logger.Debug("name", "abc")
	logger.Error("name", "1111")
}
package global

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	var err error
	CmdDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panicf("get cmd dir failed: %v\n", err)
	}

	SrcDir = filepath.Dir(CmdDir)
	ProjectDir = filepath.Dir(SrcDir)
	ConfDir = filepath.Join(ProjectDir, "conf")

	AppConfigPath = filepath.Join(ConfDir, AppConfigFile)
}

// dirs
var (
	CmdDir     string
	SrcDir     string
	ProjectDir string
	ConfDir    string
)

// file path
var (
	AppConfigPath string
)

const (
	AppConfigFile string = "app_config.yaml"
)

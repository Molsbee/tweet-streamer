package config

import (
	"os"
	"path"
	"runtime"
)

func FilePath() string {
	return path.Join(getUserHome(), ".twitter-config")
}

func getUserHome() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}

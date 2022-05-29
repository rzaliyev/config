package config

import (
	"io/ioutil"
	"os"
	"path"
)

// Get will open or create if doesn't exist a config file
// in a specified subdirectory of user home directory
// and return a content
func Get(dirName, fileName string) (data []byte, err error) {
	homeDirName, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := path.Join(homeDirName, dirName)
	configFile := path.Join(configPath, fileName)

	err = os.Mkdir(configPath, 0750)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

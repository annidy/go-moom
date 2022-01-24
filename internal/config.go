package internal

import (
	"fmt"
	"github.com/groob/plist"
	"os"
	"os/user"
	"path/filepath"
)
type Config struct {
	Count   int64   `plist:"c"`
}

var plistPath string

func init() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	plistPath = filepath.Join(dir, "Library/Preferences/.byhost.spotlight.mdquery")
}

func (config *Config)ReadConfig() (error)  {
	// Open our xmlFile
	xmlFile, err := os.Open(plistPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	err = plist.NewXMLDecoder(xmlFile).Decode(&config)
	return nil
}

func (config *Config)WriteConfig() error {

	xmlFile, err := os.Create(plistPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	encoder := plist.NewEncoder(xmlFile)
	encoder.Indent("  ")
	encoder.Encode(config)

	return nil
}

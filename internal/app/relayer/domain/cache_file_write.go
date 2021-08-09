package domain

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

var (
	userDir, _ = os.UserHomeDir()
	homeDir    = filepath.Join(userDir, ".tibc-relayer")
)

type CacheFileWriter struct {
	cacheDir      string
	cacheFilename string
}

func NewCacheFileWriter(cacheDir string, cacheFilename string) *CacheFileWriter {
	return &CacheFileWriter{cacheDir: cacheDir, cacheFilename: cacheFilename}
}

type cacheData struct {
	LatestHeight uint64 `json:"latest_height"`
}

func (w *CacheFileWriter) Write(height uint64) error {

	cacheDataObj := &cacheData{}
	cacheDataObj.LatestHeight = height

	cacheDataWriteBytes, err := json.Marshal(cacheDataObj)
	if err != nil {
		return err
	}

	cacheDir := path.Join(homeDir, w.cacheDir)
	filename := path.Join(cacheDir, w.cacheFilename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// And the home folder doesn't exist
		if _, err := os.Stat(homeDir); os.IsNotExist(err) {
			// Create the home folder
			if err = os.Mkdir(homeDir, os.ModePerm); err != nil {
				return err
			}
		}
		// Create the home config folder
		if err = os.Mkdir(cacheDir, os.ModePerm); err != nil {
			return err
		}
		// Then create the file...
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}

	} else {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}
	}
	return nil
}

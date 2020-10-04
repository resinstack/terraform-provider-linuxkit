package linuxkit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func defaultMobyConfigDir() (string, error) {
	mobyDefaultDir := ".moby"
	home, err := homedir.Dir()
	return filepath.Join(home, mobyDefaultDir), err
}

func id(input interface{}) string {
	b, _ := json.Marshal(input)
	return hash(string(b))
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

func stringPtr(s string) *string {
	return &s
}

func copyFile(src, dst string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		return err
	}

	return nil
}

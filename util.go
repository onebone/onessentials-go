package onessentials

import (
	"os"
	"io"
	"path"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)


func CopyFileIfNotExists(path, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		return CopyFile(path, dest)
	}

	return nil
}

func CopyFile(path, dest string) error {
	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	e := out.Close()
	if e != nil {
		return e
	}

	return nil
}

func InitConfig(config interface{}) (err error) {
	var base string
	base, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	if _, err = os.Stat(path.Join(base, "config.json")); os.IsNotExist(err) {
		CopyFile(path.Join(base, "resources", "config.json"), path.Join(base, "config.json"))
	}

	b, err := ioutil.ReadFile(path.Join(base, "config.json"))
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &config)
	return
}

package pipeline

import (
	"os"
	"path"
	"encoding/json"
	"io/ioutil"
)

type FileStorage struct {
	Prefix string 	// 存储路径
}

func (s FileStorage) Get(subpath, id string) ([]byte, error) {
	filename := path.Join(s.Prefix, subpath, id)
	return ioutil.ReadFile(filename)
}

func (s *FileStorage) Put(subpath string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	filename := path.Join(s.Prefix, subpath)
	if err := os.MkdirAll(path.Dir(filename), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(filename , b, 0644)
}

var store *FileStorage

func InitStore(prefix string) {
	store = &FileStorage{
		Prefix: prefix,
	}
}
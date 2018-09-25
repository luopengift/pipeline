package pipeline

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// FileStorage FileStorage
type FileStorage struct {
	Prefix string // 存储路径
}

// Get get
func (s FileStorage) Get(subpath, id string) ([]byte, error) {
	filename := path.Join(s.Prefix, subpath, id)
	return ioutil.ReadFile(filename)
}

//Put put
func (s *FileStorage) Put(subpath string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	filename := path.Join(s.Prefix, subpath)
	if err := os.MkdirAll(path.Dir(filename), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0644)
}

var store *FileStorage

// InitStore init store
func InitStore(prefix string) {
	store = &FileStorage{
		Prefix: prefix,
	}
}

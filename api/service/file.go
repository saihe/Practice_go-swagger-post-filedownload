package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type FileService interface {
	Generate(id string, name string) (string, []byte)
}

type FileServiceImpl struct{}

func NewFileService() FileService {
	return &FileServiceImpl{}
}

func (s *FileServiceImpl) Generate(id string, name string) (string, []byte) {
	fn := fmt.Sprintf("%s_%s.csv", name, id)
	f, err := ioutil.TempFile("", fn)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	len, err := f.WriteString("col1, col2, col3\n")
	if len == 0 {
		log.Fatal("ファイル書き込みできてない")
	}
	if err != nil {
		log.Fatal(err)
	}
	ret, err := f.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d", ret)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return fn, b
}

package walkdir_service

import (
	"io/ioutil"
	"path/filepath"
	"github.com/JiHanHuang/gin_vue/pkg/logging"
)

const (
	TFile = iota
	TDir
)


type Files struct {
	RelativePath string
	Name string
	T int
}

func WalkDir(dir string) (fs []Files){
	files, err := ioutil.ReadDir(dir)
	if err != nil{
		logging.Error(err)
		return nil
	}
	for _, file := range files{
		f := Files{}
		f.Name = file.Name()
		f.T = TFile
		if file.IsDir(){
			f.T = TDir
		}
		f.RelativePath, _ = filepath.Abs(f.Name)
		fs = append(fs, f)
	}
	return
}
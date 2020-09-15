package download_service

import (
	"fmt"
	"sort"
	"sync"

	"github.com/JiHanHuang/gin_vue/pkg/download"
	"github.com/JiHanHuang/gin_vue/service/download_service/ftp"
)

type Download interface {
	DownloadFile() error
	GetStatus() download.Status
	GetAttr() download.Attr
}

var downList sync.Map

func InitDownload(id int, addr, path, fileName string) Download {
	d := ftp.InitFtp(id, addr, path, fileName)
	downList.Store(id, d)
	return d
}

func GetDownList() (statutsList []Download) {
	var ids []int
	downList.Range(func(key, value interface{}) bool {
		ids = append(ids, key.(int))
		return true
	})
	sort.Ints(ids)
	for _, id := range ids {
		d, _ := downList.Load(id)
		statutsList = append(statutsList, d.(Download))
	}
	return
}

func GetDown(id int) (Download, error) {
	if d, ok := downList.Load(id); ok {
		entry, _ := d.(Download)
		return entry, nil
	}
	return nil, fmt.Errorf("Not find data. ID:%d.", id)
}

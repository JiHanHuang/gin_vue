package download_service

import (
	"fmt"
	"sync"

	"github.com/JiHanHuang/gin_vue/pkg/logging"

	"github.com/JiHanHuang/gin_vue/pkg/download"
	"github.com/JiHanHuang/gin_vue/service/download_service/ftp"
)

type Download interface {
	DownloadFile() error
	GetStatus() download.Status
	GetID() int
	GetName() string
}

var downList sync.Map

func InitDownload(id int, addr, path, fileName string) Download {
	d := ftp.InitFtp(id, addr, path, fileName)
	downList.Store(id, d)
	return d
}

func GetStatus(id int) (error, download.Status) {
	if d, ok := downList.Load(id); ok {
		entry, ret := d.(Download)
		if !ret {
			return fmt.Errorf("Type assert failed in GetStatus. ID:%d.", id), download.Status{}
		}
		return nil, entry.GetStatus()
	}
	return fmt.Errorf("Not find data. ID:%d.", id), download.Status{}
}

func GetDownList() (statutsList []Download) {
	downList.Range(func(key, value interface{}) bool {
		entry, ret := value.(Download)
		if !ret {
			logging.Error("Type assert failed in download list.")
			return false
		}
		statutsList = append(statutsList, entry)
		return true
	})
	return
}

package download_service

import (
	"io"
	"net/http"
	"os"

	"github.com/JiHanHuang/gin_vue/pkg/logging"
)

type FTP struct {
	Addr     string
	DownPath string
}

func (f *FTP) Download() error {
	logging.Info("FTP download....")
	logging.Info(f.Addr)

	resp, err := http.Get(f.Addr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(f.DownPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

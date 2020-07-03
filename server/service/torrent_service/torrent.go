package torrent_service

import (
	"fmt"
	"log"

	"github.com/JiHanHuang/gin_vue/service/torrent_service/torrentfile"
)

type Torrent struct {
	ID       int
	Name     string
	Addr     string
	BTFile   string
	DownPath string
}

func (t *Torrent) Download() error {
	fmt.Println("Torrent download...")
	tf, err := torrentfile.Open(t.BTFile)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(t.DownPath)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

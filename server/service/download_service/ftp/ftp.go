package ftp

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/JiHanHuang/gin_vue/pkg/download"
	"github.com/JiHanHuang/gin_vue/pkg/logging"
)

type FTP struct {
	ID       int
	Addr     string
	DownPath string
	FileName string
	fileSize int64
	s        *download.Status
}

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	Total uint64
	F     *FTP
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	//fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	//fmt.Printf("\rDownloading... %d B complete", wc.Total)

	wc.F.s.UpP(int(int64(wc.Total*100) / wc.F.fileSize))
}

func InitFtp(id int, addr, path, fileName string) *FTP {
	return &FTP{
		ID:       id,
		Addr:     addr,
		DownPath: path,
		FileName: fileName,
		s:        download.InitStatus(download.Waiting, 0),
	}
}

func (f *FTP) DownloadFile() error {
	logging.Debug("FTP download....")
	logging.Debug(fmt.Sprintln(f))
	go downloadWorker(f)
	return nil
}

func downloadWorker(f *FTP) {
	f.s.UpS(download.Start)
	resp, err := http.Get(f.Addr)
	if err != nil {
		f.s.UpS(download.Failed)
		return
	}
	defer resp.Body.Close()

	f.fileSize = resp.ContentLength

	out, err := os.Create(f.DownPath + "/" + f.FileName)
	if err != nil {
		f.s.UpS(download.Failed)
		return
	}
	defer out.Close()
	f.s.UpS(download.Running)
	counter := &WriteCounter{F: f}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		f.s.UpS(download.Failed)
		return
	}
	f.s.UpS(download.Finish)
}

func (f *FTP) GetStatus() download.Status {
	return *(f.s)
}

func (f *FTP) GetID() int {
	return f.ID
}

func (f *FTP) GetName() string {
	return f.FileName
}

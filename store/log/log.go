package log

import (
	"io"
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func init() {
	Logger = log.New(InitWriter(), "", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Println("Read Config……")
}
func InitWriter() io.Writer {
	//if !PathExists("./log/") {
		os.Mkdir("./log/", 0773)
	//}

	path := "./log/" + time.Now().Format("2006-01-02") + ".log"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0773)
	if err != nil {
		log.Fatal(err)
	}
	//defer f.Close()
	writers := []io.Writer{
		f,
		os.Stdout}
	return io.MultiWriter(writers...)
}

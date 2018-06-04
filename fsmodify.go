package fsmodify

import (
	"io/ioutil"
	"time"
	"path/filepath"
)

var LastModify time.Time
var timeInterval time.Duration
type fnHandler func(string)

func NewWatcher(pathToWatch string, extension string, interval int, fn fnHandler) {
	LastModify = time.Now()
	timeInterval = time.Duration(time.Duration(interval) * time.Second)

	for {
		err := watchFiles(pathToWatch, extension, fn)
		if err != nil {
			break;
		}
		time.Sleep(timeInterval)
		LastModify = time.Now()
	}
}

//func main () {
//	fmt.Println("Watching files")
//	NewWatcher("./teste",".txt", 3, func(f string) {
//		fmt.Println("Arquivo alterado->", f)
//	})
//}

// Example : LoadListFiles("/Users/test", ".html")
func watchFiles(path string, ext string, fn fnHandler) (error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		fileExt := filepath.Ext(f.Name())
		if fileExt == ext {
			diff := LastModify.Sub(f.ModTime())
			if diff < timeInterval {
				//fmt.Printf("arquivo alterado : %v \n", f.Name())
				fn(f.Name())
			}
		}
	}

	return nil
}
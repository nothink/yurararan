package shelf

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
)


// TODO 外部化
const (
	// rootPath = "/Users/kaba/Dropbox/Wasabi/verenav/"
	rootPath = "/verenav/"
)

// TODO Shelf interface 化すること

var allFiles mapset.Set = mapset.NewSet()
var allMtx *sync.Mutex = new(sync.Mutex)

func Init() {
	allFiles = getAllFileSet()
	go watch()
}

func All() []interface{} {
	return allFiles.ToSlice()
}

func Append(s []interface{}) []interface{} {
	if s == nil {
		return nil
	}

	var diff mapset.Set
	if allFiles != nil {
		diff = mapset.NewSetFromSlice(s).Difference(allFiles)
	} else {
		diff = mapset.NewSetFromSlice(s)
	}

	if c := diff.Cardinality(); c == 0 {
		return nil
	}
	// update(allFiles.Union(s))
	for _, path := range diff.ToSlice() {
		go fetch(path.(string))
	}
	return diff.ToSlice()
}

func watch() {
	for _ = range time.Tick(30 * time.Minute) {
		// 30分おきにファイルシステムのもので上書きする
		tmpAll := getAllFileSet()
		allMtx.Lock()
		allFiles = tmpAll
		allMtx.Unlock()
	}
}

func fetch(path string) {
	if strings.Contains(path, "stat100.ameba.jp") {
		res, err := http.Get(fmt.Sprintf("https://%v", path))
		if err != nil {
			log.Print("Failed: ", path, " - ", err)
			return
		}
		defer res.Body.Close()

		fullPath := filepath.Join(rootPath, path)

		if _, err := os.Stat(filepath.Dir(fullPath)); os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(fullPath), 0755)
		}

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			f, err := os.Create(fullPath)
		    if err != nil {
		        log.Print(err)
				return
		    }
		    defer f.Close()
			_, err = io.Copy(f, res.Body)
			if err != nil {
				log.Print(err)
				return
			}
			allFiles.Add(path)
		}
	}
}

func getAllFileSet() mapset.Set {
	result := mapset.NewSet()

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error{
		if err != nil {
			return err
		}
		if !info.IsDir() {
			keyPath := path[len(rootPath):]
			result.Add(keyPath)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return result
}

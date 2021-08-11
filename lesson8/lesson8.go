package lesson8

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type programFlag struct {
	DeleteDuplicate bool
	YesDelete bool
}

var Flag programFlag
type FileInfo struct {
	Name string
	Size int64
}
var files = make(map[FileInfo]string)

func checkDuplicate(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	f := FileInfo{info.Name(),info.Size()}
	if v, ok := files[f]; ok {
		fmt.Printf("%q is a duplicate of %q\n", path, v)
	} else {
		files[f] = path
	}

	return nil
}

//FindDuplicateFile - функция ищущая дубликаты файлов

func FindDuplicateFile(path string)   {
	fmt.Println(path)
	err := filepath.Walk(path, checkDuplicate)
	if err != nil {
		fmt.Println(err)
	}
}

func Start() {
	flag.BoolVar(&Flag.DeleteDuplicate,"d", false, "Delete duplicate files")

	flag.BoolVar(&Flag.YesDelete,"y", false, "Delete duplicate files")
	flag.Parse()

	//fmt.Println(Flag)
	var pathname string
	if pathname = flag.Arg(0); pathname != "" {
		// проверяем, что это дирректория и она существует
		info, err := os.Stat(pathname)
		if os.IsNotExist(err) {
			fmt.Println("path does not exist.")
			return
		}
		if !info.IsDir() {
			fmt.Println("path is not a directory")
			return
		}
	} else {
		// получаем текущий каталог
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		pathname = filepath.Dir(ex)

	}

	FindDuplicateFile(pathname)
}


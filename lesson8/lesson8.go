package lesson8

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// храним состояние флагов
var curFlag programFlag

// map для всех найденных уникальных файлов
var files = make(map[FileInfo]string)

var mutex = make(chan struct{}, 1)
var wg sync.WaitGroup

//FindDuplicateFile - функция ищущая дубликаты файлов
func FindDuplicateFile(path string) {

	err := filepath.Walk(path, checkDuplicate)
	if err != nil {
		fmt.Println(err)
	}

}

//checkDuplicate - проверка дубликата файла по имени и размеру
//занесение уникальных файлов в map
func checkDuplicate(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	f := FileInfo{info.Name(), info.Size()}
	if v, ok := files[f]; ok {
		wg.Add(1)
		go func(path, v string) {
			defer wg.Done()
			mutex <- struct{}{} // lock
			fmt.Printf("%q is a duplicate of %q\n", path, v)
			// Если разрешено - удаляем
			if curFlag.DeleteDuplicate {
				duplicateDelete(path, curFlag.ConfirmDelete)
			}
			<-mutex //unlock
		}(path, v)

	} else {
		files[f] = path
	}

	return nil
}

// Удаление дубликата
func duplicateDelete(path string, isAsk bool) error {

	// Если надо, спрашиваем разрешение на удаление
	if isAsk {
		if !askForConfirmation("Delete duplicate?") {
			return nil
		}
	}

	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("File %s deleted!\n", path)
	return nil
}

func askForConfirmation(s string) bool {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/n]: ", s)
	res, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return false
	}
	res = strings.ToLower(strings.TrimSpace(res))
	if res == "y" || res == "yes" {
		return true
	}
	return false
}

//Start - обработка флагов и старт поиска дубликатов
func Start(delFlag, confFlag bool) {
	// Читаем флаги
	if flag.Lookup("d") == nil {
		flag.BoolVar(&curFlag.DeleteDuplicate, "d", delFlag, "Delete duplicate files")
		flag.BoolVar(&curFlag.ConfirmDelete, "y", confFlag, "Confirm file deletion")
		flag.Parse()
	} else {
		curFlag.DeleteDuplicate = delFlag
		curFlag.ConfirmDelete = confFlag
	}

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
	wg.Wait()
}

package main

import (
	//	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//	flag.Parse()
	//	root := flag.Arg(0)
	//getFilelist("E:/STUDY")
	copyDir("G:\\go_workspace\\GOPATH\\src\\pjx\\example", "G:\\pjx_path\\fwhezfwhez\\example")
	//	w, err := CopyFile("E:/filecopy.go", "e:/test.go")
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}

	//	fmt.Println(w)
}
func copyDir(src string, dest string) {
	src_original := src
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			//			fmt.Println(f.Name())
			//			copyDir(f.Name(), dest+"/"+f.Name())
		} else {
			//fmt.Println(src)
			//fmt.Println(src_original)
			//fmt.Println(dest)

			dest_new := strings.Replace(src, src_original, dest, -1)
			//fmt.Println(dest_new)
			//fmt.Println("CopyFile:" + src + " to " + dest_new)
			CopyFile(src, dest_new)
		}
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

//egodic directories
func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//copy file
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	dst_slices := strings.Split(dst, "\\")
	dst_slices_len := len(dst_slices)
	dest_dir := ""
	for i := 0; i < dst_slices_len-1; i++ {
		dest_dir = dest_dir + dst_slices[i] + "\\"
	}

	b, err := PathExists(dest_dir)
	if b == false {
		err := os.Mkdir(dest_dir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}
	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

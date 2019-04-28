package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Mkdir("file", 0777)
	http.Handle("/pollux/", http.StripPrefix("/pollux/", http.FileServer(http.Dir("G:\\go_workspace\\GOPATH\\src\\test_X\\test_file_server"))))
	http.HandleFunc("/file/upload/", func(w http.ResponseWriter, r *http.Request){
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create("G:\\go_workspace\\GOPATH\\src\\test_X\\test_file_server" + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		w.Write([]byte("保存成功"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

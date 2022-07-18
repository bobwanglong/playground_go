package main

import (
	"fmt"
	"log"
	"os"
)

func fileCreate() {
	fmt.Println("<file create>")
	f, err := os.Create("bob.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("<file info>")
	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fileInfo's name:", fileInfo.Name())
	fmt.Println("fileInfo's premissions:", fileInfo.Mode())
	fmt.Println("fileInfo's modeTime:", fileInfo.ModTime())

	fmt.Println("<file premissions>")
	err = f.Chmod(0777)
	if err != nil {
		log.Fatalf("chmod file failed err=%s\n", err)
	}

	// 改变拥有者
	err = f.Chown(os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatalf("chown file failed err=%s\n", err)
	}

	// 再次获取文件信息 验证改变是否正确
	fileInfo, err = f.Stat()
	if err != nil {
		log.Fatalf("get file info second failed err=%s\n", err)
	}
	log.Printf("File change Permissions is %s\n", fileInfo.Mode())

	// 关闭文件
	err = f.Close()
	if err != nil {
		log.Fatalf("close file failed err=%s\n", err)
	}

	//删除文件
	err = os.Remove("bob.md")
	if err != nil {
		log.Fatalf("remove file failed err=%s\n", err)
	}
}

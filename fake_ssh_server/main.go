package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gliderlabs/ssh"
)

func main() {
	filePath := "./record.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	//写入文件时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	ssh.ListenAndServe(":2223", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
			writer.WriteString(ctx.User() + "\t" + password + "\n")
			writer.Flush()
			log.Println(ctx.User(), password)
			return false
		}),
	)
	for {
		time.Sleep(time.Minute)
	}
}

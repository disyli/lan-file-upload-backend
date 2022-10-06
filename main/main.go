package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func upldFiles(c *gin.Context) {
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		log.Fatal(err)
	}
	// 获取表单
	form := c.Request.MultipartForm
	//log.Println(*form)
	// 获取参数upload后面的多个文件名，存放到数组files里面，
	files := form.File["file"]
	//log.Println(form.File["file"])
	// 遍历数组，每取出一个file就拷贝一次
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		fileName := files[i].Filename
		fmt.Println(fileName)

		out, err := os.Create(fileName)
		defer out.Close()
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusCreated, "uploadFiles success! \n")
	}
}

func main() {
	router := gin.Default()

	router.POST("/upload", upldFiles)

	router.Run(":8088")
}

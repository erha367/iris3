package main

import (
	"flag"
	"fmt"
	"io"
	"iris3/tools"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var path, url, del string

func main() {
	//参数解析
	flag.StringVar(&del, "del", "", "是否删除源文件")
	flag.StringVar(&url, "url", "", "输入图片的url路径")
	flag.StringVar(&path, "path", "", "输入图片的绝对路径")
	flag.Parse()
	if url != `` {
		path2, err := downloadImage(url, `./tmp.jpg`)
		if err != nil {
			log.Println(err)
			return
		}
		path = path2
	}
	if len(del) > 0 {
		//删除源文件
		defer os.Remove(path)
	}
	//解析
	tools.DrawLines(path)
	return
}

// downloadImage 下载远程图片并保存到本地，返回文件的绝对路径
func downloadImage(url string, filename string) (string, error) {
	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("无法获取图片: %v", err)
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("无法下载图片，状态码: %d", resp.StatusCode)
	}

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("无法创建文件: %v", err)
	}
	defer file.Close()

	// 将响应体写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("无法写入文件: %v", err)
	}

	// 获取绝对路径
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return "", fmt.Errorf("无法获取绝对路径: %v", err)
	}

	return absPath, nil
}

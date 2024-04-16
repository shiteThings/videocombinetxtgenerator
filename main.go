package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 从键盘输入获取目录路径
	var dirPath string
	fmt.Print("请输入目录路径：")
	fmt.Scanln(&dirPath)

	// 检查路径是否为Windows文件路径，并替换分隔符
	if strings.Contains(dirPath, "\\") {
		dirPath = strings.ReplaceAll(dirPath, "\\", "\\\\")
	}

	// 获取目录下所有文件的文件名
	fileNames, err := getFileNames(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 拼接txt文件的完整路径
	txtFilePath := filepath.Join(dirPath, "file_names.txt")

	// 创建并打开txt文本文件
	txtFile, err := os.Create(txtFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer txtFile.Close()

	// 将文件名按行写入txt文件
	for _, fileName := range fileNames {
		txtFile.WriteString(fmt.Sprintf("file '%s'\n", fileName))
	}

	fmt.Printf("File names written to %s successfully.\n", txtFilePath)
}

// 获取目录下所有文件的文件名
func getFileNames(dirPath string) ([]string, error) {
	var fileNames []string

	// 遍历目录
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// 提取文件名并加入到fileNames切片中
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

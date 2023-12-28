package util

import (
	"encoding/json"
	"fmt"
	"github.com/crytjy/jkd-leaf/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateDir 创建文件夹
func CreateDir(path string) error {
	_, err := os.Stat(path) // 检查文件夹是否存在
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm) // 不存在则创建
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveFile 删除文件
func RemoveFile(filename string) {
	// 判断文件是否存在
	if _, err := os.Stat(filename); err == nil {
		// 如果文件不存在，则创建文件
		err := os.Remove(filename)
		if err != nil {
			log.Debug("RemoveFile:"+filename, err)
			return
		}
	}
}

// CreateFile 创建文件
func CreateFile(filePath string) {
	// 判断文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果文件不存在，则创建文件
		file, err := os.Create(filePath)
		if err != nil {
			return
		}
		defer file.Close() // 创建文件后，立即关闭文件
	}
}

// FileWrite 打开文件，并进行写入操作
func FileWrite(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil
	}
	//defer file.Close()

	return file
}

func GetFileContent(path string) interface{} {
	// 打开 JSON 文件
	file, err := os.Open(path)
	if err != nil {
		log.Debug("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Debug("Error reading file:", err)
		return nil
	}

	// 使用 interface{} 解析未知 JSON 结构
	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Debug("Error unmarshalling JSON:", err)
		return nil
	}

	return result
}

// RemoveExtension 移除文件名的扩展名
func RemoveExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

// RemoveFolder 清除文件夹下的文件
func RemoveFolder(folderPath string) {
	// 获取文件夹下的所有文件
	files, err := filepath.Glob(filepath.Join(folderPath, "*"))
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}

	// 遍历文件列表并删除每个文件
	for _, file := range files {
		RemoveFile(file)
	}
}

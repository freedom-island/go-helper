package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

// GetDirFirstLastFile 获取指定文件夹内符合指定格式的最旧与最新文件名
func GetDirFirstLastFile(format, folderPath string) ([]string, error) {
	// 使用 os.ReadDir 读取文件夹内容
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return []string{}, err
	}

	var matchedFiles []os.DirEntry

	// 遍历文件并检查文件名是否符合时间格式
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		_, err = time.Parse(format, entry.Name())
		if err == nil {
			matchedFiles = append(matchedFiles, entry)
		}
	}

	// 如果没有符合格式的文件，返回空字符串和错误
	if len(matchedFiles) == 0 {
		return []string{}, ErrFileNotFound
	}

	// 按时间排序
	sort.Slice(matchedFiles, func(i, j int) bool {
		timeI, _ := time.Parse(format, matchedFiles[i].Name())
		timeJ, _ := time.Parse(format, matchedFiles[j].Name())
		return timeI.After(timeJ)
	})

	// 返回最新、最旧的文件路径
	return []string{matchedFiles[0].Name(), matchedFiles[len(matchedFiles)-1].Name()}, nil
}

// SaveJsonToFile 保存 JSON 数据到文件
func SaveJsonToFile(data interface{}, filename string, replace bool) error {
	filename = fmt.Sprintf("tmp/%s.json", filename)

	// 序列化数据
	serialized, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}

	// 打开文件，追加模式
	fileStatus := os.O_APPEND | os.O_WRONLY | os.O_CREATE
	if replace {
		fileStatus = os.O_CREATE | os.O_RDWR | os.O_TRUNC
	}
	file, err := os.OpenFile(filename, fileStatus, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 写入实际数据
	_, err = file.Write(serialized)
	if err != nil {
		return fmt.Errorf("failed to write serialized data: %w", err)
	}

	return nil
}

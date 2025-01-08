package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// 配置文件
type Config struct {
	Host   string  `json:"host"`
	Sleep  int     `jsson:"sleep"`
	FFmpeg FFmpeg  `json:"ffmpeg"`
	Video  []Video `json:"video"`
}

// 视频转码配置
type FFmpeg struct {
	Exec string `json:"exec"`
	Type string `json:"type"`
	Gpu  string `json:"gpu"`
}

// 视频录像机配置
type Video struct {
	WsHost   string `json:"wsHost"`
	ParamMsg string `json:"paramMsg"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
}

// 获取配置
func GetConfig() Config {
	var config Config
	filePath := "config.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		FmtPrint("配置文件不存在", err)
		os.Exit(0)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		FmtPrint("读取配置文件出错", err)
		os.Exit(0)
	}
	return config
}

// 删除文件夹下的旧文件
func DeleteOldFiles(dirPath string, filesToKeep int) {
	//打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		FmtPrint("无法打开目录", err)
		return
	}
	defer dir.Close()
	//读取目录下的所有文件信息
	files, err := dir.Readdir(-1) //读取所有文件
	if err != nil {
		FmtPrint("读取目录失败", err)
		return
	}
	//过滤掉目录，仅保留文件
	var fileInfos []os.FileInfo
	for _, file := range files {
		if !file.IsDir() { //仅保留文件，排除子目录
			fileInfos = append(fileInfos, file)
		}
	}
	//如果文件数量小于或等于需要保留的文件数量，则不需要删除
	if len(fileInfos) <= filesToKeep {
		return
	}
	//按照文件的修改时间进行排序，最新的排在前面
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
	})
	//删除最旧的文件，直到只剩下指定数量的文件
	for i := filesToKeep; i < len(fileInfos); i++ {
		oldFile := filepath.Join(dirPath, fileInfos[i].Name())
		_ = os.Remove(oldFile)
	}
}

// 定义内置的打印语句
func FmtPrint(data ...any) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if len(data) == 1 {
		fmt.Println(date+": ", data[0])
	} else {
		fmt.Println(date+": ", data)
	}
}

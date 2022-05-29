package logging

import (
	"fmt"
	"github.com/jamesluo111/gin-blog/pkg/file"
	"github.com/jamesluo111/gin-blog/pkg/setting"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getFileName() string {
	return fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	//获取当前文件的绝对路径
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Getwd error:%v", err)
	}
	src := dir + "/" + filePath
	//判断目录是否需要权限
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("permission denied src:%s", src)
	}
	//判断文件目录是否存在
	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src:%s", src)
	}

	//创建文件
	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("file.Open src:%s", src)
	}
	return f, nil

}

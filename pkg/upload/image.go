package upload

import (
	"fmt"
	"github.com/jamesluo111/gin-blog/pkg/file"
	"github.com/jamesluo111/gin-blog/pkg/logging"
	"github.com/jamesluo111/gin-blog/pkg/setting"
	"github.com/jamesluo111/gin-blog/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

//获取图片完整访问 URL
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

//获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//获取图片路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

//获取图片完整路径
func GetFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

//检查文件后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)

	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

//检查文件大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

//检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd error:%v", err)
	}

	//检查文件是否存在
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir error:%v", err)
	}

	//检查权限
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission src:%s", src)
	}

	return nil
}

package article_service

import (
	"github.com/golang/freetype"
	"github.com/jamesluo111/gin-blog/pkg/file"
	"github.com/jamesluo111/gin-blog/pkg/qrcode"
	"github.com/jamesluo111/gin-blog/pkg/setting"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePost(posterName string, article *Article, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
		Qr:         qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}

	return true
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	//获取二维码储存路径
	fullPath := qrcode.GetQrCodeFullPath()
	//生成二维码图像
	fileName, path, err := a.Qr.EnCode(fullPath)
	if err != nil {
		return "", "", err
	}

	//检查合并后图像是否存在
	if !a.CheckMergedImage(path) {
		//生成合并后图像文件
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()
		//打开背景图片
		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()
		//打开生成的二维码图片
		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()
		//解码背景图片
		bgImage, err := png.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		//解码二维码图片
		qrImage, err := jpeg.Decode(qrF)

		if err != nil {
			return "", "", err
		}
		//创建一个新的rgba图像
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
		//在rgba图像上绘制背景图
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		//在已绘制背景图的rgba上指定Point上绘制二维码图像
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
		//将绘制好的 RGBA 图像以 JPEG 4：2：0 基线格式写入合并后的图像文件
		err = a.DrawPoster(&DrawText{
			JPG:      jpg,
			Merged:   mergedF,
			Title:    "gin golang",
			X0:       300,
			Y0:       30,
			Size0:    30,
			SubTitle: "---罗",
			X1:       400,
			Y1:       60,
			Size1:    20,
		}, "msyhbd.ttc")
		if err != nil {
			return "", "", err
		}
		jpeg.Encode(mergedF, jpg, nil)
	}
	return fileName, path, nil
}

type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubTitle string
	X1       int
	Y1       int
	Size1    float64
}

func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	//获取源字体路径
	fontSource := setting.AppSetting.RuntimeRootPath + setting.AppSetting.FontSavePath + fontName
	//读取源字体
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}

	//解析字体库
	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}
	//创建新的context设置一些默认值
	fc := freetype.NewContext()
	//设置屏幕每英寸的分辨率
	fc.SetDPI(72)
	//设置用于绘制文本的字体
	fc.SetFont(trueTypeFont)
	//设置文本字体的大小,以磅为单位
	fc.SetFontSize(d.Size0)
	//设置裁剪矩形以进行绘制
	fc.SetClip(d.JPG.Bounds())
	//设置目标图像
	fc.SetDst(d.JPG)
	//置绘制操作的源图像，通常为 image.Uniform
	fc.SetSrc(image.Black)
	//设置绘制的起始坐标
	pt := freetype.Pt(d.X0, d.Y0)
	//在pt位置画title
	_, err = fc.DrawString(d.Title, pt)
	//重新设置文字大小用来画二级标题
	fc.SetFontSize(d.Size1)
	//画二级标题
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1))
	if err != nil {
		return err
	}

	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}

	return nil
}

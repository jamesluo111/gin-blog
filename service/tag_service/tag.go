package tag_service

import (
	"encoding/json"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jamesluo111/gin-blog/models"
	"github.com/jamesluo111/gin-blog/pkg/export"
	"github.com/jamesluo111/gin-blog/pkg/gredis"
	"github.com/jamesluo111/gin-blog/pkg/logging"
	"github.com/jamesluo111/gin-blog/service/cache_service"
	"io"
	"strconv"
	"time"
)

type Tag struct {
	ID         int
	Name       string
	State      int
	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagById(t.ID)
}

func (t *Tag) Get() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)
	//从redis中获取数据
	cache := cache_service.Tag{
		Id:    t.ID,
		Name:  t.Name,
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	//从数据库中获取数据
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, tags, 3600)
	return tags, nil

}

func (t *Tag) Add() error {
	tag := map[string]interface{}{
		"name":       t.Name,
		"created_by": t.CreatedBy,
		"state":      t.State,
	}
	return models.AddTag(tag)
}

func (t *Tag) Edit() error {
	return models.EditTag(t.ID, map[string]interface{}{
		"name":        t.Name,
		"state":       t.State,
		"modified_by": t.ModifiedBy,
	})
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if t.State != -1 {
		maps["state"] = t.State
	}
	if t.Name != "" {
		maps["name"] = t.Name
	}
	return maps
}

func (t *Tag) Export() (string, error) {
	////先从数据库中获取数据
	//tags, err := t.Get()
	//if err != nil {
	//	return "", err
	//}
	//titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	//
	////创建xlsx文件
	//file := xlsx.NewFile()
	////给文件加一个工作表
	//sheet, err := file.AddSheet("标签信息")
	//if err != nil {
	//	return "", err
	//}
	////新增标题列
	//row := sheet.AddRow()
	////定义单元格
	//var cell *xlsx.Cell
	////循环titles给当前列添加单元格内容
	//for _, title := range titles {
	//	cell = row.AddCell()
	//	cell.Value = title
	//}
	//
	//for _, v := range tags {
	//	values := []string{
	//		strconv.Itoa(v.Id),
	//		v.Name,
	//		v.CreatedBy,
	//		strconv.Itoa(v.CreatedOn),
	//		v.ModifiedBy,
	//		strconv.Itoa(v.ModifiedOn),
	//	}
	//	row = sheet.AddRow()
	//	for _, value := range values {
	//		cell = row.AddCell()
	//		cell.Value = value
	//	}
	//}
	//
	//timeNow := strconv.Itoa(int(time.Now().Unix()))
	//fileName := "tags-" + timeNow + ".xlsx"
	//fullPath := export.GetExcelFullPath() + fileName
	//err = file.Save(fullPath)
	//if err != nil {
	//	return "", err
	//}
	//return fileName, nil

	//先从数据库中获取数据
	tags, err := t.Get()
	if err != nil {
		return "", err
	}
	x1 := excelize.NewFile()
	x1.SetSheetName("Sheet1", "标签信息")
	row := 1
	for index, v := range tags {
		if index == 0 {
			//设置表头信息
			title := []string{
				"ID",
				"标签名",
				"创建人",
				"创建时间",
				"修改人",
				"修改时间",
			}
			x1.SetSheetRow("标签信息", "A1", &title)
		}
		values := []string{
			strconv.Itoa(v.Id),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}
		row++
		x1.SetSheetRow("标签信息", "A"+strconv.Itoa(row), &values)
	}
	timeNow := strconv.Itoa(int(time.Now().Unix()))
	fileName := "tags-" + timeNow + ".xlsx"
	fullPath := export.GetExcelFullPath() + fileName
	err = x1.SaveAs(fullPath)
	if err != nil {
		return "", err
	}
	return fileName, nil

}

func (t *Tag) Import(r io.Reader) error {
	//打开excel文件流
	xlsx1, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	//打开标签信息工作表
	rows := xlsx1.GetRows("标签信息")
	for irow, row := range rows {
		//判断excel表头是否按照要求
		if irow == 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			if data[1] != "名称" && data[2] != "创建人" {
				return errors.New("格式错误!")
			}
		}
		//irow行数据,判断工作表内是否为空
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			var maps = map[string]interface{}{
				"name":       data[1],
				"state":      1,
				"created_by": data[2],
			}
			if err := models.AddTag(maps); err != nil {
				return err
			}
		}

	}
	return nil
}

package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_CHECK_IMAGE_FOEMAT: "图片大小或后缀不对",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "图片路径错误或权限不足",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "图片保存失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
	ERROR_GET_ARTICLE_FAIL:          "文章获取错误",
	ERROR_TOTLE_ARTICLE_FAIL:        "文章获取总数量失败",
	ERROR_GETALL_ARTICLE_FAIL:       "批量获取文章失败",
	ERROR_ADD_ARTICLE_FAIL:          "文章添加失败",
	ERROR_EDIT_ARTICLE_FAIL:         "文章修改失败",
	ERROR_DELETE_ARTICLE_FAIL:       "文章删除失败",
	ERROR_GET_TAG_FAIL:              "标签获取失败",
	ERROR_COUNT_TAG_FAIL:            "标签数量获取失败",
	ERROR_EXIST_BY_NAME_FAIL:        "标签是否存在报错",
	ERROR_ADD_TAG_FAIL:              "标签添加报错",
	ERROR_EDIT_TAG_FAIL:             "标签编辑报错",
	ERROR_DELETE_TAG_FAIL:           "标签删除报错",
	ERROR_EXPORT_TAG_FAIL:           "标签导出报错",
	ERROR_IMPORT_TAG_FAIL:           "标签导入报错",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "文章二维码生成报错",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[code]
}

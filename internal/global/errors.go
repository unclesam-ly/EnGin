package global

import "errors"

// 预定义全局通用的系统级/数据库级错误
var (
	ErrDbQuery  = errors.New("数据库查询异常")
	ErrDbSave   = errors.New("数据保存失败")
	ErrDbDelete = errors.New("数据删除失败")
	ErrDbUpdate = errors.New("数据更新失败")
)

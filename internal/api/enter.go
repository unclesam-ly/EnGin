package api

import (
	"EnGin/internal/api/auth_api"
	"EnGin/internal/api/user_api"
)

type Api struct {
	AuthApi auth_api.AuthApi
	UserApi user_api.UserApi
}

// App 统一导出全局 API 实例
var App = new(Api)

package auth_api

import (
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/service/auth_ser"
	"EnGin/internal/utils/jwts"
	"EnGin/internal/utils/res"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (AuthApi) LoginView(c *gin.Context) {
	cr := middleware.GetBind[LoginRequest](c)

	u, err := auth_ser.Login(c.Request.Context(), cr.Username, cr.Password)
	if err != nil {
		if errors.Is(err, auth_ser.ErrUserNotFound) {
			global.Log.Warn("用户不存在", zap.String("username", cr.Username), zap.Error(err))
			res.BadRequest("用户名或密码错误", c)
			return
		}

		if errors.Is(err, auth_ser.ErrAccount) {
			res.FailWithMsg("用户名或密码错误", c)
			return
		}

		global.Log.Error("登陆系统异常", zap.Error(err))
		res.InternalError("登陆失败,请稍后重试", c)
		return
	}

	// 提取预加载的角色 ID 列表
	var roleIds []uint
	for _, role := range u.Edges.Roles {
		roleIds = append(roleIds, uint(role.ID))
	}

	claims := jwts.Claims{
		UserID: uint(u.ID),
		Roles:  roleIds, // 将关联查询出的角色 ID 塞进 token
		Email:  u.Email,
	}

	accessToken, err := jwts.GenAccessToken(claims)
	if err != nil {
		res.InternalError("生成 accessToken 失败", c)
		return
	}

	refreshToken, err := jwts.GenRefreshToken(uint(u.ID))
	if err != nil {
		res.InternalError("生成 refreshToken 失败", c)
		return
	}

	res.OkWithData(LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, c)
}

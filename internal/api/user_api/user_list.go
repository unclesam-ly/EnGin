package user_api

import (
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/service/user_ser"
	"EnGin/internal/utils/res"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserListRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	UsernameKey string `form:"username_key"` // 可选模糊搜索
}

type UserListResponse struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	IP       string    `json:"ip"`
	CreateAt time.Time `json:"create_at"`
}

func (UserApi) UserListView(c *gin.Context) {
	cr := middleware.GetBind[UserListRequest](c)

	list, count, err := user_ser.GetUserListWithPage(
		c.Request.Context(),
		cr.UsernameKey,
		cr.Page,
		cr.Limit,
	)
	if err != nil {
		global.Log.Error("查询用户列表失败", zap.Error(err))
		res.InternalError("获取用户列表失败", c)
		return
	}

	var users = make([]UserListResponse, 0, count)
	for _, u := range list {
		users = append(users, UserListResponse{
			Username: u.Username,
			Email:    u.Email,
			IP:       u.IP,
			CreateAt: u.CreatedAt,
		})
	}

	res.OkWithList(users, count, c)
}

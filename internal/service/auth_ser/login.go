package auth_ser

import (
	"EnGin/internal/db"
	"EnGin/internal/ent"
	"EnGin/internal/ent/user"
	"EnGin/internal/global"
	"EnGin/internal/utils/pwd"
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("用户不存在")
	ErrAccount      = errors.New("账号和密码错误")
)

func Login(ctx context.Context, username, password string) (*ent.User, error) {
	u, err := db.Client.User.Query().Where(user.UsernameEQ(username)).WithRoles().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}

		return nil, global.ErrDbQuery
	}

	if !pwd.CompareHashAndPassword(u.Password, password) {
		return nil, ErrAccount
	}

	return u, nil
}

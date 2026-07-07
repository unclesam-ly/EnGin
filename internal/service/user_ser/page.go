package user_ser

import (
	"EnGin/internal/db"
	"EnGin/internal/ent"
	"EnGin/internal/ent/user"
	"EnGin/internal/service/common"
	"context"

	"entgo.io/ent/dialect/sql"
)

func GetUserListWithPage(ctx context.Context, usernameKey string, page, limit int) ([]*ent.User, int, error) {
	return common.QueryList(ctx, db.Client.User.Query(), common.QueryOptions[*ent.UserQuery]{
		Page:  page,
		Limit: limit,
		Query: func(q *ent.UserQuery) *ent.UserQuery {
			// 模糊匹配（编译期安全，杜绝拼写错误）
			if usernameKey != "" {
				q = q.Where(user.UsernameContains(usernameKey))
			}

			// 预加载关联角色，并按照创建时间倒序
			return q.WithRoles().Order(user.ByCreatedAt(sql.OrderDesc()))
		},
	})
}

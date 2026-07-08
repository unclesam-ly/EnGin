package ent

// GetPrimaryRole 获取用户主角色（权限最高的角色）
func (u *User) GetPrimaryRole() *Role {
	// 如果没有预加载角色，或者角色列表为空，直接返回 nil
	if u.Edges.Roles == nil || len(u.Edges.Roles) == 0 {
		return nil
	}

	primary := u.Edges.Roles[0]
	for _, r := range u.Edges.Roles {
		if r.Code < primary.Code {
			primary = r
		}
	}

	return primary
}

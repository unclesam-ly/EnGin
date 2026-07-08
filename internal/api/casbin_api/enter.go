package casbin_api

type CasbinApi struct {
}

type PolicyRequest struct {
	RoleID string `json:"role_id" binding:"required" label:"角色ID"`
	Path   string `json:"path" binding:"required" label:"访问路径"`
	Method string `json:"method" binding:"required" label:"请求方法"`
}

type RoleInheritanceRequest struct {
	RoleID      string `json:"role_id" binding:"required" label:"角色ID"`
	InheritFrom string `json:"inherit_from" binding:"required" label:"继承自角色ID"`
}

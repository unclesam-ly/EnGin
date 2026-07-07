package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin 定义公共的基础混入
type BaseMixin struct {
	mixin.Schema
}

// Fields 定义所有表都拥有的公共字段
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间：默认当前时间，不可修改（Immutable）
		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		// 更新时间：默认当前时间，并且在每次数据 Update 时自动更新为当前时间
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

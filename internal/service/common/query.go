package common

import "context"

// EntQuery 定义满足 Ent 所有查询构建器的泛型接口
// Q 为具体的查询器类型（如 *ent.UserQuery），T 为返回的实体类型（如 *ent.User）
type EntQuery[Q any, T any] interface {
	Limit(int) Q
	Offset(int) Q
	Count(ctx context.Context) (int, error)
	All(ctx context.Context) ([]T, error)
}

// QueryOptions 通用列表查询配置项
type QueryOptions[Q any] struct {
	Page  int
	Limit int
	Query func(Q) Q // Query 闭包：在这里传入具体的过滤条件（Where）、排序（Order）和预加载（With）
}

func QueryList[Q EntQuery[Q, T], T any](ctx context.Context, query Q, opts QueryOptions[Q]) (list []T, count int, err error) {
	if opts.Query != nil {
		query = opts.Query(query)
	}

	count, err = query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	page := opts.Page
	if page <= 0 {
		page = 1
	}

	limit := opts.Limit
	if limit <= 0 {
		limit = 10 // 默认每页 10 条
	}
	offset := (page - 1) * limit

	list, err = query.Limit(limit).Offset(offset).All(ctx)

	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

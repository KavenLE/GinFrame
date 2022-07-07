package types

type PageForm struct {
	pageNum  int64 `form:"pageNum"`
	pageSize int64 `form:"pageSize"`
}

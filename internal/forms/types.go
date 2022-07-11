package forms

type PageForm struct {
	pageNum  int64 `form:"pageNum"`
	pageSize int64 `form:"pageSize"`
}
type TestForm struct {
	Param string `json:"param"`
}

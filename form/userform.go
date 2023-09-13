package form

type User struct {
	Id   int64 `json:"id" `
	Name string `json:"name" binding:"required"`
}

package form

type User struct{
	Name string `json:"name" binding:"required"`
}
package form

type User struct{
	Id string `json:"name" binding:"required"`
	Name string `json:"name" binding:"required"`
}
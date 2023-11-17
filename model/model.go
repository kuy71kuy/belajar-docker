package model

type Config struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// dibuat UserResponse karena untuk kebutuhan jwt, kalo ki ta tetap pakai User saja, nanti di database malah bikin kolom baru yaitu token yang mana gaperlu dibuat di database jg
type UserResponse struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

type Book struct {
	Id        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Writer    string `json:"writer" form:"writer"`
	Publisher string `json:"publisher" form:"publisher"`
}

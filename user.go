package todo

type User struct {
	Id       int    `form:"-" db:"id"`
	Name     string `form:"name"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password1" binding:"required"`
}

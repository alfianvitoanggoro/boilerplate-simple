package user_model

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
)

type User struct {
	ID       uint
	Username string
	Password string
	Role     Role
}

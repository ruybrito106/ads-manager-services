package users

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

var (
	superUser = &User{
		ID:       "superuser",
		Password: "12345",
	}
)

func (u *User) IsSuperuser() bool {
	return u.ID == superUser.ID && u.Password == superUser.Password
}

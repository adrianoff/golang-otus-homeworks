package hw10programoptimization

//go:generate /home/adrianoff/go/bin/easyjson -all user.go
type User struct {
	ID       int    `json:"-"`
	Name     string `json:"-"`
	Username string `json:"-"`
	Email    string
	Phone    string `json:"-"`
	Password string `json:"-"`
	Address  string `json:"-"`
}

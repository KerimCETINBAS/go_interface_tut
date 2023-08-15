package services

type User struct {
	name string
}
type userService struct {
	users []User
}

type IUserService interface {
	AddUser(value string)
	GetUsers() []User
}

func UserService() IUserService {
	return &userService{}
}

func (s *userService) AddUser(value string) {
	s.users = append(s.users, User{name: value})
}

func (s *userService) GetUsers() []User {
	return s.users
}

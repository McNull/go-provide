package pkg

type UserRepository interface {
	FindAll() []User
}

type userRepository struct {
	users []User
}

func (r *userRepository) FindAll() []User {
	return r.users
}

func NewUserRepository(users []User) *userRepository {
	return &userRepository{
		users: users,
	}
}

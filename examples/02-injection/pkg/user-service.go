package pkg

type UserService interface {
	PrintAll() string
}

type userService struct {
	repository UserRepository
}

func (s *userService) PrintAll() string {
	users := s.repository.FindAll()
	result := ""
	for _, user := range users {
		result += user.Name + "\n"
	}
	return result
}

func NewUserService(repository UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

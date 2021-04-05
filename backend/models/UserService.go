package models

func NewUserService(services *Services) UserService {
	cg := &userMongo{services}
	cv := newClientValidator(cg, services)
	defineClientValidators(cv)
	return &userService{
		UserModel: cv,
		services:  services,
	}
}

// UserService is a set of methods used to manipulate and
// work with the user model
type UserService interface {
	// Authenticate will verify the provided email address and
	// password are correct. If they are correct, the user
	// corresponding to that email will be returned.
	Authenticate() (*User, error)

	UserModel
}
type UserDB interface {
	// Methods for querying for single user
	FindById(id uint32) (*User, error)
}

var _ UserService = &userService{}

type userService struct {
	UserModel
	services *Services
}

func (us *userService) Authenticate() (*User, error) {
	var user User

	return &user, nil
}

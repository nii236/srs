package mock

import "github.com/nii236/srs"

// UserService represents a mock implementation of srs.UserService.
type UserService struct {
	UserFn      func(id int) (*srs.User, error)
	UserInvoked bool

	UsersFn      func() ([]*srs.User, error)
	UsersInvoked bool

	// additional function implementations...
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int) (*srs.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}

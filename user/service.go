package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	// IsEmailAvailable(input CheckEmailUserInput) (User, error)
	SaveAvatar(ID string, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.ID = uuid.NewString()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.RoleID = "pegawai"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil

}

func (s *service) LoginUser(input LoginUserInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == "0" {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil

}

func (s *service) SaveAvatar(ID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.Avatar = fileLocation
	updateUser, err := s.repository.Update(user)

	if err != nil {
		return updateUser, err
	}
	return updateUser, nil
}

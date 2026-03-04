package user

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	"github.com/DotNicolasPenha/Posts-CRUD/pkg/JWT"
	"github.com/DotNicolasPenha/Posts-CRUD/pkg/hasher"
)

type Service struct {
	repository Repository
}

func NewService(repository *Repository) *Service {
	if repository == nil {
		logger.Fatal("repository of user service is nil")
	}
	return &Service{
		repository: *repository,
	}
}
func (s *Service) AddUser(createUserDto CreateUserDTO) error {
	if createUserDto.Username == "" {
		return ErrNameUserIsNil
	}
	if createUserDto.Bio == "" {
		return ErrBioIsNil
	}
	if createUserDto.Password == "" {
		return ErrPasswordIsNil
	}
	password, err := hasher.HashPassword(createUserDto.Password)
	if err != nil {
		return err
	}
	createUserDto.Password = password
	return s.repository.Insert(createUserDto)
}

func (s *Service) GetUsers() ([]UserPreviewPerson, error) {
	return s.repository.FindMany()
}

func (s *Service) LoginUser(userLogin LoginUserDTO) (string, error) {
	if userLogin.Username == "" {
		return "", ErrNameUserIsNil
	}
	if userLogin.Password == "" {
		return "", ErrPasswordIsNil
	}

	user, err := s.repository.FindByNameUser(userLogin.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrNotFound
	}

	if !hasher.ComparePassword(user.PasswordHash, userLogin.Password) {
		return "", ErrIncorrectPassword
	}

	token, err := JWT.GenerateToken(user.ID.UUID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

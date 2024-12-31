package usecase

import (
	"librarybackend/domain"
	"time"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
	TokenGen       domain.TokenGenerator
	PasswordSvc    domain.PasswordService
}

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration, tokenGen domain.TokenGenerator, passwordSvc domain.PasswordService) domain.UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		contextTimeout: timeout,
		TokenGen:       tokenGen,
		PasswordSvc:    passwordSvc,
	}
}

func (uc *UserUseCase) CreateAccount(user domain.User) (domain.User, domain.ErrorResponse) {

	check, _ := uc.UserRepository.GetAllUserByEmial(user.Email)

	if check.Email != "" {
		return domain.User{}, domain.ErrorResponse{Message: "Email already exist", Status: 400}
	}

	// user.Role = strings.ToLower(user.Role)

	hashedPassword, err := uc.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, domain.ErrorResponse{Message: "Failed to hash password", Status: 500}
	}
	user.Password = hashedPassword

	res, err := uc.UserRepository.CreateAccount(user)

	if err != nil {
		return domain.User{}, domain.ErrorResponse{Message: "Failed to create user", Status: 500}
	}

	return res, domain.ErrorResponse{}
}

func (uc *UserUseCase) Login(newUser domain.User) (domain.LoginSucessResponse, domain.ErrorResponse) {

	user, err := uc.UserRepository.GetAllUserByEmial(newUser.Email)
	if err != nil {
		return domain.LoginSucessResponse{}, domain.ErrorResponse{Message: "User not found", Status: 404}
	}

	match := uc.PasswordSvc.CheckPasswordHash(newUser.Password, user.Password)

	if !match {
		return domain.LoginSucessResponse{}, domain.ErrorResponse{Message: "Invalid email or password", Status: 400}
	}

	token, err := uc.TokenGen.GenerateToken(user)

	if err != nil {
		return domain.LoginSucessResponse{}, domain.ErrorResponse{Message: "Failed to generate token", Status: 500}
	}

	user.Password = ""

	return domain.LoginSucessResponse{
		Message:    "Login successful",
		Status:     200,
		AcessToken: token,
		UserData:   user,
	}, domain.ErrorResponse{}
}

func (uc *UserUseCase) GetByID(id string) (domain.User, domain.ErrorResponse) {
	user, err := uc.UserRepository.GetByID(id)

	if err != nil {
		return domain.User{}, domain.ErrorResponse{Message: "User not found", Status: 404}
	}

	return user, domain.ErrorResponse{}
}

func (uc *UserUseCase) UpdateProfile(id string, user domain.User) (domain.User, domain.ErrorResponse) {
	_, err := uc.UserRepository.UpdateProfile(id, user)

	if err != nil {

		return domain.User{}, domain.ErrorResponse{Message: "Failed to update user", Status: 500}
	}

	return user, domain.ErrorResponse{}
}

func (uc *UserUseCase) GetAllUser() ([]domain.User, error) {
	users, err := uc.UserRepository.GetAllUser()

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (uc *UserUseCase) GetUserByID(id string) (domain.User, error) {

	user, err := uc.UserRepository.GetByID(id)

	if err != nil {
		return domain.User{}, err

	}

	return user, nil
}

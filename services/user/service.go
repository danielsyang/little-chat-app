package user

import (
	"chat-app/database/models/users"
	authService "chat-app/services/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserDomain struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPass), err
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateUser(createUserDomain *CreateUserDomain) (id string, err error) {

	hashedPassword, err := hashPassword(createUserDomain.Password)

	if err != nil {
		return "", err
	}

	user := users.UserModel{
		FirstName: createUserDomain.FirstName,
		LastName:  createUserDomain.LastName,
		Email:     createUserDomain.Email,
		Password:  hashedPassword,
	}

	result, err := users.CreateUser(&user)

	if err != nil {
		return "", err

	}

	return result.String(), nil
}

func ValidateCredentials(email *string, password *string) (id string, err error) {

	credentials, err := users.GetUserCredentialsByEmail(email)

	if err != nil {
		return "", err
	}

	err = verifyPassword(*password, credentials.Password)

	if err != nil {
		return "", errors.New("invalid pasword")
	}

	token, err := authService.GenerateToken(credentials.ID.String())

	if err != nil {
		return "", err
	}

	return token, nil
}

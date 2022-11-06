package users

import (
	"chat-app/database"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CredentialsModel struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func CreateUser(user *UserModel) (uuid.UUID, error) {
	db := database.GetConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO users (first_name, last_name, email, password) values ($1, $2, $3, $4) RETURNING id;`

	var id uuid.UUID
	err := db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)

	return id, err

}

func GetUserCredentialsByEmail(email *string) (credentialsModel CredentialsModel, err error) {
	db := database.GetConnection()
	defer db.Close()

	sqlStatement := `select id, email, password from users u where email=$1`

	row := db.QueryRow(sqlStatement, email)
	err = row.Scan(&credentialsModel.ID, &credentialsModel.Email, &credentialsModel.Password)

	return credentialsModel, err
}

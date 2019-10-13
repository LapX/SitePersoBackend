package authentication

import (
	"github.com/LapX/SitePersoBackend/dependencies/database"
)

type UserEmailPicture struct {
	Email   string
	Picture string
}

func GetUser(token string) UserEmailPicture {
	email, picture := database.FetchUser(token)
	return UserEmailPicture{email, picture}
}

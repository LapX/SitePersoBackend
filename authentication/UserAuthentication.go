package authentication

import "github.com/LapX/SitePersoBackend/repository"

type UserEmailPicture struct {
	Email   string
	Picture string
}

func GetUser(token string) UserEmailPicture {
	email, picture := repository.FetchUser(token)
	return UserEmailPicture{email, picture}
}

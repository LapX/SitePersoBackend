package resource

import (
	"encoding/json"
	"github.com/LapX/SitePersoBackend/authentication"
	"net/http"
)

func oauthGoogleLogin(response http.ResponseWriter, request *http.Request) {
	user := authentication.LoginUser(response)
	http.Redirect(response, request, user, http.StatusTemporaryRedirect)
}

func oauthGoogleCallback(response http.ResponseWriter, request *http.Request) {
	token := authentication.LoginCallback(response, request)
	http.Redirect(response, request, "http://lapx.github.io/SitePersoFrontend/?token="+token, http.StatusTemporaryRedirect)
}

func getUserInfo(response http.ResponseWriter, request *http.Request) {
	token, err := request.URL.Query()["token"]

	if err {
		userEmailPicture := authentication.GetUser(token[0])
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(userEmailPicture)
	} else {
		json.NewEncoder(response).Encode("token query param missing")
	}
}

package oAuth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/LapX/SitePersoBackend/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//Code follows this tutorial : https://dev.to/douglasmakey/oauth2-example-with-go-3n8a

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Endpoint:     google.Endpoint,
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func OauthGoogleLogin(response http.ResponseWriter, request *http.Request) {
	oauthState := generateStateOauthCookie(response)
	user := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(response, request, user, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(response http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(response, &cookie)

	return state
}

func OauthGoogleCallback(response http.ResponseWriter, request *http.Request) {
	var userInfo repository.UserInfo
	oauthState, _ := request.Cookie("oauthstate")

	if request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(response, request, "http://localhost:3000/SitePersoFrontend/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(response, request, "http://localhost:3000/SitePersoFrontend/", http.StatusTemporaryRedirect)
		return
	}

	json.Unmarshal(data, &userInfo)
	repository.StoreUserInfo(userInfo)
	http.Redirect(response, request, "http://localhost:3000/SitePersoFrontend/", http.StatusTemporaryRedirect)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

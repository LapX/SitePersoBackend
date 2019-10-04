package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	Endpoint:     google.Endpoint,
	RedirectURL:  "https://lapx.github.io/SitePersoFrontend/",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthGoogleLogin(response http.ResponseWriter, request *http.Request) {
	clientId, clientSecret := getOauthInfo()
	googleOauthConfig.ClientID = clientId
	googleOauthConfig.ClientSecret = clientSecret

	oauthState := generateStateOauthCookie(response)
	user := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(response, request, user, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func oauthGoogleCallback(response http.ResponseWriter, request *http.Request) {
	oauthState, _ := request.Cookie("oauthstate")

	if request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(response, request, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(response, request, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(response, "UserInfo: %s\n", data)
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

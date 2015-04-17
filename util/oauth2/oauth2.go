// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package oauth2 contains Martini handlers to provide
// user login via an OAuth 2.0 backend.
package oauth2

import (
	"net/http"
	"net/url"

	"github.com/dorajistyle/goyangi/util/log"
	"golang.org/x/oauth2"
)

type AuthResponse struct {
	State        string `form:"state"`
	Code         string `form:"code"`
	Authuser     int    `form:"authuser"`
	NumSessions  int    `form:"num_sessions"`
	prompt       string `form:"prompt"`
	ImageName    string `form:"imageName"`
	SessionState string `form:"session_state"`
}

// Your credentials should be obtained from the Google
// Developer Console (https://console.developers.google.com).
func OauthURL(conf *oauth2.Config) string {
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.

	url := conf.AuthCodeURL("state")
	log.Debugf("\nVisit the URL for the auth dialog: %v\n", url)
	return url
}

func OauthRequest(url *url.URL, conf *oauth2.Config, authResponse AuthResponse) (*http.Response, *oauth2.Token, error) {
	var res *http.Response
	var req *http.Request
	var token *oauth2.Token
	var err error

	log.Info("Oauth Redirect performed.")
	// Handle the exchange code to initiate a transport.
	log.Debugf("AuthResponse :%v\n", authResponse)
	token, err = conf.Exchange(oauth2.NoContext, authResponse.Code)
	// log.Debugf("Token AccessToken :%s\n", token.AccessToken)
	// log.Debugf("Token TokenType :%s\n", token.TokenType)
	// log.Debugf("Token RefreshToken :%s\n", token.RefreshToken)
	// log.Debugf("Token Expiry :%s\n", token.Expiry)
	if err != nil {
		return res, token, err
	}
	// if err == nil {
	client := conf.Client(oauth2.NoContext, token)
	log.Debugf("url : %s\n", url.String())
	req, err = http.NewRequest("GET", "url.String()", nil)
	if err != nil {
		return res, token, err
	}
	// if err == nil {
	req.URL = url
	res, err = client.Do(req)
	// if err != nil {
	// 	return res, token, err
	// }
	return res, token, nil
	// if err == nil {
	// 	return res, token, err
	// }
	// }
	// }

	// return nil, token, err

}

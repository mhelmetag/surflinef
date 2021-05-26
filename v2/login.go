package surflinef

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type LoginQuery struct {
	IsShortLived bool `url:"isShortLived"`
}

type LoginPayload struct {
	AuthorizationString string `json:"authorizationString"`
	DeviseID            string `json:"device_id"`
	DeviseType          string `json:"device_type"`
	Forced              bool   `json:"forced"`
	GrantType           string `json:"grant_type"`
	Password            string `json:"password"`
	Username            string `json:"username"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type ErrorLoginResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func DefaultLoginPayload(username string, password string) LoginPayload {
	return LoginPayload{
		AuthorizationString: "Basic NWM1OWU3YzNmMGI2Y2IxYWQwMmJhZjY2OnNrX1FxWEpkbjZOeTVzTVJ1MjdBbWcz",
		DeviseID:            "",
		DeviseType:          "",
		Forced:              true,
		GrantType:           "password",
		Password:            password,
		Username:            username,
	}
}

func (c *Client) PostLogin(lq LoginQuery, lp LoginPayload) (LoginResponse, error) {
	vs, err := query.Values(lq)

	if err != nil {
		return LoginResponse{}, err
	}

	qs := vs.Encode()

	s := c.FullURL(qs)
	u, err := url.Parse(s)
	if err != nil {
		return LoginResponse{}, err
	}

	p, err := json.Marshal(lp)
	if err != nil {
		return LoginResponse{}, err
	}

	pr := bytes.NewReader(p)

	r, err := http.Post(u.String(), "application/json", pr)
	if err != nil {
		return LoginResponse{}, err
	}

	defer r.Body.Close()

	if r.StatusCode == 200 {
		var lr LoginResponse
		err = json.NewDecoder(r.Body).Decode(&lr)
		if err != nil {
			return LoginResponse{}, err
		}

		return lr, nil
	} else {
		var elr ErrorLoginResponse
		err = json.NewDecoder(r.Body).Decode(&elr)
		if err != nil {
			return LoginResponse{}, err
		}

		err = errors.New(elr.ErrorDescription)

		return LoginResponse{}, err
	}
}

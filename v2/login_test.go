package surflinef

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLogin(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/login.json")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	c := Client{BaseURL: bu}
	lq := LoginQuery{
		IsShortLived: false,
	}
	lp := DefaultLoginPayload("fred.frederickson@example.com", "Password")

	lr, err := c.PostLogin(lq, lp)
	if err != nil {
		t.Fatal(err)
	}

	eT := "43ab1foo4ca3d74b59dd96fd8a82a135f914d1b2"
	aT := lr.AccessToken
	if aT != eT {
		t.Errorf("Got '%s', expected '%s'", aT, eT)
	}
}

func TestLoginError(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/login-error.json")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}))

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	c := Client{BaseURL: bu}
	lq := LoginQuery{
		IsShortLived: false,
	}
	lp := DefaultLoginPayload("fred.frederickson@example.com", "WrongPassword")

	_, err = c.PostLogin(lq, lp)
	if err == nil {
		t.Fatalf("Login should've returned an error but didn't")
	}

	eM := "Invalid email and password combination"
	aM := err.Error()
	if aM != eM {
		t.Errorf("Got '%s', expected '%s'", aM, eM)
	}
}

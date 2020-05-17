package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// FacebookProvider godoc
type FacebookProvider struct {
	Config oauth2.Config
}

// NewFacebookProvider godoc
func NewFacebookProvider(clientID string, clientSecret string, redirectURL string) *FacebookProvider {
	return &FacebookProvider{
		Config: oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"public_profile"},
			Endpoint:     facebook.Endpoint,
		},
	}
}

// BuildURL godoc
func (p *FacebookProvider) BuildURL() string {
	authURL, err := url.Parse(p.Config.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", p.Config.ClientID)
	parameters.Add("scope", strings.Join(p.Config.Scopes, " "))
	parameters.Add("redirect_uri", p.Config.RedirectURL)
	parameters.Add("response_type", "code")
	authURL.RawQuery = parameters.Encode()
	return authURL.String()
}

// GetProfile godoc
func (p *FacebookProvider) GetProfile(code string) (map[string]string, error) {
	token, err := p.Config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	json.Unmarshal(response, &m)
	m["picture"] = "http://graph.facebook.com/" + m["id"] + "/picture?type=square"
	return m, nil
}

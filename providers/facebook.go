package providers

import (
	"log"
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
	oauthStateString := "thisshouldbearandomstring"
	authURL, err := url.Parse(p.Config.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", p.Config.ClientID)
	parameters.Add("scope", strings.Join(p.Config.Scopes, " "))
	parameters.Add("redirect_uri", p.Config.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	authURL.RawQuery = parameters.Encode()
	return authURL.String()
}

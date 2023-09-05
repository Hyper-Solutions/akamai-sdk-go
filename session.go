package akamai

import "net/http"

type Session struct {
	apiKey, jwtKey string
	client         *http.Client
}

// NewSession creates a new Session that can be used to make requests to the Hyper Solutions API.
func NewSession(apiKey string) *Session {
	return &Session{
		apiKey: apiKey,
		client: http.DefaultClient,
	}
}

// WithJwtKey adds the JWT Key to the session. If not empty, a signature will be added to each request.
func (s *Session) WithJwtKey(jwt string) *Session {
	s.jwtKey = jwt
	return s
}

// WithClient sets a new client that will be used to make requests to the Hyper Solutions API.
func (s *Session) WithClient(client *http.Client) *Session {
	s.client = client
	return s
}

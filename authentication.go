package harvest

import (
	"encoding/base64"
	"fmt"
)

type AuthenticationService struct {
	client *Client
}

type Session struct {
	Authorization string
}

func (s *AuthenticationService) Acquire(username, password string) (bool, error) {
	apiEndpoint := "account/who_am_i"

	authorization := encodeString(username, password)
	authorization = fmt.Sprintf("Basic %s", authorization)
	session := new(Session)
	session.Authorization = authorization

	s.client.session = session
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		s.client.session = nil
		return false, err
	}

	resp, err := s.client.Do(req, nil)

	if err != nil {
		s.client.session = nil
		return false, fmt.Errorf("Auth at Harvest instance failed (HTTP(S) request). %s", err)
	}
	if resp != nil && resp.StatusCode != 200 {
		s.client.session = nil
		return false, fmt.Errorf("Auth at Harvest instance failed (HTTP(S) request). Status code: %d", resp.StatusCode)
	}

	return true, nil
}

func encodeString(user string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, password)))

}

// Authenticated reports if the current Client has an authenticated session with Harvest
func (s *AuthenticationService) Authenticated() bool {
	if s != nil {
		return s.client.session != nil
	}
	return false
}

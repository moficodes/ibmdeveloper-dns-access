package ibmcloud

import (
	"log"
	"time"
)

var endpoints *IdentityEndpoints

func cacheIdentityEndpoints() error {
	if endpoints == nil {
		var err error
		endpoints, err = getIdentityEndpoints()
		if err != nil {
			return err
		}
	}
	return nil
}

// GetIdentityEndpoints returns the list of endpoints for IBMCloud IAM
func GetIdentityEndpoints() (*IdentityEndpoints, error) {
	err := cacheIdentityEndpoints()
	if err != nil {
		return nil, err
	}
	return endpoints, nil
}

// IAMAuthenticate uses the api key to authenticate and return a Session
func IAMAuthenticate(apikey string) (*Session, error) {
	err := cacheIdentityEndpoints()
	if err != nil {
		log.Println("error with cached data")
		return nil, err
	}
	token, err := getTokenFromIAM(endpoints.TokenEndpoint, apikey)

	if err != nil {
		log.Println("error with token data")
		return nil, err
	}
	return &Session{Token: token}, nil
}

// Authenticate uses the one time passcode to authenticate and return a Session
func Authenticate(otp string) (*Session, error) {
	err := cacheIdentityEndpoints()
	if err != nil {
		log.Println("error with cached data")
		return nil, err
	}
	token, err := getToken(endpoints.TokenEndpoint, otp)

	if err != nil {
		log.Println("error with token data")
		return nil, err
	}
	return &Session{Token: token}, nil
}

// GetAccounts get a list of accounts for the current session
func (s *Session) GetAccounts() (*Accounts, error) {
	return s.getAccountsWithEndpoint(nil)
}

// IsValid checks if session is expired or not
func (s *Session) IsValid() bool {
	now := time.Now().Unix()
	difference := int64(s.Token.Expiration) - now
	return difference > 100 // expires in 3600 second. keeping 100 second buffer
}

func (s *Session) getAccountsWithEndpoint(nextURL *string) (*Accounts, error) {
	if !s.IsValid() {
		log.Println("Access token expired.")
		token, err := upgradeToken(endpoints.TokenEndpoint, s.Token.RefreshToken, "")
		if err != nil {
			return nil, err
		}
		log.Println("Token Refreshed.")
		s.Token = token
	}
	accounts, err := getAccounts(nextURL, s.Token.AccessToken)
	if err != nil {
		return nil, err
	}
	if accounts.NextURL != nil {
		nextAccounts, err := s.getAccountsWithEndpoint(accounts.NextURL)
		if err != nil {
			return nil, err
		}
		nextAccounts.Resources = append(nextAccounts.Resources, accounts.Resources...)
		return nextAccounts, nil
	}
	return accounts, nil
}

// GetAccountResources return AccountResources
func (s *Session) GetAccountResources(accountID string) (*AccountResources, error) {
	return getAccountResources(s.Token.AccessToken, accountID)
}

func (s *Session) GetUserInfo() (*UserInfo, error) {
	err := cacheIdentityEndpoints()
	if err != nil {
		return nil, err
	}
	if !s.IsValid() {
		token, err := upgradeToken(endpoints.TokenEndpoint, s.Token.RefreshToken, "")
		if err != nil {
			return nil, err
		}
		log.Println("Token Refreshed.")
		s.Token = token
	}

	return getUserInfo(endpoints.UserinfoEndpoint, s.Token.AccessToken)
}

func (s *Session) GetUserPreference(accountID, userID string) (*User, error) {
	if !s.IsValid() {
		token, err := upgradeToken(endpoints.TokenEndpoint, s.Token.RefreshToken, "")
		if err != nil {
			return nil, err
		}
		log.Println("Token Refreshed.")
		s.Token = token
	}
	return getUserPreference(accountID, userID, s.Token.AccessToken)
}

func bindAccountToToken(refreshToken, accountID string) (*Session, error) {
	err := cacheIdentityEndpoints()
	if err != nil {
		return nil, err
	}
	token, err := upgradeToken(endpoints.TokenEndpoint, refreshToken, accountID)
	if err != nil {
		return nil, err
	}
	return &Session{Token: token}, nil
}

// BindAccountToToken upgrades session with account
func (s *Session) BindAccountToToken(accountID string) (*Session, error) {
	session, err := bindAccountToToken(s.Token.RefreshToken, accountID)
	if err != nil {
		return nil, err
	}
	return session, err
}

// RenewSession renews session with refresh token
func (s *Session) RenewSession() (*Session, error) {
	err := cacheIdentityEndpoints()
	if err != nil {
		return nil, err
	}
	token, err := upgradeToken(endpoints.TokenEndpoint, s.Token.RefreshToken, "")
	if err != nil {
		return nil, err
	}
	return &Session{Token: token}, nil
}

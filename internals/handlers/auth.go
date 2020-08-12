package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/moficodes/ibmdeveloper-domain/pkg/ibmcloud"
)

const (
	sessionName  = "cloud_session"
	accessToken  = "access_token"
	refreshToken = "refresh_token"
	expiration   = "expiration"
	cookiePath   = "/api"
)

func AuthenticationHandler(c echo.Context) error {
	accountLogin := new(AccountLogin)
	if err := c.Bind(accountLogin); err != nil {
		return err
	}

	session, err := ibmcloud.Authenticate(accountLogin.OTP)
	if err != nil {
		log.Println("could not authenticate with the otp provided")
		log.Println(err.Error())
		return err
	}

	setCookie(c, session)

	return c.JSON(http.StatusOK, StatusOK{Message: "success"})
}

func LoginHandler(c echo.Context) error {
	session, err := getCloudSessions(c)
	if err != nil {
		return err
	}

	if !session.IsValid() {
		return err
	}

	return c.JSON(http.StatusOK, StatusOK{Message: "success"})
}

func TokenEndpointHandler(c echo.Context) error {
	endpoints, err := ibmcloud.GetIdentityEndpoints()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, endpoints)
}

func setCookie(c echo.Context, session *ibmcloud.Session) {
	accessTokenCookie := &http.Cookie{Name: accessToken, Value: session.Token.AccessToken, Path: cookiePath}
	c.SetCookie(accessTokenCookie)

	refreshTokenCookie := &http.Cookie{Name: refreshToken, Value: session.Token.RefreshToken, Path: cookiePath}
	c.SetCookie(refreshTokenCookie)

	expirationStr := strconv.Itoa(session.Token.Expiration)

	expirationCookie := &http.Cookie{Name: expiration, Value: expirationStr, Path: cookiePath}
	c.SetCookie(expirationCookie)
}

func getCloudSessions(c echo.Context) (*ibmcloud.Session, error) {
	var accessToken string
	var refreshToken string
	var expirationTime int
	accessTokenVal, err := c.Cookie(accessToken)
	if err != nil {
		bearerToken := c.Request().Header.Get("Authorization")
		if bearerToken == "" {
			return nil, err
		}
		parsedToken := strings.Split(bearerToken, " ")
		if len(parsedToken) != 2 {
			return nil, err
		}
		accessToken = parsedToken[1]
	} else {
		accessToken = accessTokenVal.Value
	}

	refreshTokenVal, err := c.Cookie(refreshToken)
	if err != nil {
		refreshToken = c.Request().Header.Get("X-Auth-Refresh-Token")
		if refreshToken == "" {
			return nil, err
		}
	} else {
		refreshToken = refreshTokenVal.Value
	}

	expirationValStr, err := c.Cookie(expiration)
	if err != nil {
		expirationTime = 0
	} else {
		expirationTime, err = strconv.Atoi(expirationValStr.Value)
		if err != nil {
			return nil, err
		}
	}

	session := &ibmcloud.Session{
		Token: &ibmcloud.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Expiration:   expirationTime,
		},
	}

	return session.RenewSession()
}

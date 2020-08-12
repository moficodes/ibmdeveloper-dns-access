package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserInfoHandler(c echo.Context) error {
	session, err := getCloudSessions(c)

	if err != nil {
		log.Println(err)
		return err
	}

	userInfo, err := session.GetUserInfo()
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, userInfo)
}

func UserPreferenceHandler(c echo.Context) error {
	session, err := getCloudSessions(c)

	accountID := c.Param("accountID")
	userID := c.Param("userID")

	if err != nil {
		log.Println(err)
		return err
	}

	userPreference, err := session.GetUserPreference(accountID, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, userPreference)
}
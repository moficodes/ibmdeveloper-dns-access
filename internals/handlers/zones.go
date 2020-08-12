package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moficodes/ibmdeveloper-domain/pkg/dns"
)

// ListZoneHandler returns a JSON list of zones
func ListZoneHandler(c echo.Context) error {
	client, err := dns.GetDNSimpleClient()
	if err != nil {
		return err
	}

	accountID, err := dns.GetAccountID(client)
	if err != nil {
		return err
	}

	zones, err := dns.GetZones(client, *accountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, zones)
}

func ZoneRecordsHandler(c echo.Context) error {
	client, err := dns.GetDNSimpleClient()
	if err != nil {
		return err
	}

	accountID, err := dns.GetAccountID(client)
	if err != nil {
		return err
	}

	zone := c.Param("zone")

	records, err := dns.GetRecords(client, *accountID, zone)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, records)
}

func CreateNewRecord(c echo.Context) error {
	session, err := getCloudSessions(c)
	if err != nil {
		return err
	}

	session, err = session.RenewSession()
	if err != nil {
		return err
	}

	client, err := dns.GetDNSimpleClient()
	if err != nil {
		return err
	}

	accountID, err := dns.GetAccountID(client)
	if err != nil {
		return err
	}

	createRecord := new(CreateRecord)

	if err := c.Bind(createRecord); err != nil {
		return err
	}

	_, err = dns.CreateRecord(client, *accountID, createRecord.Zone,
		createRecord.ZoneID, createRecord.RecordType,
		createRecord.RecordZone, createRecord.Content, createRecord.TTL)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, StatusOK{Message: "success"})
}

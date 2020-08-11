package dns

import (
	"context"
	"strings"

	"github.com/dnsimple/dnsimple-go/dnsimple"
)

// GetZones returns list of zone for an account
func GetZones(client *dnsimple.Client, accountID string) ([]dnsimple.Zone, error) {
	zones, err := client.Zones.ListZones(context.Background(), accountID, nil)
	if err != nil {
		return nil, err
	}
	res := make([]dnsimple.Zone, 0)
	for _, zone := range zones.Data {
		if strings.Contains(zone.Name, "ibmdeveloper") {
			res = append(res, zone)
		}
	}

	return res, nil
}

// GetRecords returns the list of zone records for a zone
func GetRecords(client *dnsimple.Client, accountID, zone string) ([]dnsimple.ZoneRecord, error) {
	zoneRecords, err := client.Zones.ListRecords(context.Background(), accountID, zone, nil)
	if err != nil {
		return nil, err
	}

	return zoneRecords.Data, nil
}

// CreateRecord creates a new record at the specified zone
func CreateRecord(client *dnsimple.Client, accountID, zone, zoneID, recordType, recordZone, content string, ttl int) (*dnsimple.ZoneRecordResponse, error) {
	attribute := dnsimple.ZoneRecordAttributes{
		Name:    &recordZone,
		ZoneID:  zoneID,
		Content: content,
		Type:    recordType,
		TTL:     ttl,
	}
	resp, err := client.Zones.CreateRecord(context.Background(), accountID, zone, attribute)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

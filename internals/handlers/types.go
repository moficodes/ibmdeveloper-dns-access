package handlers

type CreateRecord struct {
	Zone       string `json:"zone"`
	ZoneID     string `json:"zoneID"`
	RecordType string `json:"recordType"`
	RecordZone string `json:"recordZone"`
	Content    string `json:"content"`
	TTL        int    `json:"ttl"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

//zone, zoneID, recordType, recordZone, content string, ttl int

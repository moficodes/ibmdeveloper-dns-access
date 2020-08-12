package handlers

type CreateRecord struct {
	Zone       string `json:"zone"`
	ZoneID     string `json:"zoneID"`
	RecordType string `json:"recordType"`
	RecordZone string `json:"recordZone"`
	Content    string `json:"content"`
	TTL        int    `json:"ttl"`
}

type StatusOK struct {
	Message string `json:"message"`
}

//IBMCloud handler related types

// AccountLogin ...
type AccountLogin struct {
	OTP string `json:"otp"`
}

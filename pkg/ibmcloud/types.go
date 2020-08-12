package ibmcloud

import "time"

type Session struct {
	Token *Token
}

// Token is used in every request that need authentication
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ImsUserID    int    `json:"ims_user_id"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Expiration   int    `json:"expiration"`
	Scope        string `json:"scope"`
}

type IdentityEndpoints struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	PasscodeEndpoint                  string   `json:"passcode_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
}

//ACCOUNT RELATED TYPES
type Accounts struct {
	NextURL      *string   `json:"next_url"`
	TotalResults int       `json:"total_results"`
	Resources    []Account `json:"resources"`
}
type Metadata struct {
	GUID      string    `json:"guid"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type TermsAndConditions struct {
	Required  bool      `json:"required"`
	Accepted  bool      `json:"accepted"`
	Timestamp time.Time `json:"timestamp"`
}
type OrganizationsRegion struct {
	GUID   string `json:"guid"`
	Region string `json:"region"`
}
type Linkages struct {
	Origin string `json:"origin"`
	State  string `json:"state"`
}
type PaymentMethod struct {
	Type           string      `json:"type"`
	Started        time.Time   `json:"started"`
	Ended          string      `json:"ended"`
	CurrencyCode   string      `json:"currencyCode"`
	AnniversaryDay interface{} `json:"anniversaryDay"`
}
type History struct {
	Type               string    `json:"type"`
	State              string    `json:"state"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	CurrencyCode       string    `json:"currencyCode"`
	CountryCode        string    `json:"countryCode"`
	BillingCountryCode string    `json:"billingCountryCode"`
	BillingSystem      string    `json:"billingSystem"`
}
type BluemixSubscriptions struct {
	Type                  string        `json:"type"`
	State                 string        `json:"state"`
	PaymentMethod         PaymentMethod `json:"payment_method"`
	SubscriptionID        string        `json:"subscription_id"`
	PartNumber            string        `json:"part_number"`
	SubscriptionTags      []interface{} `json:"subscriptionTags"`
	PaygPendingTimestamp  time.Time     `json:"payg_pending_timestamp"`
	History               []History     `json:"history"`
	CurrentStateTimestamp time.Time     `json:"current_state_timestamp"`
	SoftlayerAccountID    string        `json:"softlayer_account_id"`
	BillingSystem         string        `json:"billing_system"`
}

type Resource struct {
	ResourceID string `json:"resource_id"`
}

type Results struct {
	ResourceID string      `json:"resource_id"`
	IsError    string      `json:"isError"`
	ISError    interface{} `json:"is_error"`
}

type ResourceGroups struct {
	Resources []Group `json:"resources"`
}
type Group struct {
	ID                string        `json:"id"`
	Crn               string        `json:"crn"`
	AccountID         string        `json:"account_id"`
	Name              string        `json:"name"`
	State             string        `json:"state"`
	Default           bool          `json:"default"`
	EnableReclamation bool          `json:"enable_reclamation"`
	QuotaID           string        `json:"quota_id"`
	QuotaURL          string        `json:"quota_url"`
	PaymentMethodsURL string        `json:"payment_methods_url"`
	ResourceLinkages  []interface{} `json:"resource_linkages"`
	TeamsURL          string        `json:"teams_url"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

// USER RELATED TYPES
type Users struct {
	TotalResults int    `json:"total_results"`
	Limit        int    `json:"limit"`
	FirstURL     string `json:"first_url"`
	NextURL      string `json:"next_url"`
	Resources    []User `json:"resources"`
}
type User struct {
	ID             string `json:"id"`
	IamID          string `json:"iam_id"`
	Realm          string `json:"realm"`
	UserID         string `json:"user_id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	State          string `json:"state"`
	Email          string `json:"email"`
	Phonenumber    string `json:"phonenumber"`
	Altphonenumber string `json:"altphonenumber"`
	Photo          string `json:"photo"`
	AccountID      string `json:"account_id"`
}

type UserInfo struct {
	Active     bool     `json:"active"`
	RealmID    string   `json:"realmId"`
	Identifier string   `json:"identifier"`
	IamID      string   `json:"iam_id"`
	GivenName  string   `json:"given_name"`
	FamilyName string   `json:"family_name"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Sub        string   `json:"sub"`
	Account    Account  `json:"account"`
	Iat        int      `json:"iat"`
	Exp        int      `json:"exp"`
	Iss        string   `json:"iss"`
	GrantType  string   `json:"grant_type"`
	ClientID   string   `json:"client_id"`
	Scope      string   `json:"scope"`
	Acr        int      `json:"acr"`
	Amr        []string `json:"amr"`
}
type Account struct {
	Bss       string `json:"bss"`
	Ims       string `json:"ims"`
	ImsUserID string `json:"ims_user_id"`
	Valid     bool   `json:"valid"`
}

type ErrorMessage struct {
	ErrorDescription string  `json:"error_description"`
	Trace            string  `json:"trace"`
	Error            []Error `json:"error"`
	Errors           []Error `json:"errors"`
}

type Target struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
type Error struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Target   Target `json:"target"`
}

type URL struct {
	Href   string `json:"href"`
	Offset string `json:"offset"`
}

type AccountResources struct {
	Resources []AccountResource `json:"resources"`
}
type AccountResource struct {
	ID                string        `json:"id"`
	Crn               string        `json:"crn"`
	AccountID         string        `json:"account_id"`
	Name              string        `json:"name"`
	State             string        `json:"state"`
	Default           bool          `json:"default"`
	EnableReclamation bool          `json:"enable_reclamation"`
	QuotaID           string        `json:"quota_id"`
	QuotaURL          string        `json:"quota_url"`
	PaymentMethodsURL string        `json:"payment_methods_url"`
	ResourceLinkages  []interface{} `json:"resource_linkages"`
	TeamsURL          string        `json:"teams_url"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type ErrorItems struct {
	Code             string `json:"code"`
	Description      string `json:"description"`
	RecoveryCLI      string `json:"recoveryCLI"`
	RecoveryUI       string `json:"recoveryUI"`
	TerseDescription string `json:"terseDescription"`
	Type             string `json:"type"`
}
type NonCriticalErrors struct {
	IncidentID string       `json:"incidentID"`
	Items      []ErrorItems `json:"items"`
}

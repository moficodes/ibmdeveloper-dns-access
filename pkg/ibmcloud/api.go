package ibmcloud

// TODO: return errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// protocol
const protocol = "https://"

// subdomains
const (
	subdomainIAM                = "iam."
	subdomainUserManagement     = "user-management."
	subdomainAccounts           = "accounts."
	subdomainResourceController = "resource-controller."
	subdomainClusters           = "containers."
	subdomainUsers              = "users."
	subdomainTags               = "tags.global-search-tagging."
	subdomainBilling            = "billing."
)

// domain
const api = "cloud.ibm.com"

// endpoints
const (
	identityEndpoint       = protocol + subdomainIAM + api + "/identity/.well-known/openid-configuration"
	userPreferenceEndpoint = protocol + subdomainUserManagement + api + "/v2/accounts"
	accountsEndpoint       = protocol + subdomainAccounts + api + "/coe/v2/accounts"
	resourcesEndpoint      = protocol + subdomainResourceController + api + "/v2/resource_instances"
	resourceKeysEndpoint   = protocol + subdomainResourceController + api + "/v2/resource_keys"
	containersEndpoint     = protocol + subdomainClusters + api + "/global/v1"
	usersEndpoint          = protocol + subdomainUsers + api + "/v2"
	tagEndpoint            = protocol + subdomainTags + api + "/v3/tags"
	billingEndpoint        = protocol + subdomainBilling + api + "/v4/accounts"
	resourceEndoint        = protocol + subdomainResourceController + api + "/v1/resource_groups"
)

const (
	clusterEndpoint     = containersEndpoint + "/clusters"
	versionEndpount     = containersEndpoint + "/versions"
	locationEndpoint    = containersEndpoint + "/locations"
	zonesEndpoint       = containersEndpoint + "/zones"
	datacentersEndpoint = containersEndpoint + "/datacenters"
)

// grant types
const (
	passcodeGrantType     = "urn:ibm:params:oauth:grant-type:passcode"
	apikeyGrantType       = "urn:ibm:params:oauth:grant-type:apikey"
	refreshTokenGrantType = "refresh_token"
)

const basicAuth = "Basic Yng6Yng="

//// useful for loagging
// bodyBytes, err := ioutil.ReadAll(resp.Body)
// if err != nil {
// 	panic(err)
// }
// bodyString := string(bodyBytes)
// log.Println(bodyString)
////

func timeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s took %s\n", name, elapsed)
}

func getError(resp *http.Response) error {
	var errorTemplate ErrorMessage
	if err := json.NewDecoder(resp.Body).Decode(&errorTemplate); err != nil {
		return err
	}
	if errorTemplate.Error != nil {
		return errors.New(errorTemplate.Error[0].Message)
	}
	if errorTemplate.Errors != nil {
		return errors.New(errorTemplate.Errors[0].Message)
	}
	return errors.New("unknown")
}

func getIdentityEndpoints() (*IdentityEndpoints, error) {
	result := &IdentityEndpoints{}
	err := fetch(identityEndpoint, nil, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getToken(endpoint string, otp string) (*Token, error) {
	header := map[string]string{
		"Authorization": basicAuth,
	}

	form := url.Values{}
	form.Add("grant_type", passcodeGrantType)
	form.Add("passcode", otp)

	result := Token{}
	err := postForm(endpoint, header, nil, form, &result)

	if err != nil {
		log.Println("error in post form")
		return nil, err
	}

	return &result, nil
}

func getTokenFromIAM(endpoint string, apikey string) (*Token, error) {
	header := map[string]string{
		"Authorization": basicAuth,
	}

	form := url.Values{}
	form.Add("grant_type", apikeyGrantType)
	form.Add("apikey", apikey)

	result := &Token{}
	err := postForm(endpoint, header, nil, form, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func upgradeToken(endpoint string, refreshToken string, accountID string) (*Token, error) {
	header := map[string]string{
		"Authorization": basicAuth,
	}

	form := url.Values{}
	form.Add("grant_type", refreshTokenGrantType)
	form.Add("refresh_token", refreshToken)
	if accountID != "" {
		form.Add("bss_account", accountID)
	}

	result := &Token{}
	err := postForm(endpoint, header, nil, form, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getUserInfo(endpoint string, token string) (*UserInfo, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint cannot be empty")
	}
	header := map[string]string{
		"Authorization": "Bearer " + token,
	}
	var result UserInfo
	err := fetch(endpoint, header, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func getUserPreference(accountID, userID, token string) (*User, error) {
	endpoint := fmt.Sprintf("%s/%s/users/%s", userPreferenceEndpoint, accountID, userID)

	header := map[string]string{
		"Authorization": "Bearer " + token,
	}

	var result User
	err := fetch(endpoint, header, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func getAccounts(endpoint *string, token string) (*Accounts, error) {
	if endpoint == nil {
		endpointString := accountsEndpoint
		endpoint = &endpointString
	} else {
		endpointString := accountsEndpoint + *endpoint
		endpoint = &endpointString
	}

	header := map[string]string{
		"Authorization": "Bearer " + token,
	}
	var result Accounts
	err := fetch(*endpoint, header, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func getAccountResources(token, accountID string) (*AccountResources, error) {
	var result AccountResources
	header := map[string]string{
		"Authorization": "Bearer " + token,
	}

	query := map[string]string{
		"account_id": accountID,
	}

	err := fetch(resourceEndoint, header, query, &result)
	if err != nil {
		return nil, err
	}
	//"/v1/resource_groups?account_id=9b13b857a32341b7167255de717172f5"
	return &result, nil
}

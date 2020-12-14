package models

//Request holds the data sent by the user
type Request struct {
	IP               string
	CountryWhiteList []string
	RequestID        string
}

//Response is the data sent back to the user
type Response struct {
	IPOnWhiteList     bool     `json:"givenIpIsOnWhiteList"`
	GivenIP           string   `json:"givenIp"`
	CountryForGivenIP string   `json:"countryForGivenIp"`
	WhiteListGiven    []string `json:"whiteListGiven"`
}

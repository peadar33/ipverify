package verify

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/oschwald/geoip2-golang"
	"github.com/peadar33/ipverify/dbconnection"
	"github.com/peadar33/ipverify/models"
	"github.com/sirupsen/logrus"
)

var db *geoip2.Reader

func init() {
	db = dbconnection.GetDbclient().Client
}

//CheckTheWhitelist ... CheckTheWhitelist
func CheckTheWhitelist(req models.Request) (models.Response, error) {

	var wg sync.WaitGroup
	wg.Add(len(req.CountryWhiteList))

	//Start building the response
	response := models.Response{IPOnWhiteList: false, GivenIP: req.IP, WhiteListGiven: req.CountryWhiteList}

	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(req.IP)
	if ip == nil {
		fmt.Printf("IP: %s\n", ip.String())
		err := errors.New("IP did not parse")
		return response, err
	}

	country, err := getTheIPCountry(ip)
	if err != nil {
		logrus.Error(err)
	}

	response.CountryForGivenIP = country
	//A go routine might be overkill here.
	for i := 0; i < len(req.CountryWhiteList); i++ {
		go func(i int) {
			defer wg.Done()
			if country == req.CountryWhiteList[i] {
				logrus.Infof("RequestID: %s. IP : %s is from %s", req.RequestID, ip.String(), country)
				logrus.Infof("RequestID: %s. Whitelist Match Found: %s = %s", req.RequestID, country, req.CountryWhiteList[i])
				response.IPOnWhiteList = true
				//break
			}
		}(i)
	}

	wg.Wait()

	if !response.IPOnWhiteList {
		logrus.Infof("RequestID: %s. IP : %s is from %s. No match found on whitelist.", req.RequestID, ip.String(), country)
	}

	return response, err
}

func getTheIPCountry(ip net.IP) (string, error) {
	record, err := db.Country(ip)
	if err != nil {
		return "no_country_found", err
	}
	country := ""
	if len(record.Country.Names["en"]) == 0 {
		country = "no_country_found"
	} else {
		country = record.Country.Names["en"]
	}

	return country, nil
}

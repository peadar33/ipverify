package dbconnection

import (
	"log"

	"github.com/oschwald/geoip2-golang"
)

//DB ... DB
type DB struct {
	Client *geoip2.Reader
}

//GetDbclient ... GetDbclient
func GetDbclient() *DB {
	//Open the country ip list database
	db, err := geoip2.Open("./data/GeoLite2-Country_20201208/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()

	return &DB{
		Client: db,
	}
	//return db
}

//Close ... Close
func (d *DB) Close() error {
	return d.Client.Close()
}

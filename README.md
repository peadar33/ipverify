# ipverify
Verify country for IP provided is or isn't on the whitelist provided.


## Access

GET http://localhost:10000/api/v1/whitelist

Authorization: Basic Auth

Credentials: username: abc pwd: 123


## Payload

```
Is on whitelist
{
    "ip": "24.98.94.55",
    "countryWhiteList": ["Ireland", 
    "Canada", 
    "United States", 
    "United Kingdom"]
}
Is not on whitelist
{
    "ip": "24.98.94.55",
    "countryWhiteList": ["Ireland", 
    "Canada", 
    "United States", 
    "United Kingdom"]
}
bad IP
{
    "ip": "0.98.94.55",
    "countryWhiteList": ["Ireland", 
    "Canada", 
    "United States", 
    "United Kingdom"]
}
```

## Response
```
{
    "givenIpIsOnWhiteList": true,
    "givenIp": "24.98.94.55",
    "countryForGivenIp": "United States",
    "whiteListGiven": [
        "Ireland",
        "Canada",
        "United States",
        "United Kingdom"
    ]
}
```
## Logging
Logrus is used to handle the logs. In a production environment the logs would be forwared by an agent to something like Splunk or Elasticsearch

## Kepping the database fresh
Option 1. Currently the data is stored locally in GeoLite2-Country.mmdb The data will be refreshed once per week. A cron trigger will excute a function that will download the latest GeoLite2-Country.mmdb from https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=<Key>&suffix=tar.gz and replace the old GeoLite2-Country.mmdb.
 
Options 2. Store the GeoLite2-Country data in a cache service such as Redis. This will make the data available to other services that might need it. Update the cache similar to option 1. 

## Future Enhancements
Add the option for the user to specify the language they want to search in. Currently once a country is found for a given IP the service only uses the english name for the country to search for a match in the whitelist, in some cases the whitelist provided may be in a different language, e'g' United States and (Spanish)Estados Unidos. It would be good if the service could handle a whitelist of any language.......or maybe it should do this by default.

## License
[MIT](https://choosealicense.com/licenses/mit/)



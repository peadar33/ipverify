package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/peadar33/ipverify/verify"

	logging "github.com/peadar33/ipverify/logging"
	"github.com/peadar33/ipverify/models"
	"github.com/sirupsen/logrus"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/peadar33/ipverify/security"
)

func init() {
	logging.CreateLog()
}

func convertReqBodyToRequestStruct(reqBody []byte, request models.Request) (models.Request, error) {
	var err error
	if len(reqBody) == 0 {
		err = errors.New("Request body is empty")
		fmt.Println(err)
		logrus.Warning(err.Error())
		return request, err
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		logrus.Warningf("Unmarshal failed. Error: %s", err.Error())
		return request, err
	}

	logrus.Infof("RequestID: %s. Request to struct successful", fmt.Sprintf("%s", request.RequestID))
	return request, nil
}

//Handlers Start **********************
func isAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all good")
}

func verifyIPHandler(w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	//create a new requestID to track the request
	requestID, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Something went wrong while creating the requestID: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Oops, something happened on our side. Please try again.")
		return
	}
	request.RequestID = fmt.Sprintf("%s", requestID)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Hi, heads up, there is something up with the request body. Please review the documentation for required fields.")
		return
	}

	//Build a request struct out of the request body data
	request, err = convertReqBodyToRequestStruct(reqBody, request)
	//Let the user know if there is and issue with the request body.
	if err != nil || len(request.IP) <= 0 || len(request.CountryWhiteList) <= 0 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Hi, heads up, there is something up with the request body. Please review the documentation for required fields.")
		return
	}

	//Verify the IP
	response, err := verify.CheckTheWhitelist(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Oops, something happened on our side. Please try again.")
		return
	}
	w.Header().Set("Content-type", "applciation/json")
	if !response.IPOnWhiteList {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(response)
}

//Handlers End **********************

func handleRequests() {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", isAlive).Methods("GET")
	router.HandleFunc("/api/v1/whitelist", security.BasicAuthMiddleware(verifyIPHandler))

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	logrus.Info("API starting...")
	handleRequests()
}

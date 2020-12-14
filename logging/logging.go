package logging

import (
	"fmt"
	"os"

	logrus "github.com/sirupsen/logrus"
)

//CreateLog ... CreateLog
func CreateLog() {
	// open a file
	f, err := os.OpenFile("verifyIPApi.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	//defer f.Close()

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(f)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

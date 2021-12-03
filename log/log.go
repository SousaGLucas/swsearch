package log

// responsible for logging system errors

import "log"

func SetLog(err error) {
	log.Fatal(err.Error())
}

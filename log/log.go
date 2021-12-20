package log

// package responsible for logging system errors

import "log"

func SetLog(err error) {
	log.Fatal(err.Error())
}

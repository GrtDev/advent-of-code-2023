package utils

import (
	"log"
)

func FatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

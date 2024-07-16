package helpers

import "log"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalError(str string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", str, err)
	}
}

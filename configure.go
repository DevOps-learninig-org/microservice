package microservice

import (
  "reflect"
)

const envPrefix string = "MYSERVICE"

type Config struct {
	LISTEN_ADDERSS string
	LISTEN_PORT string

	STORAGE_ADDRESS string
	STORAGE_PORT string
}

func getVar() {

}

func Configure() (Config, error) {
        result := Config{
		LISTEN_ADDERSS:		"127.0.0.1",
		LISTEN_PORT:		"3000",

		STORAGE_ADDRESS:	"127.0.0.1",
		STORAGE_PORT:		"3306",
	}

	v := reflect.ValueOf(result)
	f := v.Type()

        for i := 0; i < v.NumField(); i++ {
		f.Field(i).Name
	}

        return result, nil
}

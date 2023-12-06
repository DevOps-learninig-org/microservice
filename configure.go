package microservice

import (

)

const envPrefix string = "MYSERVICE"

type Config struct {
	LISTEN_ADDERSS string
	LISTEN_PORT string

	STORAGE_ADDRESS string
	STORAGE_PORT string
}

interface configFieldsTypes {
	uint | string
}

func setConfigField[ T configFieldsTypes ]( *field T, defaultVal T ) {
	
}

func Configure() (Config, error) {
        result := Config{}
	setConfigField[string]( &result.LISTEN_ADDERSS, "127.0.0.1" )
	setConfigField[string]( &result.LISTEN_PORT, "3000" )
	setConfigField[string]( &result.STORAGE_ADDRESS, "127.0.0.1" )
	setConfigField[string]( &result.STORAGE_PORT, "3306" )
	
        return result, nil
}

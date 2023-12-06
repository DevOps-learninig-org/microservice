package microservice

import (

)

type Config struct {
	LISTEN_ADDERSS string
	LISTEN_PORT string

	STORAGE_ADDRESS string
	STORAGE_PORT string
}

type configFieldsTypes interface {
	int | string
}

func setConfigField[ T configFieldsTypes ]( prefix string, field *T, defaultVal T ) {
	
}

func Configure(envPrefix string) (Config, error) {
        result := Config{}
	setConfigField[string]( envPrefix, &result.LISTEN_ADDERSS, "127.0.0.1" )
	setConfigField[string]( envPrefix, &result.LISTEN_PORT, "3000" )
	setConfigField[string]( envPrefix, &result.STORAGE_ADDRESS, "127.0.0.1" )
	setConfigField[string]( envPrefix, &result.STORAGE_PORT, "3306" )
	
        return result, nil
}

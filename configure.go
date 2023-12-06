package microservice

const envPrefix string = "MYSERVICE"

type config struct {
	LISTEN_ADDERSS string
	LISTEN_PORT string

	STORAGE_ADDRESS string
	STORAGE_PORT string
}

func Configure() {

}

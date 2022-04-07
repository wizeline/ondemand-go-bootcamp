package config

type config struct {
	Port string
}

var singleInstance *config

func GetInstance() *config {
	if singleInstance == nil {
		singleInstance = &config{Port: "8080"}
	}

	return singleInstance
}

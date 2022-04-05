package config

import "log"

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var (
	Env            = EnvDevelopment
	Port           = "8080"
	APIHost        string
	DefaultAPIHost = "http://localhost:8080"
	DBType         = "sqlite"
	DBName         = "last9.db"
	DBHost         = "localhost"
	DBPort         = "5432"
	AWSAPIID       = ""
	AWSSecretKey   = ""
)

// Initialize initializes all the env variables for this package.
// This function should be called only once during the application.
//
// Example:
//     LoadFromJSON("config.json", "dash.json")
//     AddEnvEntry("HOST", "localhost", &host)
//     AddEnvEntry("PORT", "8080", &port)
//     ...
//     config.Initialize()
func Initialize(files ...string) {
	if initDone {
		panic("config initialization done already")
	}
	LoadFromJSON(files...)

	addNewEnvEntry("ENV", &Env, EnvDevelopment)
	addNewEnvEntry("PORT", &Port, "8080")
	addNewEnvEntry("API_HOST", &APIHost, DefaultAPIHost)

	addNewEnvEntry("DB_TYPE", &DBType, DBType)
	addNewEnvEntry("DB_NAME", &DBName, DBName)
	addNewEnvEntry("DB_HOST", &DBHost, DBHost)
	addNewEnvEntry("DB_PORT", &DBPort, DBPort)

	addNewEnvEntry("AWS_API_ID", &AWSAPIID, AWSAPIID)
	addNewEnvEntry("AWS_SECRET_KEY", &AWSSecretKey, AWSSecretKey)

	// load all the env variables. Must be called at the end.
	load()
	log.Println("Inited config...👍")
}

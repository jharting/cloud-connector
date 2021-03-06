package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	ENV_PREFIX = "CLOUD_CONNECTOR"

	URL_APP_NAME                               = "URL_App_Name"
	URL_PATH_PREFIX                            = "URL_Path_Prefix"
	URL_BASE_PATH                              = "URL_Base_Path"
	OPENAPI_SPEC_FILE_PATH                     = "OpenAPI_Spec_File_Path"
	HTTP_SHUTDOWN_TIMEOUT                      = "HTTP_Shutdown_Timeout"
	SERVICE_TO_SERVICE_CREDENTIALS             = "Service_To_Service_Credentials"
	PROFILE                                    = "Enable_Profile"
	MQTT_BROKER_ADDRESS                        = "MQTT_Broker_Address"
	MQTT_BROKER_TLS_CERT_FILE                  = "MQTT_Broker_Tls_Cert_File"
	MQTT_BROKER_TLS_KEY_FILE                   = "MQTT_Broker_Tls_Key_File"
	MQTT_BROKER_TLS_CA_CERT_FILE               = "MQTT_Broker_Tls_CA_Cert_File"
	MQTT_BROKER_TLS_SKIP_VERIFY                = "MQTT_Broker_Tls_Skip_Verify"
	MQTT_BROKER_JWT_GENERATOR_IMPL             = "MQTT_Broker_JWT_Generator_Impl"
	MQTT_BROKER_JWT_FILE                       = "MQTT_Broker_JWT_File"
	DEFAULT_MQTT_BROKER_ADDRESS                = "ssl://localhost:8883"
	KAFKA_BROKERS                              = "Kafka_Brokers"
	DEFAULT_KAFKA_BROKER_ADDRESS               = "kafka:29092"
	CLIENT_ID_TO_ACCOUNT_ID_IMPL               = "Client_Id_To_Account_Id_Impl"
	CLIENT_ID_TO_ACCOUNT_ID_CONFIG_FILE        = "Client_Id_To_Account_Id_Config_File"
	CLIENT_ID_TO_ACCOUNT_ID_DEFAULT_ACCOUNT_ID = "Client_Id_To_Account_Id_Default_Account_Id"
)

type Config struct {
	UrlAppName                          string
	UrlPathPrefix                       string
	UrlBasePath                         string
	OpenApiSpecFilePath                 string
	HttpShutdownTimeout                 time.Duration
	ServiceToServiceCredentials         map[string]interface{}
	Profile                             bool
	MqttBrokerAddress                   string
	MqttBrokerTlsCertFile               string
	MqttBrokerTlsKeyFile                string
	MqttBrokerTlsCACertFile             string
	MqttBrokerTlsSkipVerify             bool
	MqttBrokerJwtGeneratorImpl          string
	MqttBrokerJwtFile                   string
	KafkaBrokers                        []string
	ClientIdToAccountIdImpl             string
	ClientIdToAccountIdConfigFile       string
	ClientIdToAccountIdDefaultAccountId string
}

func (c Config) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s: %s\n", URL_PATH_PREFIX, c.UrlPathPrefix)
	fmt.Fprintf(&b, "%s: %s\n", URL_APP_NAME, c.UrlAppName)
	fmt.Fprintf(&b, "%s: %s\n", URL_BASE_PATH, c.UrlBasePath)
	fmt.Fprintf(&b, "%s: %s\n", OPENAPI_SPEC_FILE_PATH, c.OpenApiSpecFilePath)
	fmt.Fprintf(&b, "%s: %s\n", HTTP_SHUTDOWN_TIMEOUT, c.HttpShutdownTimeout)
	fmt.Fprintf(&b, "%s: %t\n", PROFILE, c.Profile)
	fmt.Fprintf(&b, "%s: %s\n", MQTT_BROKER_ADDRESS, c.MqttBrokerAddress)
	fmt.Fprintf(&b, "%s: %s\n", MQTT_BROKER_TLS_CERT_FILE, c.MqttBrokerTlsCertFile)
	fmt.Fprintf(&b, "%s: %s\n", MQTT_BROKER_TLS_KEY_FILE, c.MqttBrokerTlsKeyFile)
	fmt.Fprintf(&b, "%s: %s\n", MQTT_BROKER_TLS_CA_CERT_FILE, c.MqttBrokerTlsCACertFile)
	fmt.Fprintf(&b, "%s: %v\n", MQTT_BROKER_TLS_SKIP_VERIFY, c.MqttBrokerTlsSkipVerify)
	fmt.Fprintf(&b, "%s: %s\n", MQTT_BROKER_JWT_GENERATOR_IMPL, c.MqttBrokerJwtGeneratorImpl)
	fmt.Fprintf(&b, "%s: %s\n", MQTT_BROKER_JWT_FILE, c.MqttBrokerJwtFile)
	fmt.Fprintf(&b, "%s: %s\n", KAFKA_BROKERS, c.KafkaBrokers)
	fmt.Fprintf(&b, "%s: %s\n", CLIENT_ID_TO_ACCOUNT_ID_IMPL, c.ClientIdToAccountIdImpl)
	fmt.Fprintf(&b, "%s: %s\n", CLIENT_ID_TO_ACCOUNT_ID_CONFIG_FILE, c.ClientIdToAccountIdConfigFile)
	fmt.Fprintf(&b, "%s: %s\n", CLIENT_ID_TO_ACCOUNT_ID_DEFAULT_ACCOUNT_ID, c.ClientIdToAccountIdDefaultAccountId)

	return b.String()
}

func GetConfig() *Config {
	options := viper.New()

	options.SetDefault(URL_PATH_PREFIX, "api")
	options.SetDefault(URL_APP_NAME, "cloud-connector")
	options.SetDefault(OPENAPI_SPEC_FILE_PATH, "/opt/app-root/src/api/api.spec.file")
	options.SetDefault(HTTP_SHUTDOWN_TIMEOUT, 2)
	options.SetDefault(SERVICE_TO_SERVICE_CREDENTIALS, "")
	options.SetDefault(PROFILE, false)
	options.SetDefault(KAFKA_BROKERS, []string{DEFAULT_KAFKA_BROKER_ADDRESS})
	options.SetDefault(MQTT_BROKER_ADDRESS, DEFAULT_MQTT_BROKER_ADDRESS)
	options.SetDefault(MQTT_BROKER_TLS_SKIP_VERIFY, false)
	options.SetDefault(MQTT_BROKER_JWT_GENERATOR_IMPL, "jwt_file_reader")
	options.SetDefault(MQTT_BROKER_JWT_FILE, "cloud-connector-mqtt-jwt.txt")
	options.SetDefault(CLIENT_ID_TO_ACCOUNT_ID_IMPL, "config_file_based")
	options.SetDefault(CLIENT_ID_TO_ACCOUNT_ID_CONFIG_FILE, "client_id_to_account_id_map.json")
	options.SetDefault(CLIENT_ID_TO_ACCOUNT_ID_DEFAULT_ACCOUNT_ID, "111000")

	options.SetEnvPrefix(ENV_PREFIX)
	options.AutomaticEnv()

	return &Config{
		UrlPathPrefix:                       options.GetString(URL_PATH_PREFIX),
		UrlAppName:                          options.GetString(URL_APP_NAME),
		UrlBasePath:                         buildUrlBasePath(options.GetString(URL_PATH_PREFIX), options.GetString(URL_APP_NAME)),
		OpenApiSpecFilePath:                 options.GetString(OPENAPI_SPEC_FILE_PATH),
		HttpShutdownTimeout:                 options.GetDuration(HTTP_SHUTDOWN_TIMEOUT) * time.Second,
		ServiceToServiceCredentials:         options.GetStringMap(SERVICE_TO_SERVICE_CREDENTIALS),
		Profile:                             options.GetBool(PROFILE),
		KafkaBrokers:                        options.GetStringSlice(KAFKA_BROKERS),
		MqttBrokerAddress:                   options.GetString(MQTT_BROKER_ADDRESS),
		MqttBrokerTlsCertFile:               options.GetString(MQTT_BROKER_TLS_CERT_FILE),
		MqttBrokerTlsKeyFile:                options.GetString(MQTT_BROKER_TLS_KEY_FILE),
		MqttBrokerTlsCACertFile:             options.GetString(MQTT_BROKER_TLS_CA_CERT_FILE),
		MqttBrokerTlsSkipVerify:             options.GetBool(MQTT_BROKER_TLS_SKIP_VERIFY),
		MqttBrokerJwtGeneratorImpl:          options.GetString(MQTT_BROKER_JWT_GENERATOR_IMPL),
		MqttBrokerJwtFile:                   options.GetString(MQTT_BROKER_JWT_FILE),
		ClientIdToAccountIdImpl:             options.GetString(CLIENT_ID_TO_ACCOUNT_ID_IMPL),
		ClientIdToAccountIdConfigFile:       options.GetString(CLIENT_ID_TO_ACCOUNT_ID_CONFIG_FILE),
		ClientIdToAccountIdDefaultAccountId: options.GetString(CLIENT_ID_TO_ACCOUNT_ID_DEFAULT_ACCOUNT_ID),
	}
}

func buildUrlBasePath(pathPrefix string, appName string) string {
	return fmt.Sprintf("/%s/%s/v1", pathPrefix, appName)
}

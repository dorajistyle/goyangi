package config

// Constants for database.
const (
	MysqlProtocol = "tcp"
	// for DEVELOPMENT
	MysqlHostDev     = "127.0.0.1"
	MysqlPortDev     = "3306"
	MysqlUserDev     = "root"
	MysqlPasswordDev = ""
	MysqlDatabaseDev = "goyangi_dev"
	MysqlOptionsDev  = "charset=utf8&parseTime=True"
	MysqlDSLDev      = MysqlUserDev + ":" + MysqlPasswordDev + "@" + MysqlProtocol + "(" + MysqlHostDev + ":" + MysqlPortDev + ")/" + MysqlDatabaseDev + "?" + MysqlOptionsDev
	// for TEST
	MysqlHostTest     = "127.0.0.1"
	MysqlPortTest     = "3306"
	MysqlUserTest     = "root"
	MysqlPasswordTest = ""
	MysqlDatabaseTest = "goyangi_test"
	MysqlOptionsTest  = "charset=utf8&parseTime=True"
	MysqlDSLTest      = MysqlUserTest + ":" + MysqlPasswordTest + "@" + MysqlProtocol + "(" + MysqlHostTest + ":" + MysqlPortTest + ")/" + MysqlDatabaseTest + "?" + MysqlOptionsTest
	// for PRODUCTION
	MysqlHost          = "127.0.0.1"
	MysqlPort          = "3306"
	MysqlUser          = "root"
	MysqlPassword      = ""
	MysqlDatabase      = "goyangi"
	MysqlOptions       = "charset=utf8&parseTime=True"
	MysqlDSLProduction = MysqlUser + ":" + MysqlPassword + "@" + MysqlProtocol + "(" + MysqlHost + ":" + MysqlPort + ")/" + MysqlDatabase + "?" + MysqlOptions
	
	PostgresProtocol = "tcp"
	// for DEVELOPMENT
	PostgresHostDev              = "127.0.0.1"
	PostgresPortDev              = "5432"
	PostgresUserDev              = "postgres"
	PostgresPasswordDev          = "postgres"
	PostgresDatabaseDev          = "goyangi_dev"
	PostgresOptionsDev           = "sslmode=disable"
	PostgresDSLDev               = "host="+ PostgresHostDev +" port="+ PostgresPortDev + " dbname="+ PostgresDatabaseDev + " " + PostgresOptionsDev +" user="+ PostgresUserDev + " password=" + PostgresPasswordDev
	PostgresDSLMigrateDev        = PostgresUserDev + ":" + PostgresPasswordDev + "@" + PostgresHostDev + ":" + PostgresPortDev + "/" + PostgresDatabaseDev + "?" + PostgresOptionsDev
	// for TEST
	PostgresHostTest             = "127.0.0.1"
	PostgresPortTest             = "5432"
	PostgresUserTest             = "postgres"
	PostgresPasswordTest         = "postgres"
	PostgresDatabaseTest         = "goyangi_test"
	PostgresOptionsTest          = "sslmode=disable"
	PostgresDSLTest              = "host="+ PostgresHostTest +" port="+ PostgresPortTest + " dbname="+ PostgresDatabaseTest + " " + PostgresOptionsTest +" user="+ PostgresUserTest + " password=" + PostgresPasswordTest
	PostgresDSLMigrateTest       = PostgresUserTest + ":" + PostgresPasswordTest + "@" + PostgresHostTest + ":" + PostgresPortTest + "/" + PostgresDatabaseTest + "?" + PostgresOptionsTest
	// for PRODUCTION
	PostgresHost                 = "127.0.0.1"
	PostgresPort                 = "5432"
	PostgresUser                 = "postgres"
	PostgresPassword             = "postgres"
	PostgresDatabase             = "goyangi_production"
	PostgresOptions              = "sslmode=disable"
	PostgresDSLProduction        = "host="+ PostgresHost +" port="+ PostgresPort + " dbname="+ PostgresDatabase + " " + PostgresOptions +" user="+ PostgresUser + " password=" + PostgresPassword
	PostgresDSLMigrateProduction = PostgresUser + ":" + PostgresPassword + "@" + PostgresHost + ":" + PostgresPort + "/" + PostgresDatabase + "?" + PostgresOptions
	
	
)

// MysqlDSL return mysql DSL.
func MysqlDSL() string {
	var mysqlDSL string
	switch Environment {
	case "DEVELOPMENT":
		mysqlDSL = MysqlDSLDev
	case "TEST":
		mysqlDSL = MysqlDSLTest
	default:
		mysqlDSL = MysqlDSLProduction
	}
	return mysqlDSL
}

// PostgresDSL return postgres DSL.
func PostgresDSL() string {
	var postgresDSL string
	switch Environment {
		case "DEVELOPMENT":
		postgresDSL = PostgresDSLDev
		case "TEST":
		postgresDSL = PostgresDSLTest
		default:
		postgresDSL = PostgresDSLProduction
	}
	return postgresDSL
}

// PostgresDSL return postgres DSL.
func PostgresMigrateDSL() string {
	var postgresDSL string
	switch Environment {
		case "DEVELOPMENT":
		postgresDSL = PostgresDSLMigrateDev
		case "TEST":
		postgresDSL = PostgresDSLMigrateTest
		default:
		postgresDSL = PostgresDSLMigrateProduction
	}
	return postgresDSL
}

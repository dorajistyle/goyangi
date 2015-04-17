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

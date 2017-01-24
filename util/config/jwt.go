package config

// Constants for Authentication
const (
	JWTSigningMethodClient  = "HMAC256" // HMAC256 | RSA256
	JWTSigningKeyHMACClient = ""
	JWTSigningKeyRSAClient  = "secretRSA"
	JWTPublicKeyRSAClient   = "secretRSAPUBLIC"
	JWTSigningMethodServer  = "HMAC256" // HMAC256 | RSA256
	JWTExpriationHourServer = 24
	JWTSigningKeyHMACServer = ""
	JWTSigningKeyRSAServer  = `-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----`
	JWTPublicKeyRSAServer = `-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----`
)

package main

import (
	"fmt"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
)

// Hardcoded credentials - security issue
var (
	// Database credentials hardcoded
	dbUser     = "admin"
	dbPassword = "Password123!"
	dbHost     = "prod-database.internal.company.com"
	dbConnStr  = "mysql://root:rootpassword@localhost:3306/production"

	// API credentials
	apiKey    = "api_key_a1b2c3d4e5f6g7h8i9j0"
	apiSecret = "secret_xyz789abc123def456"

	// Generic secrets
	encryptionKey = "aes256_encryption_key_do_not_share"
	signingSecret = "hmac_sha256_signing_secret_value"
	authToken     = "bearer_token_abc123xyz789"

	// Internal service credentials
	serviceAccount = "svc-account@project.iam.gserviceaccount.com"
	serviceKey     = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBg..."
)

func main() {
	fmt.Println("Hello world")

	// Reference to avoid unused variable errors
	_ = dbUser
	_ = dbPassword
	_ = dbHost
	_ = dbConnStr
	_ = apiKey
	_ = apiSecret
	_ = encryptionKey
	_ = signingSecret
	_ = authToken
	_ = serviceAccount
	_ = serviceKey
}

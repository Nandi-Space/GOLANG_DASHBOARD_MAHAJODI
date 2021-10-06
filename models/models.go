package models

import (
	"github.com/golang-jwt/jwt"
)

type APIResponse struct {
	Status  bool        `json:"status ,omitempty"`
	Message bool        `json:"message, omitempty"`
	Body    interface{} `json:"body, omitempty"`
}

type Config struct {
	Database struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Name     string `toml:"name"`
	} `toml:"database"`

	Logging struct {
		Level string `toml:"logging"`
	} `toml:"logging"`

	Server struct {
		Listen string `toml:"listen"`
	} `toml:"server"`

	Firebase struct {
		CredentialPath string `toml:"credentialPath"`
	} `toml:"firebase"`

	AWS struct {
		BucketName string `toml:"bucketName"`
		Region     string `toml:"region"`
		Secret     string `toml:"secret"`
		ID         string `toml:"id"`
	} `toml:"aws"`

	Stripe struct {
		Key           string `json:"key"`
		WebHookSecret string `json:"webhookSecret"`
	}

	Mail struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"mail"`

	Keys struct {
		PrivateKey string `toml:"privateKey"`
		PublicKey  string `toml:"publicKey"`
	} `toml:"keys"`
}

type JWTToken struct {
	jwt.StandardClaims
	UserID int64  `json:"userId"`
	Email  string `json:"email"`
}

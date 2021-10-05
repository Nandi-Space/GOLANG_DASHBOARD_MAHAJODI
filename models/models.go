package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type APIResponse struct {
	Status  bool        `json:"status" omitempty`
	Message bool        `json:"message" omitempty`
	Body    interface{} `json:"body" omitempty`
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
	AdminID int64  `json:"adminId"`
	Mobile  string `json:"mobile"`
}

type Filters struct {
	ID              int64     `json:"id"`
	HigherAge       int64     `json:"higher_age"`
	LowerAge        int64     `json:"lower_age"`
	Gender          *[]string `json:"gender"`
	MaritalStatus   *[]string `json:"marital_status"`
	Profession      *[]string `json:"profession"`
	Cast            *[]string `json:"cast"`
	Country         *[]string `json:"country"`
	ResidingCountry *[]string `json:"residing_country"`
	District        *[]string `json:"district"`
	ResidentStatus  *[]string `json:"resident_status"`
	Citizenship     *[]string `json:"citizenship"`
	UserID          int64     `json:"user_id"`
	Religion        *[]string `json:"religion"`
	Education       *[]string `json:"education"`
	Siblings        *[]string `json:"siblings"`
}

type Message struct {
	Type        string    `json:"type,required"`
	ID          string    `json:"id"`
	Data        string    `json:"data,omitempty"`
	RecipientID int64     `json:"recipient_id,omitempty"`
	SenderID    int64     `json:"sender_id,omitempty"`
	SentAt      time.Time `json:"sent_at,omitempty"`
}

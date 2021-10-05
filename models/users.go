package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string ` json:"password,omitempty"`
	//PlanType           int64     `json:"plan_type,omitempty"`
	OTP                string    `json:"otp,omitempty"`
	IsVerified         int       `json:"is_verified,omitempty"`
	CreatedAt          time.Time ` json:"created_at,omitempty"`
	UpdatedAt          time.Time ` json:"updated_at,omitempty"`
	DetailID           int64     `json:"detail_id,omitempty"`
	UserID             int64     `json:"user_id,omitempty"`
	Community          string    ` json:"community,omitempty"`
	DOB                string    `json:"date_of_birth,omitempty"`
	CurrentAddress     string    ` json:"current_address,omitempty"`
	MaritalStatus      string    ` json:"marital_status,omitempty"`
	Citizenship        string    ` json:"citizenship,omitempty"`
	Country            string    ` json:"country,omitempty"`
	District           string    ` json:"district,omitempty"`
	Eduction           string    ` json:"education,omitempty"`
	PlanId             int64     ` json:"plan_id,omitempty"`
	Siblings           string    ` json:"siblings,omitempty"`
	Religion           string    ` json:"religion,omitempty"`
	Height             string    ` json:"height,omitempty"`
	Age                int32     `json:"age,omitempty"`
	Bio                string    ` json:"bio,omitempty"`
	Profession         string    ` json:"profession,omitempty"`
	Caste              string    ` json:"caste,omitempty"`
	Path1              string    ` json:"path1,omitempty"`
	Path2              string    ` json:"path2,omitempty"`
	Path3              string    ` json:"path3,omitempty"`
	InterestedIn       string    ` json:"interested_in,omitempty"`
	Gender             string    ` json:"gender,omitempty"`
	Status             string    ` json:"status,omitempty"`
	SentStatus         string    ` json:"sent_status,omitempty"`
	FavStatus          string    ` json:"fav_status,omitempty"`
	BlockStatus        string    ` json:"block_status,omitempty"`
	ResidentStatus     string    ` json:"resident_status,omitempty"`
	DateOfSubscription string    ` json:"date_of_subscription,omitempty"`
	IsOnline           bool      ` json:"is_online,omitempty"`

	LastMessage string    ` json:"last_message,omitempty"`
	PaidAt      time.Time `json:"-"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

}

func (u *User) Validate(action string) error {
	switch action {
	case "update":

		if u.UserName == "" {
			return errors.New("required nickname")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if u.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		return nil
	default:

		if u.UserName == "" {
			return errors.New("required User Name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

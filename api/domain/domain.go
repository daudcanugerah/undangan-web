package domain

import "time"

type RoleType uint

const (
	RoleAdmin RoleType = iota + 1
	RoleUser
)

type UserClaim struct {
	ID    string
	Email string
	Role  RoleType
}

type SafeUser struct {
	ID        string
	Email     string
	Name      string
	Profile   string
	Role      RoleType
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	Profile   string
	Role      RoleType
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PublicTemplate struct {
	ID            string
	Name          string
	Description   string
	PriceInterval string
	Price         int
	Type          string
	Tags          []string
	CoverImage    string
	State         int
	Slug          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MessageTemplate struct {
	Text     string
	Provider string
}

type UserTemplate struct {
	ID              string
	UserID          string // reference to User ID
	BaseTemplateID  string // reference to PublicTemplate ID
	State           int
	Slug            string
	URL             string
	MessageTemplate []MessageTemplate
	Name            string
	CoverImage      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ExpireAt        time.Time
}

type Guest struct {
	ID             string
	UserTemplateID string // refernce to UserTemplate ID
	Name           string
	Group          string
	Person         int
	Tags           []string
	Telp           string
	Address        string
	Message        string
	Attend         bool
	ViewAt         *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

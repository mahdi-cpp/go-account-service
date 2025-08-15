package go_account_service

import "time"

func (u *User) SetID(id string)                 { u.ID = id }
func (u *User) SetCreationDate(t time.Time)     { u.CreationDate = t }
func (u *User) SetModificationDate(t time.Time) { u.ModificationDate = t }
func (u *User) GetID() string                   { return u.ID }
func (u *User) GetCreationDate() time.Time      { return u.CreationDate }
func (u *User) GetModificationDate() time.Time  { return u.ModificationDate }

type User struct {
	ID               string    `json:"id"`
	Username         string    `json:"username"`
	PhoneNumber      string    `json:"phoneNumber"`
	Email            string    `json:"email"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Bio              string    `json:"bio"`
	AvatarURL        string    `json:"avatarURL"`
	IsOnline         bool      `json:"isOnline"`
	LastSeen         time.Time `json:"lastSeen"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}

type Update struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatarURL"`
	IsOnline    *bool  `json:"isOnline"`

	// Date
	LastSeen         *time.Time `json:"lastSeen"`
	ModificationDate *time.Time `json:"modificationDate"`
}

type Options struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatarURL"`
	IsOnline    *bool  `json:"isOnline"`

	UsernameQuery string `json:"usernameQuery"`

	// Date filters
	LastSeen         *time.Time `json:"lastSeen"`
	ModificationDate *time.Time `json:"modificationDate"`
	CreatedAfter     *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore    *time.Time `json:"createdBefore,omitempty"`

	// Pagination
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`

	// Sorting
	SortBy    string `json:"sortBy,omitempty"`    // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"
}

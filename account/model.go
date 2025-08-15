package account

import "time"

func (u *User) SetID(id string)                 { u.ID = id }
func (u *User) SetCreationDate(t time.Time)     { u.CreationDate = t }
func (u *User) SetModificationDate(t time.Time) { u.ModificationDate = t }
func (u *User) GetID() string                   { return u.ID }
func (u *User) GetCreationDate() time.Time      { return u.CreationDate }
func (u *User) GetModificationDate() time.Time  { return u.ModificationDate }

type User struct {

	// Core Identity & Basic Information
	ID          string `json:"id"`          // Unique identifier for the user
	Username    string `json:"username"`    // User's login username (must be unique)
	DisplayName string `json:"displayName"` // Name displayed publicly (can differ from Username)
	PhoneNumber string `json:"phoneNumber"` // User's phone number
	Email       string `json:"email"`       // User's email address
	FirstName   string `json:"firstName"`   // User's first name
	LastName    string `json:"lastName"`    // User's last name
	Bio         string `json:"bio"`         // Short biography or "About Me" section
	AvatarURL   string `json:"avatarURL"`   // URL to the user's profile picture
	IsVerified  bool   `json:"isVerified"`  // Indicates if the account is verified (e.g., for official accounts)

	// Presence & Connectivity (for chat, social apps)
	IsOnline      bool      `json:"isOnline"`      // Current online status of the user
	LastSeen      time.Time `json:"lastSeen"`      // Last timestamp the user was seen online
	StatusMessage string    `json:"statusMessage"` // User's custom status message (e.g., "Busy", "Available")

	// Privacy & Social Features (for social, photo, music apps)
	ProfileVisibility string   `json:"profileVisibility"` // Profile visibility setting: "public", "private", "friendsOnly"
	FollowerCount     int      `json:"followerCount"`     // Number of followers this user has
	FollowingCount    int      `json:"followingCount"`    // Number of Users this user is following
	BlockedUserIDs    []string `json:"blockedUserIDs"`    // List of user IDs blocked by this user

	// Preferences & Customization
	PreferredLanguage string   `json:"preferredLanguage"` // User's preferred language (e.g., "en-US", "fa-IR")
	Timezone          string   `json:"timezone"`          // User's timezone (e.g., "Asia/Tehran")
	ThemePreference   string   `json:"themePreference"`   // User's UI theme preference: "light", "dark", "system"
	Interests         []string `json:"interests"`         // List of user's interests (for content recommendations)

	// Account Status & Usage
	SubscriptionTier   string `json:"subscriptionTier"`   // User's subscription level (e.g., "free", "premium", "pro")
	AccountStatus      string `json:"accountStatus"`      // Current status of the user's account: "active", "suspended", "deactivated"
	IsTwoFactorEnabled bool   `json:"isTwoFactorEnabled"` // Indicates if two-factor authentication is enabled

	// Timestamps & Tracking
	LastActivity     time.Time `json:"lastActivity"`     // Timestamp of the user's last public activity in the app
	CreationDate     time.Time `json:"creationDate"`     // Date and time when the user account was created
	ModificationDate time.Time `json:"modificationDate"` // Date and time of the last modification to user's profile data

	// Generic/Extensible Metadata (for highly specific or future data)
	Metadata map[string]string `json:"metadata"` // Flexible field for storing additional, application-specific data
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

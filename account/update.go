package account

import "github.com/mahdi-cpp/api-go-pkg/update"

// Initialize updater
var chatUpdater = update.NewUpdater[User, Update]()

func init() {

	// Basic Info Updates
	chatUpdater.AddScalarUpdater(func(c *User, u Update) {
		if u.Username != "" {
			c.Username = u.Username
		}
		if u.PhoneNumber != "" {
			c.PhoneNumber = u.PhoneNumber
		}
		if u.Username != "" {
			c.Username = u.Username
		}
		if u.Email != "" {
			c.Email = u.Email
		}
		if u.FirstName != "" {
			c.FirstName = u.FirstName
		}
		if u.LastName != "" {
			c.LastName = u.LastName
		}
		if u.Bio != "" {
			c.Bio = u.Bio
		}
		if u.AvatarURL != "" {
			c.AvatarURL = u.AvatarURL
		}
		if u.IsVerified != nil {
			c.IsVerified = *u.IsVerified
		}
	})

	// Presence & Connectivity (for chat, social apps)
	chatUpdater.AddScalarUpdater(func(c *User, u Update) {
		if u.IsOnline != nil {
			c.IsOnline = *u.IsOnline
		}
		if u.LastSeen != nil {
			c.LastSeen = *u.LastSeen
		}
		if u.StatusMessage != "" {
			c.StatusMessage = u.StatusMessage
		}
	})

	/// Privacy & Social Features (for social, photo, music apps)
	chatUpdater.AddScalarUpdater(func(c *User, u Update) {
		if u.ProfileVisibility != "" {
			c.ProfileVisibility = u.ProfileVisibility
		}
		if u.IsVerified != nil {
			c.IsVerified = *u.IsVerified
		}
		if u.FollowerCount != 0 {
			c.FollowerCount = u.FollowerCount
		}
		if u.FollowingCount != 0 {
			c.FollowingCount = u.FollowingCount
		}
	})

	// Banned Users Collection Updates
	chatUpdater.AddCollectionUpdater(func(c *User, u Update) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.BlockedUsers,
			Add:         u.AddBlockedUsers,
			Remove:      u.RemoveBlockedUsers,
		}
		c.BlockedUserIDs = update.ApplyCollectionUpdate(c.BlockedUserIDs, op)
	})

}

// Update applies the updates to a chat
func (u *User) Update(update Update) *User {
	chatUpdater.Apply(u, update)
	return u
}

//func (c *User) Save() error {
//	c.mutex.Lock()
//	defer c.mutex.Unlock()
//
//	return utils.WriteData(c, c.Filepath)
//}

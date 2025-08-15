package go_account_service

import (
	"github.com/mahdi-cpp/MessageKit/pkg/search_manager"
	"strings"
)

var PHAssetLessFuncs = map[string]search_manager.LessFunction[*User]{
	"id":               func(a, b *User) bool { return a.ID < b.ID },
	"creationDate":     func(a, b *User) bool { return a.CreationDate.Before(b.CreationDate) },
	"modificationDate": func(a, b *User) bool { return a.ModificationDate.Before(b.ModificationDate) },
	"title":            func(a, b *User) bool { return a.Username < b.Username },
}

func GetLessFunc(sortBy, sortOrder string) search_manager.LessFunction[*User] {

	fn, exists := PHAssetLessFuncs[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "end" {
		return func(a, b *User) bool { return !fn(a, b) }
	}
	return fn
}

func BuildUserSearchCriteria(with Options) search_manager.SearchCriteria[*User] {

	return func(c *User) bool {

		// ID filter
		if with.ID != "" && c.ID != with.ID {
			return false
		}

		// Title search_manager (case-insensitive)
		if with.UsernameQuery != "" {
			query := strings.ToLower(with.UsernameQuery)
			username := strings.ToLower(c.Username)
			if !strings.Contains(username, query) {
				return false
			}
		}

		// Username exact match
		if with.Username != "" && c.Username != with.Username {
			return false
		}

		// Boolean flags
		if with.IsOnline != nil && c.IsOnline != *with.IsOnline {
			return false
		}

		// Date filters
		if with.CreatedAfter != nil && c.CreationDate.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && c.CreationDate.After(*with.CreatedBefore) {
			return false
		}
		//if with.ActiveAfter != nil && c.LastMessage != nil &&
		//	c.LastMessage.CreationDate.Before(*with.ActiveAfter) {
		//	return false
		//}

		return true
	}
}

func Search(chats []*User, with Options) []*User {

	// Build criteria
	criteria := BuildUserSearchCriteria(with)

	// Execute search_manager
	results := search_manager.Search(chats, criteria)

	// Sort results if needed
	if with.SortBy != "" {
		lessFn := GetLessFunc(with.SortBy, with.SortOrder)
		if lessFn != nil {
			search_manager.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*User, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	// Apply pagination
	start := with.Offset
	end := start + with.Limit
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}

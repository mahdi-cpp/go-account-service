package account

import (
	"github.com/mahdi-cpp/api-go-pkg/search"
	"strings"
)

var PHAssetLessFuncs = map[string]search.LessFunction[*User]{
	"id":               func(a, b *User) bool { return a.ID < b.ID },
	"creationDate":     func(a, b *User) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"modificationDate": func(a, b *User) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
	"title":            func(a, b *User) bool { return a.UserName < b.UserName },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*User] {

	fn, exists := PHAssetLessFuncs[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "end" {
		return func(a, b *User) bool { return !fn(a, b) }
	}
	return fn
}

func BuildUserSearchCriteria(with Options) search.SearchCriteria[*User] {

	return func(c *User) bool {

		// ID filter
		if with.ID != "" && c.ID != with.ID {
			return false
		}

		// Title search_manager (case-insensitive)
		if with.UsernameQuery != "" {
			query := strings.ToLower(with.UsernameQuery)
			username := strings.ToLower(c.UserName)
			if !strings.Contains(username, query) {
				return false
			}
		}

		// Username exact match
		if with.Username != "" && c.UserName != with.Username {
			return false
		}

		// Boolean flags
		if with.IsOnline != nil && c.IsOnline != *with.IsOnline {
			return false
		}

		// Date filters
		if with.CreatedAfter != nil && c.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && c.CreatedAt.After(*with.CreatedBefore) {
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
	results := search.Search(chats, criteria)

	// Sort results if needed
	if with.SortBy != "" {
		lessFn := GetLessFunc(with.SortBy, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
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

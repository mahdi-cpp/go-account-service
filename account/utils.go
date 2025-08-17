package account

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// GetUserId retrieves the userID from the request header.
// It returns the userID string and an error if the header is missing.
func GetUserId(c *gin.Context) (string, error) {
	// Get the "userID" header from the request.
	// If the header is not present, c.GetHeader will return an empty string "".
	userIDStr := c.GetHeader("userID")

	// Check if the retrieved userID string is empty.
	if userIDStr == "" {
		// If it's empty, return an empty string and a new error.
		// errors.New is used for simple, static error messages.
		// fmt.Errorf is useful if you need to include dynamic information in your error message.
		return "", errors.New("userID header is missing or empty")
		// Or using fmt.Errorf for a more descriptive error:
		// return "", fmt.Errorf("missing or empty 'userID' header in request for path: %s", c.Request.URL.Path)
	}

	// If the userID is found and not empty, return it and a nil error.
	return userIDStr, nil
}

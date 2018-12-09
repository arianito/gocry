package cry

import "golang.org/x/crypto/bcrypt"

// Hash a password to something unreadable to human being
func Bcrypt(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedPassword)
}

// Check if the password matches the hashed one
func Check(hashed string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}
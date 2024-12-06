package db

import (
	"errors"
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" binding:"required"`
	Password string `gorm:"size:255;not null;" json:"password" binding:"required"`
}

// GetUserById retrieves a user from the database by their unique ID.
// It returns the User object along with any associated groceries,
// or an error if the user is not found or there is a database issue.
func GetUserById(uid uint) (User, error) {
	var user User

	db, err := ConnectDB()

	if err != nil {
		log.Println(err)
		return User{}, err
	}
	if err := db.Preload("Groceries").Where("id=?", uid).Find(&user).Error; err != nil {
		return user, errors.New("user not found")

	}

	return user, nil
}

// HashedPassword sets the User's Password field to a bcrypt-hashed version of
// the current value, and trims/escapes the Username field. It returns an error
// if the hashing operation fails. It modifies the User directly and does not
// return a new User.
func (u *User) HashedPassword() error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
		return err
	}

	u.Password = string(hashPass)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

// VerifyPassword checks if a given plaintext password matches a bcrypt-hashed
// password. It returns a nil error if the password is valid, or an error if
// the password does not match the hash, or if there is an issue with the hash
// itself.
func VerifyPassword(password, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}

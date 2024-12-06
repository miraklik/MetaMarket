package handlers

import (
	"net/http"
	"nft-marketplace/db"
	"nft-marketplace/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Server struct {
	db *gorm.DB
}

// NewServer returns a new Server instance with the given database connection. It is
// used to create a new server with a database connection that is already open. The
// caller is responsible for ensuring that the database connection is valid and will
// stay open for the duration of the server's lifetime. If the database connection is
// closed, the server will panic.
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

// Register takes a username and password as input and creates a new user. If the
// input is invalid, it returns a 400 Bad Request response. If the user is created
// successfully, it returns a 201 Created response with a message indicating that
// the user was created. If there is a database error, it returns a 500 Internal
// Server Error response.
func (s *Server) Register(c *gin.Context) {
	var Input RegisterUserInput

	if err := c.ShouldBind(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := db.User{Username: Input.Username, Password: Input.Password}
	user.HashedPassword()

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := utils.ValidatePassword(Input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Login takes a username and password as input and returns a valid JWT token if the
// credentials are valid. If the credentials are invalid, it returns an error. The function
// queries the database for the user with the given username and if the user is found, it
// verifies the password using bcrypt. If the verification fails, it returns an error.
// Otherwise, it generates a JWT token using the user's ID and returns it.
func (s *Server) Login(c *gin.Context) {
	var input LoginUserInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := db.User{Username: input.Username, Password: input.Password}

	token, err := s.LoginCheck(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}

// LoginCheck takes a username and password as input and returns a valid JWT token if the
// credentials are valid. If the credentials are invalid, it returns an error. The function
// queries the database for the user with the given username and if the user is found, it
// verifies the password using bcrypt. If the verification fails, it returns an error.
// Otherwise, it generates a JWT token using the user's ID and returns it.
func (s *Server) LoginCheck(username, password string) (string, error) {
	var err error

	user := db.User{}

	if err = s.db.Model(db.User{}).Where("username=?", username).Take(&user).Error; err != nil {
		return "", err
	}

	err = db.VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

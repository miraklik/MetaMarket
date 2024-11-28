package auth

import (
	"net/http"
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

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Register(c *gin.Context) {
	var Input RegisterUserInput

	if err := c.ShouldBind(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{Username: Input.Username, Password: Input.Password}
	user.HashedPassword()

	if err := s.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (s *Server) Login(c *gin.Context) {
	var input LoginUserInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{Username: input.Username, Password: input.Password}

	token, err := s.LoginCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "User logged in"})
}

func (s *Server) LoginCheck(username, password string) (string, error) {
	var err error

	user := User{}

	if err = s.db.Model(User{}).Where("username=?", username).Take(&user).Error; err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil

}

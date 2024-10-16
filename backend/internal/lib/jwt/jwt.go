package jwt

import (
	"internal/internal/domain/models"
	"internal/internal/lib/jwt"
	"time"

	
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := 
}
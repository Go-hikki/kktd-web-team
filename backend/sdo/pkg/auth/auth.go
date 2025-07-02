package auth

import (
	"context"
	"errors"
	"strings"
	"time"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/m1sty32/sdo/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db        *sqlx.DB
	redis     *redis.Client
	secretKey string
}

func NewAuthService(db *sqlx.DB, redis *redis.Client, secretKey string) *AuthService {
	return &AuthService{db: db, redis: redis, secretKey: secretKey}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (s *AuthService) Register(ctx context.Context, req RegisterRequest) error {
	// Проверяем, существует ли пользователь с таким email
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1`
	err := s.db.GetContext(ctx, &count, query, req.Email)
	if err != nil {
		logrus.Errorf("Failed to check existing user: %v", err)
		return err
	}
	if count > 0 {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("Failed to hash password: %v", err)
		return err
	}

	user := models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query = `INSERT INTO users (email, password_hash, role, created_at, updated_at) 
	          VALUES (:email, :password_hash, :role, :created_at, :updated_at)`
	_, err = s.db.NamedExecContext(ctx, query, user)
	if err != nil {
		logrus.Errorf("Failed to register user: %v", err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("user with this email already exists")
		}
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	var user models.User
	query := `SELECT id, email, password_hash, role FROM users WHERE email = $1`
	err := s.db.GetContext(ctx, &user, query, req.Email)
	if err != nil {
		logrus.Errorf("Failed to find user: %v", err)
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		logrus.Errorf("Invalid password: %v", err)
		return nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		logrus.Errorf("Failed to generate token: %v", err)
		return nil, err
	}

	// Сохраняем токен в Redis
	err = s.redis.Set(ctx, "token:"+user.Email, tokenString, 24*time.Hour).Err()
	if err != nil {
		logrus.Errorf("Failed to store token in Redis: %v", err)
		// Не прерываем выполнение, так как это не критично
	}

	return &LoginResponse{Token: tokenString}, nil
}
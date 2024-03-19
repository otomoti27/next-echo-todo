package user_test

import (
	"echo-api/domain"
	"echo-api/service/user"
	"echo-api/service/user/mocks"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUp(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	s := user.NewService(mockRepo)

	mockUser := domain.User{Email: "test@example.com", Password: "password", Name: "Test User"}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(mockUser.Password), bcrypt.DefaultCost)
	mockUserWithHash := domain.User{Email: "test@example.com", Password: string(hashedPassword), Name: "Test User"}

	mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = mockUserWithHash
	})

	result, err := s.SignUp(mockUser)
	assert.Nil(t, err)
	assert.Equal(t, mockUserWithHash.Email, result.Email)
	// assert.NotEqual(t, "password", result.Password)
}

func TestLogIn(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	s := user.NewService(mockRepo)

	mockUser := domain.User{Email: "test@example.com", Password: "password", Name: "Test User"}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(mockUser.Password), bcrypt.DefaultCost)
	mockUserWithHash := domain.User{ID: 1, Email: "test@example.com", Password: string(hashedPassword), Name: "Test User"}

	mockRepo.On("GetByEmail", mock.AnythingOfType("*domain.User"), mock.AnythingOfType("string")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = mockUserWithHash
	})

	os.Setenv("SECRET", "secret")
	tokenString, err := s.LogIn(mockUser)
	assert.Nil(t, err)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	assert.Nil(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, float64(mockUserWithHash.ID), claims["user_id"])
	assert.WithinDuration(t, time.Now().Add(time.Hour*72), time.Unix(int64(claims["exp"].(float64)), 0), time.Second)
}

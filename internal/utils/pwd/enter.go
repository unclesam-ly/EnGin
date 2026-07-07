package pwd

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("密码加密失败", zap.Error(err))
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

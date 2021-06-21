package auth

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/huprince/quick-gin/connections"
	"github.com/huprince/quick-gin/models"
	"github.com/huprince/quick-gin/services"
)

// GenerateBasicToken 生成 Baisc 认证 token
func GenerateBasicToken(username, password string) string {
	userInfo := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(userInfo))
}

// GetBasicAuthInfo 从 basicToken 中获取用户信息
func GetBasicAuthInfo(basicToken string) (username, password string, err error) {
	value, err := base64.StdEncoding.DecodeString(basicToken)
	if err != nil {
		return "", "", err
	}
	strValue := string(value)
	if idx := strings.Index(strValue, ":"); idx > 0 {
		return strValue[:idx], strValue[idx:], nil
	}
	return "", "", nil
}

func CheckBaiscAuth(token string) (bool, error) {
	if username, password, err := GetBasicAuthInfo(token); err != nil {
		db, err := connections.InitDB()
		if err != nil {
			return false, err
		}
		user := &models.User{}
		r := db.Where(&models.User{Email: username}).First(user)
		if r.RowsAffected > 0 && services.IsCorrectPassword(user.Password, password) {
			return true, nil
		}
	}
	return false, errors.New("Auth fail")
}
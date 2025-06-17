package utils

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/pquerna/otp"
)

func ExtractSecretFromTOTPURL(totpURL string) (string, error) {
	parsedURL, err := url.Parse(totpURL)
	if err != nil {
		return "", err
	}

	// 获取查询参数
	queryParams := parsedURL.Query()

	// 从查询参数中提取 "secret"
	secret := queryParams.Get("secret")
	if secret == "" {
		return "", errors.New("secret not found in TOTP URL")
	}

	return secret, nil
}
func ExtractUidFromOTPURL(otpURL string) (string, error) {
	key, err := otp.NewKeyFromURL(otpURL)
	if err != nil {
		return "", fmt.Errorf("invalid OTP URL: %v", err)
	}
	return key.AccountName(), nil // 返回 Uid
}

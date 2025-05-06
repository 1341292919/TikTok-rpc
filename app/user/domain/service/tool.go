package service

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/pkg/crypt"
	"TikTok-rpc/pkg/utils"
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func (uc *UserService) OptSecret(username string, user *model.MFAMessage) (*model.MFA, error) {
	var MFA = &model.MFA{}
	var buf bytes.Buffer

	if user.Secret == "" {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "tiktok",
			AccountName: username,
			Period:      30,
			Digits:      otp.DigitsSix,
			SecretSize:  20,
		})
		if err != nil {
			return nil, err
		}

		user.Secret = key.String()
	}
	//生成二维码
	key, err := otp.NewKeyFromURL(user.Secret)
	if err != nil {
		return nil, err
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return nil, err
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	qrcode := base64.StdEncoding.EncodeToString(buf.Bytes())

	secret, err := utils.ExtractSecretFromTOTPURL(user.Secret)
	if err != nil {
		return nil, err
	}

	MFA.Secret = secret
	MFA.Qrcode = qrcode
	return MFA, nil
}

func (uc *UserService) TotpValidate(code, secret string) bool {
	return totp.Validate(code, secret)
}
func (uc *UserService) PasswordHash(password string) (string, error) {
	return crypt.PasswordHash(password)
}
func (uc *UserService) PasswordVerify(password, hash string) bool {
	return crypt.VerifyPassword(password, hash)
}

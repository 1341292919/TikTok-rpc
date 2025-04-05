package service

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/pkg/utils"
	"bytes"
	"encoding/base64"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/png"
)

func OptSecret(user *model.User) (*model.MFA, error) {
	var MFA = &model.MFA{}
	var buf bytes.Buffer

	if user.OptSecret == "" {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "tiktok",
			AccountName: user.UserName,
			Period:      30,
			Digits:      otp.DigitsSix,
			SecretSize:  20,
		})
		if err != nil {
			return nil, err
		}

		user.OptSecret = key.String()
	}
	//生成二维码
	key, err := otp.NewKeyFromURL(user.OptSecret)
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

	secret, err := utils.ExtractSecretFromTOTPURL(user.OptSecret)
	if err != nil {
		return nil, err
	}

	MFA.Secret = secret
	MFA.Qrcode = qrcode
	return MFA, nil
}

func TotpValidate(code, secret string) bool {
	return totp.Validate(code, secret)
}

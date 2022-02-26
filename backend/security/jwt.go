package security

import (
	"github.com/lestrrat-go/jwx/jwa"
	jwt2 "github.com/lestrrat-go/jwx/jwt"
	"notans/backend/common"
	config2 "notans/backend/config"
	"time"
)

type Jwt struct {
	Config *config2.Config
}

func (jwt *Jwt) Create(subject string) string {
	token, err := jwt2.NewBuilder().
		IssuedAt(time.Now()).
		Issuer("Notansv1").
		Subject(subject).
		Build()
	if err != nil {
		common.LogPrintln("JWT::Create::BUILD", err.Error())
		return ""
	}
	sign, err := jwt2.NewSerializer().
		Sign(jwa.HS512, []byte(jwt.Config.AppKey)).
		Serialize(token)

	if err != nil {
		common.LogPrintln("JWT::Create::SIGN", err.Error())
		return ""
	}

	return string(sign)
}

func (jwt *Jwt) Validate(jwtt string) string {
	token, err := jwt2.ParseString(jwtt, jwt2.WithVerify(jwa.HS512, []byte(jwt.Config.AppKey)))
	if err != nil {
		common.LogPrintln("JWT::Validate", err.Error())
		return ""
	}
	return token.Subject()
}

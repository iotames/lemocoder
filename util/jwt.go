package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"
)

func GetJwtBySecret(keyBytes []byte, bodyInfo map[string]interface{}) (string, error) {
	headerInfo := map[string]interface{}{"typ": "JWT", "alg": "HS256"}
	var sig, sstr string
	var err error
	if sstr, err = toJwtString(headerInfo, bodyInfo); err != nil {
		return "", err
	}
	sig = Base64UrlEncode(GetSha256BySecret(sstr, keyBytes))
	return strings.Join([]string{sstr, sig}, "."), nil
}

func toJwtString(headerInfo, bodyInfo map[string]interface{}) (string, error) {
	var err error
	parts := make([]string, 2)
	for i := range parts {
		var jsonValue []byte
		if i == 0 {
			jsonValue, err = json.Marshal(headerInfo)
		} else {
			jsonValue, err = json.Marshal(bodyInfo)
			log.Println("----jwt--toJsonString---:", string(jsonValue))
		}
		if err != nil {
			return "", err
		}

		parts[i] = Base64UrlEncode(jsonValue)
	}

	return strings.Join(parts, "."), nil
}

type JsonWebToken struct {
	secret, TokenString string
}

func NewJwt(secret string) *JsonWebToken {
	return &JsonWebToken{secret: secret}
}

func (j *JsonWebToken) Create(claims map[string]interface{}, expiresin time.Duration) string {
	claims["exp"] = time.Now().Add(expiresin).Unix()
	log.Println("-------jwt---Create--Info:", claims)
	tokenStr, err := GetJwtBySecret([]byte(j.secret), claims)
	if err != nil {
		log.Println("---jwt create error:", err)
	}
	log.Println("-----jwt-----Create----Token:", tokenStr)
	// log.Println("-------jwt.Create----secret=", string(j.secret))
	j.TokenString = tokenStr
	return j.TokenString
}

func JsonDecodeUseNumber(infoBytes []byte, result interface{}) error {
	// 未设置UseNumber长整型会丢失精度
	decoder := json.NewDecoder(bytes.NewReader(infoBytes))
	decoder.UseNumber()
	return decoder.Decode(result)
}

func (j *JsonWebToken) Decode(jwtStr string) (segInfo map[string]interface{}, err error) {
	segTokens := strings.Split(jwtStr, ".")
	if len(segTokens) != 3 {
		err = errors.New("token is not a jwt")
		return
	}
	var infoBytes []byte
	infoBytes, err = Base64UrlDecode(segTokens[1])
	if err != nil {
		return
	}
	// log.Println("---jwt---Decode--jwtStr:", jwtStr)
	// log.Println("---jwt---Decode--original--InfoJson:", string(infoBytes))
	// err = json.Unmarshal(infoBytes, &segInfo)
	JsonDecodeUseNumber(infoBytes, &segInfo)
	// log.Println("---jwt---Decode--Info:", segInfo)
	return
}

func (j *JsonWebToken) Parse(tokenStr string) (result map[string]interface{}, err error) {
	tokenSplit := strings.Split(tokenStr, `.`)
	if len(tokenSplit) != 3 {
		err = errors.New("jwt must split point to len 3")
		return
	}

	bodyInfoBytes, err := Base64UrlDecode(tokenSplit[1])
	if err != nil {
		return
	}
	// log.Println("--jwt---Parse----jwtStr:", tokenStr)
	// log.Println("--jwt---Parse--original--InfoJson=", string(bodyInfoBytes))
	bodyInfo := map[string]interface{}{}
	// 时间戳 int64 转json会变 float64
	// err = json.Unmarshal(bodyInfoBytes, &bodyInfo)
	err = JsonDecodeUseNumber(bodyInfoBytes, &bodyInfo)
	if err != nil {
		return
	}
	// TODO 处理长整型会丢失精度
	exp, ok := bodyInfo["exp"]
	if ok {
		expiredAt, _ := exp.(json.Number).Int64()
		if expiredAt < time.Now().Unix() {
			err = errors.New("jwt token is expired")
			return
		}
	}
	var okToken string
	log.Println("--jwt---Parse---Info=", bodyInfo)
	okToken, err = GetJwtBySecret([]byte(j.secret), bodyInfo)
	if err != nil {
		return
	}

	if okToken != tokenStr {
		err = errors.New("token sign checked fail")
		return
	}
	result = bodyInfo
	err = nil
	return
}

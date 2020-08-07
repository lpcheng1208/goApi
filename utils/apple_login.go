package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"goApi/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//请求苹果用户信息
func GetAppleLoginData(clientId string, clientSecret string, code string) ([]byte, error) {
	params := map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  "",
	}
	form := url.Values{}
	for k, v := range params {
		form.Set(k, v)
	}

	var request *http.Request
	var err error
	if request, err = http.NewRequest("POST", "https://appleid.apple.com/auth/token",
		strings.NewReader(form.Encode())); err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); nil != err {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	return data, err
}

//生成client_secret
func GetAppleSecret(keyId string, teamId string, clientId string, keySecret string) string {
	token := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": keyId,
		},
		Claims: jwt.MapClaims{
			"iss": teamId,
			"iat": time.Now().Unix(),
			// constraint: exp - iat <= 180 days
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"aud": "https://appleid.apple.com",
			"sub": clientId,
		},
		Method: jwt.SigningMethodES256,
	}

	ecdsaKey, _ := AuthKeyFromBytes([]byte(keySecret))
	ss, _ := token.SignedString(ecdsaKey)
	return ss
}

//JWT加密
func AuthKeyFromBytes(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("token: AuthKey must be a valid .p8 PEM file")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, errors.New("token: AuthKey must be of type ecdsa.PrivateKey")
	}

	return pkey, nil
}

//解密苹果返回的data中的id_token中的用户id和客户端post的id否一致
func CheckAppleID(data []byte, id string, err error) bool {
	//fmt.Println("苹果服务器返回信息为：", string(data))
	if err != nil || strings.Contains(string(data), "error") {
		//fmt.Println("APPLE登陆失败[error 1]，请重试或使用其他方式登陆")
		return false
	}
	//fmt.Println("苹果服务器返回信息OK")
	Au := model.AppleAuth{}
	err = json.Unmarshal(data, &Au)
	if err != nil {
		//fmt.Println("APPLE登陆失败[error 2]，请重试或使用其他方式登陆")
		return false
	}
	//fmt.Println("结构体赋值OK", Au)
	var userDecode []byte
	parts := strings.Split(Au.IdToken, ".")
	userDecode, err = jwt.DecodeSegment(parts[1])
	if err != nil {
		//fmt.Println("APPLE登陆失败[error 3]，请重试或使用其他方式登陆")
		return false
	}
	//fmt.Println("Au.IdToken解码OK", userDecode)
	It := model.AppleIdToken{}
	err = json.Unmarshal(userDecode, &It)
	if err != nil {
		//fmt.Println("APPLE登陆失败[error 4]，请重试或使用其他方式登陆")
		return false
	}
	//fmt.Println("结构体赋值OK", It)
	if It.Sub != id {
		//fmt.Println("APPLE登陆失败[error 5]，请重试或使用其他方式登陆")
		return false
	}
	return true
}

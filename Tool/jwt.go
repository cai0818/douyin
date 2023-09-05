package Tool

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Userid int `json:"username"`
	jwt.StandardClaims
}

// 编码
func Encode(userid int) string {
	mySigningKey := []byte("cai")
	c := new(MyClaims)
	c.Userid = userid
	c.StandardClaims = jwt.StandardClaims{
		//		NotBefore: time.Now().Unix() - 60,
		ExpiresAt: time.Now().Unix() + 12*60*60,
		Issuer:    "cai",
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, e := t.SignedString(mySigningKey)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(s)
	return s
}

// 解码
func Code(s string) int {
	mySigningKey := []byte("cai")
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("%s", err)
	}
	//fmt.Println(token.Claims.(*MyClaims).Userid)
	return token.Claims.(*MyClaims).Userid
}

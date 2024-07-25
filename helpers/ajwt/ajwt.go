package ajwt

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/alaa2amz/g1/config"
)

var(
	 Alg= jwt.SigningMethodHS256
	 Key= config.JWTKey
 )

func Token(mc jwt.MapClaims ) (string,error) {
	tokenStruct :=jwt.NewWithClaims(Alg,mc)
	tokenString,err := tokenStruct.SignedString(Key)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

func Valid(token string) (*jwt.Token,error) {
	tokenStruct,err:= jwt.Parse(token,func (token *jwt.Token) (interface{}, error){
		// Don't forget to validate the alg is what you expect:
        	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        	}
        	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
        	return []byte(Key), nil
	})
	if err != nil {
		return nil,err
	}
	return  tokenStruct,nil
}


func key(token *jwt.Token) (interface{}, error){
// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(Key), nil
}

func EasyClaims(name,role string,exp int) jwt.MapClaims {
return  jwt.MapClaims{
                "sub": name,                    // Subject (user identifier)
                "iss": "alaazak",                  // Issuer
                "aud": role,           // Audience (user role)
                "exp": time.Now().Add(time.Second*60*2).Unix(), // Expiration time
                "iat": time.Now().Unix(),                 // Issued at
        }}

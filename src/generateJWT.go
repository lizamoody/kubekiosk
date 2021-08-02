package main

import (
	"fmt"

	"os"

	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

/*type custClaims struct {
	Issuer  string   `json:"iss,omitempty"`
	Subject string   `json:"sub,omitempty"`
	Groups  []string `json:"groups,omitempty`
	//Expiry    *NumericDate `json:"exp,omitempty"`
	//NotBefore *NumericDate `json:"nbf,omitempty"`
	//IssuedAt  *NumericDate `json:"iat,omitempty"`
	ID string `json:"jti,omitempty"`
} */

//– log in,

//create json web token payload,
//and then that info should ideally have group membership so that doen’t need to look it up again,
//that get’s passed back to user and then on end-point just need to
//validate that token is still validated (authenticated by proper key and still valid) and
//verify that payload is correct
func generateJWT(user string) (token string) {
	username := user

	//public key
	key := []byte(os.Getenv("ACCESS_SECRET"))

	//sign token
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}

	//generate jwt with claims
	cl := jwt.Claims{
		Subject: username,
		Issuer:  "BL",
		//	NotBefore: jwt.NewNumericDate(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}

	//sign token
	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		panic(err)
	}
	token = raw
	fmt.Println("token " + raw)

	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		panic(err)
	}

	out := jwt.Claims{}
	if err := tok.Claims(key, &out); err != nil {
		panic(err)
	}
	//	fmt.Print("subject: %s, Issuer: %s\n", out.Subject, out.Issuer)

	//validate token
	claims := jwt.Claims{}
	if err := tok.Claims(key, &claims); err != nil {
		panic(err)
	}

	err = claims.Validate(jwt.Expected{
		Subject: username,
		Issuer:  "BL",
		//NotBefore: jwt.NewNumericDate(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	})
	if err != nil {
		panic(err)

	}
	fmt.Printf("valid!")
	return token
	/*

		//generate key pair
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		fmt.Println(privateKey)
		if err != nil {
			panic(err)
		}

		//instantiate a signer using RS256 with private key
		signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)
		if err != nil {
			panic(err)
		}

		//Sign payload and return JWS object
		var payload = []byte(os.Getenv("ACCESS_SECRET"))
		object, err := signer.Sign(payload)
		if err != nil {
			panic(err)
		}

		//serialize the encrypted object using the full serialization format
		serialized := object.FullSerialize()

		//Parse serialized object, protected JWS obj
		object, err = jose.ParseSigned(serialized)
		if err != nil {
			panic(err)
		}

		//verify signature on the payload
		output, err := object.Verify(&privateKey.PublicKey)
		if err != nil {
			panic(err)
		}

		fmt.Print("output: " + string(output))*/
}

func readJWT(token string) {
	key := []byte(os.Getenv("ACCESS_SECRET"))
	tok, err := jwt.ParseSigned(token)
	if err != nil {
		panic(err)
	}
	out := jwt.Claims{}
	if err := tok.Claims(key, &out); err != nil {
		panic(err)
	}

	fmt.Println("Parsed user: " + out.Subject)
	fmt.Println("Parsed issuer: " + out.Issuer)

	/*
		fmt.Print("subject: %s, Issuer: %s\n", out.Subject, out.Issuer)
		claims := jwt.Claims{}
		if err := tok.Claims(key, &claims); err != nil {
			panic(err)
		}
		fmt.Println("Parsed username: ")
		err = claims.Validate(jwt.Expected{
			Subject: username,
			Issuer:  "BL",
			//NotBefore: jwt.NewNumericDate(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		})
		if err != nil {
			panic(err)

		}*/

}

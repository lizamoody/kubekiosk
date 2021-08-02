package server
import(
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
  )
  // For testing create the RSA key pair in the code
  	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
  	}
  	// create Square.jose signing key
  	key := jose.SigningKey{Algorithm: jose.RS256, Key: rsaPrivateKey}
  
  	// create a Square.jose RSA signer, used to sign the JWT
  	var signerOpts = jose.SignerOptions{}
  	signerOpts.WithType("JWT")
 	rsaSigner, err := jose.NewSigner(key, &signerOpts)
	if err != nil {
	  log.Fatalf("failed to create signer:%+v", err)
  	}

  	// create an instance of Builder that uses the rsa signer
	builder := jwt.Signed(rsaSigner)

	// Claims represents public claim values (as specified in RFC 7519).
type Claims struct {
	Issuer    string      `json:"iss,omitempty"`
	Subject   string      `json:"sub,omitempty"`
	Audience  Audience    `json:"aud,omitempty"`
	Expiry    NumericDate `json:"exp,omitempty"`
	NotBefore NumericDate `json:"nbf,omitempty"`
	IssuedAt  NumericDate `json:"iat,omitempty"`
	ID        string      `json:"jti,omitempty"`
}
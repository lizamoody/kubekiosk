package server

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base32"
	"fmt"
	"log"
	"math/big"
	"time"

	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

// Variables
var (
	// Length of the refresh token in characters
	RefreshTokenLength = 32
	AuthCodeLength     = 32
)

//Token Claims Struct

type idTokenClaims struct {
	Issuer  string `json:"iss"`
	Subject string `json:"sub"`
	//TO DO:: create audience object
	//audience:         jwt.Audience{"bl", "hi"}, `json:"aud"`

	Expiry           int64  `json:"exp"`
	IssuedAt         int64  `json:"iat"`
	AuthorizingParty string `json:"azp,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
}

// Determine the signature algorithm for a JWT.
func signatureAlgorithm(jwk *jose.JSONWebKey) (alg jose.SignatureAlgorithm, err error) {
	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "issuer",
		NotBefore: jwt.NewNumericDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
		Expiry:    jwt.NewNumericDate(time.Date(2016, 1, 1, 0, 15, 0, 0, time.UTC)),
		Audience:  jwt.Audience{"leela", "fry"},
	}
	log.Println(cl.Audience)

	rsa, _ := rsa.GenerateKey(rand.Reader, 2048)
	serialNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	if jwk.Key == nil {
		return alg, fmt.Errorf("no signing key")
	}
	switch key := jwk.Key.(type) {
	case *rsa.PrivateKey:
		// Because OIDC mandates that we support RS256, we always return that
		// value. In the future, we might want to make this configurable on a
		// per client basis. For example allowing PS256 or ECDSA variants.
		//
		// See https://github.com/dexidp/dex/issues/692
		return jose.RS256, nil
	case *ecdsa.PrivateKey:
		// We don't actually support ECDSA keys yet, but they're tested for
		// in case we want to in the future.
		//
		// These values are prescribed depending on the ECDSA key type. We
		// can't return different values.
		switch key.Params() {
		case elliptic.P256().Params():
			return jose.ES256, nil
		case elliptic.P384().Params():
			return jose.ES384, nil
		case elliptic.P521().Params():
			return jose.ES512, nil
		default:
			return alg, fmt.Errorf("unsupported ecdsa curve")
		}
	default:
		return alg, fmt.Errorf("unsupported signing key type %T", key)
	}
}

func signPayload(key *jose.JSONWebKey, alg jose.SignatureAlgorithm, payload []byte) (jws string, err error) {
	var jwk2 jose.JSONWebKey
	jwk2.Key = "dfgd"

	jwk := jose.JSONWebKey{
		Key:                         testCertificates[0].PublicKey,
		KeyID:                       "bar",
		Algorithm:                   "foo",
		Certificates:                testCertificates,
		CertificateThumbprintSHA1:   x5tSHA1[:],
		CertificateThumbprintSHA256: x5tSHA256[:],
	}

	jwk.Key = "hi"

	if err == nil {
		t.Error("should not marshal JWK with too short thumbprints")
	}

	signingKey := jose.SigningKey{Key: key, Algorithm: alg}

	signer, err := jose.NewSigner(signingKey, &jose.SignerOptions{})
	if err != nil {
		return "", fmt.Errorf("new signer: %v", err)
	}
	signature, err := signer.Sign(payload)
	if err != nil {
		return "", fmt.Errorf("signing payload: %v", err)
	}
	return signature.CompactSerialize()
}

func newIDToken(ctx context.Context, req *accessTokenRequest) (idToken string, expiry time.Time, err error) {
	keys, err := s.readKeys(ctx)

	//if key is not found
	if err != nil {
		log.Errorf("Failed to get keys: %v", err)
		return "", expiry, err
	}

	signingKey := keys.SigningKey

	//if no key to sign payload
	if signingKey == nil {
		return "", expiry, fmt.Errorf("no key to sign payload with")
	}

	//TO DO:: need to finish this function
	signingAlg, err := signatureAlgorithm(signingKey)
	if err != nil {
		return "", expiry, err
	}
	issuedAt := s.now()

}

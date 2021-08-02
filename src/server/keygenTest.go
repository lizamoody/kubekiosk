package server

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

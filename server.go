package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Response string `json:"response"`
}

func GetData(w http.ResponseWriter, r *http.Request) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)

	publicKey := privateKey.PublicKey

	var reqData Request

	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		CheckError(err)
	}

	resp := RSA_OAEP_Encrypt(reqData.Message, publicKey)
	jsonResponse := Response{
		Response: resp,
	}

	json, err := json.Marshal(jsonResponse)
	if err != nil {
		CheckError(err)
	}

	w.Write([]byte(json))

}

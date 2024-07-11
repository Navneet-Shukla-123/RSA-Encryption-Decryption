package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Response string `json:"response"`
}

var CipherText string

var PrivateKey *rsa.PrivateKey
var PublicKey rsa.PublicKey

func init() {
	var err error
	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)

	PublicKey = PrivateKey.PublicKey
}

func GetData(w http.ResponseWriter, r *http.Request) {

	var reqData Request

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		CheckError(err)
	}

	resp := RSA_OAEP_Encrypt(reqData.Message, PublicKey)
	jsonResponse := Response{
		Response: resp,
	}

	json, err := json.Marshal(jsonResponse)
	if err != nil {
		CheckError(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(json))

}

func PostData(w http.ResponseWriter, r *http.Request) {

	var reqData Request

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		CheckError(err)
	}

	log.Println("Reqdata is ", reqData.Message)

	decryptedData := RSA_OAEP_Decrypt(reqData.Message, *PrivateKey)
	resp := Response{
		Response: decryptedData,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		CheckError(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(json))

}

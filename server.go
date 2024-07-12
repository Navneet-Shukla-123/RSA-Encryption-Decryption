package main

import (
	"channels/redis"
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

type RedisResponse struct {
	Message string `json:"message"`
	Key     string `json:"key"`
	Value   string `json:"value"`
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
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		log.Println("Error in reading the response body ", err)
		w.Write([]byte("error in reading the response body"))
		return
	}

	resp := RSA_OAEP_Encrypt(reqData.Message, PublicKey)
	jsonResponse := Response{
		Response: resp,
	}

	json, err := json.Marshal(jsonResponse)
	if err != nil {
		log.Println("error in converting to json ", err)
		w.Write([]byte("error in convertinng to json"))
		return
	}

	w.Write([]byte(json))

}

func PostData(w http.ResponseWriter, r *http.Request) {

	var reqData Request
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		log.Println("Error in reading the request body ", err)
		w.Write([]byte("error in reading the request body"))
		return
	}

	log.Println("Reqdata is ", reqData.Message)

	decryptedData := RSA_OAEP_Decrypt(reqData.Message, *PrivateKey)
	resp := Response{
		Response: decryptedData,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error in converting to json ",err)
		w.Write([]byte("error in converting to json"))
		return
	}

	log.Println("Message from redis is ",reqData.Message)
	err = redis.InsertIntoRedis(decryptedData, reqData.Message)
	if err != nil {
		log.Println("Error in inserting to redis ", err)
		w.Write([]byte("Error in inserting to redis"))
		return
	}
	log.Println("Data successfully inserted to Redis")
	w.Write([]byte(json))

}

func GetFromRedis(w http.ResponseWriter, r *http.Request) {
	var req Request

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error reading the request body ", err)
		w.Write([]byte("Error reading the request body"))
		return
	}

	redisData, err := redis.GetFromDB(req.Message)
	if err != nil {
		log.Println("Error in getting the data from redis ", err)
		w.Write([]byte("Error getting data from redis"))
		return
	}
	resp := RedisResponse{
		Message: "value fetched from redis successfully",
		Key:     req.Message,
		Value:   redisData,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error in converting response to json ", err)
		w.Write([]byte("error in converting response to json"))
		return
	}

	w.Write(json)
}

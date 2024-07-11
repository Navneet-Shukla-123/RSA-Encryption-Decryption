package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func RSA_OAEP_Encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	CheckError(err)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
func CheckError(e error) {
	if e != nil {
		log.Println(e)
	}
}
func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)

	publicKey := privateKey.PublicKey
	secretMessage := "This is super secret  message!"

	encryptedMessage := RSA_OAEP_Encrypt(secretMessage, publicKey)

	log.Println("Cipher Text:", encryptedMessage)

	RSA_OAEP_Decrypt(encryptedMessage, *privateKey)

}

func RSA_OAEP_Decrypt(cipherText string, privKey rsa.PrivateKey) string {

	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label:=[]byte("OAEP Encrypted")

	rng:=rand.Reader
	plainText,err:=rsa.DecryptOAEP(sha256.New(),rng,&privKey,ct,label)
	CheckError(err)

	fmt.Println("PlainText:",string(plainText))
	return string(plainText)

}

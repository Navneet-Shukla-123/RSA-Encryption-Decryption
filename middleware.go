package main

import "net/http"

func DecryptRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		RSA_Encrypt := r.Header.Get("hash")

		dcrypt := RSA_OAEP_Decrypt(RSA_Encrypt, *PrivateKey)
		r.Header.Add("decrypt", dcrypt)
		f(w,r)
	}
}

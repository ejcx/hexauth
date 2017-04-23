package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jose "gopkg.in/square/go-jose.v2"
)

var (
	MagicPublicKey *jose.JSONWebKey
	DevJWK         = "http://localhost:8000/magic/pubkey"
	err            error
)

type MagicCallback struct {
	ToEmail        string
	CustomerSecret string
	CompanyName    string
	Verified       bool

	SentTime     time.Time
	VerifiedTime time.Time
}

func init() {
	MagicPublicKey, err = GetJWKByURL(DevJWK)
	if err != nil {
		log.Fatalf("Could not fetch jwk: %s\n", err)
	}
}

func GetJWKByURL(url string) (jwk *jose.JSONWebKey, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jwk = new(jose.JSONWebKey)
	err = jwk.UnmarshalJSON(b)
	return jwk, err
}
func HandleVerify(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("key")
	if k == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	obj, err := jose.ParseSigned(k)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	plaintext, err := obj.Verify(MagicPublicKey)

	if err != nil {
		log.Printf("Could not verify the plaintext: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(string(plaintext))
}

func (m *MagicCallback) GetEmail() string {
	return m.ToEmail
}
func main() {
	http.HandleFunc("/", HandleVerify)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

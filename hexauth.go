package hexauth

import (
	"fmt"
	"net/http"
	"time"
)

type MagicCallback struct {
	ToEmail        string
	CustomerSecret string
	CompanyName    string
	Verified       bool

	SentTime     time.Time
	VerifiedTime time.Time
}

func HandleVerify(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("key")
	fmt.Println(k)
}

func (m *MagicCallback) GetEmail() string {
	return m.ToEmail
}

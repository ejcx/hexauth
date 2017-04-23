package hexauth

import "testing"

func TestGetJWK(t *testing.T) {
	pub, err := GetJWKByURL(DevJWK)
	if err != nil {
		t.Errorf("Could not fetch jwk: %s", err)
	}
	if !pub.Valid() {
		t.Errorf("Public key is not valid")
	}
}

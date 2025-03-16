package encrypt

import "testing"

func TestEncryptMobile(t *testing.T) {
	mobile := "13800138000"
	encryptedMobile, err := EnMobile(mobile)
	if err != nil {
		t.Fatal(err)
	}
	decryptedMobile, err := DeMobile(encryptedMobile)
	if err != nil {
		t.Fatal(err)
	}
	if decryptedMobile != mobile {
		t.Fatalf("expected %s, but got %s", mobile, decryptedMobile)
	}
}

func TestEncryptPassword(t *testing.T) {
	password := "hufhaihfiuahf"
	hashedPassword, err := EnPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	if !CheckPassword(password, hashedPassword) {
		t.Fatal("password not match")
	}
}

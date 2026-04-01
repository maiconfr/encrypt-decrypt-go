package functions

import (
	"testing"
)

var testBytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var testSecret = "1234567890123456"

func TestEncrypt(t *testing.T) {
	originalText := "Encrypting this string"
	
	encryptedText, err := Encrypt(originalText, testSecret, testBytes)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}
	
	if encryptedText == "" {
		t.Fatal("Encrypted text is empty")
	}
}

func TestDecrypt(t *testing.T) {
	originalText := "Encrypting this string"
	
	// First encrypt
	encryptedText, err := Encrypt(originalText, testSecret, testBytes)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}
	
	// Then decrypt
	decryptedText, err := Decrypt(encryptedText, testSecret, testBytes)
	if err != nil {
		t.Fatalf("Error decrypting: %v", err)
	}
	
	if decryptedText != originalText {
		t.Errorf("Decrypted text does not match original. Original: %s, Decrypted: %s", originalText, decryptedText)
	}
}

func TestEncode(t *testing.T) {
	data := []byte("test data")
	encoded := Encode(data)
	
	if encoded == "" {
		t.Fatal("Encode returned empty string")
	}
	
	// Test if it's valid base64
	decoded := Decode(encoded)
	if string(decoded) != string(data) {
		t.Errorf("Decode/Encode does not match. Original: %s, Decoded: %s", string(data), string(decoded))
	}
}

func TestDecode(t *testing.T) {
	// Test valid base64
	validBase64 := "dGVzdCBkYXRh" // "test data" in base64
	decoded := Decode(validBase64)
	
	if string(decoded) != "test data" {
		t.Errorf("Decode did not work correctly. Expected: 'test data', Received: '%s'", string(decoded))
	}
	
	// Test invalid base64 - this will panic
	defer func() {
		if r := recover(); r != nil {
			t.Log("Expected panic for invalid base64:", r)
		}
	}()
	
	Decode("invalid-base64!")
}

func TestEncryptWithInvalidKey(t *testing.T) {
	invalidKey := "123"
	
	_, err := Encrypt("test", invalidKey, testBytes)
	if err == nil {
		t.Error("Expected error with invalid key, but none occurred")
	}
}

func TestDecryptWithInvalidKey(t *testing.T) {
	invalidKey := "123"
	
	_, err := Decrypt("someBase64String", invalidKey, testBytes)
	if err == nil {
		t.Error("Expected error with invalid key, but none occurred")
	}
}

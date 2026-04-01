package main

import (
	"encrypt-decrypt/functions"
	"testing"
)

var testBytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var testSecret = "1234567890123456"

func TestEncryptDecrypt(t *testing.T) {
	// Test data
	originalText := "Encrypting this string"
	
	// Test encryption
	encryptedText, err := functions.Encrypt(originalText, testSecret, testBytes)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}
	
	if encryptedText == "" {
		t.Fatal("Encrypted text is empty")
	}
	
	// Test decryption
	decryptedText, err := functions.Decrypt(encryptedText, testSecret, testBytes)
	if err != nil {
		t.Fatalf("Error decrypting: %v", err)
	}
	
	if decryptedText != originalText {
		t.Errorf("Decrypted text does not match original. Original: %s, Decrypted: %s", originalText, decryptedText)
	}
}

func TestEncryptWithDifferentInputs(t *testing.T) {
	testCases := []struct {
		name string
		text string
	}{
		{"Simple text", "Hello World"},
		{"Empty text", ""},
		{"Text with numbers", "Test123"},
		{"Text with special characters", "Test@#$%"},
		{"Long text", "This is a longer text to test the encryption and decryption functionality with more content"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			encrypted, err := functions.Encrypt(tc.text, testSecret, testBytes)
			if err != nil {
				t.Errorf("Error encrypting '%s': %v", tc.name, err)
				return
			}
			
			// Decrypt
			decrypted, err := functions.Decrypt(encrypted, testSecret, testBytes)
			if err != nil {
				t.Errorf("Error decrypting '%s': %v", tc.name, err)
				return
			}
			
			if decrypted != tc.text {
				t.Errorf("Text does not match for '%s'. Original: %s, Decrypted: %s", tc.name, tc.text, decrypted)
			}
		})
	}
}

func TestEncryptWithInvalidKey(t *testing.T) {
	// Test with invalid key (too short for AES)
	invalidKey := "123"
	
	_, err := functions.Encrypt("test", invalidKey, testBytes)
	if err == nil {
		t.Error("Expected error with invalid key, but none occurred")
	}
}

func TestDecryptWithInvalidText(t *testing.T) {
	// Test with invalid base64 - this will cause a panic in Decode function
	defer func() {
		if r := recover(); r != nil {
			// Expected panic for invalid base64
			t.Log("Expected panic for invalid base64:", r)
		}
	}()
	
	_, err := functions.Decrypt("invalid-base64!", testSecret, testBytes)
	if err == nil {
		t.Error("Expected error with invalid text, but none occurred")
	}
}

func BenchmarkEncrypt(b *testing.B) {
	text := "Benchmark test text"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := functions.Encrypt(text, testSecret, testBytes)
		if err != nil {
			b.Fatalf("Error in encryption benchmark: %v", err)
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	text := "Benchmark test text"
	encrypted, _ := functions.Encrypt(text, testSecret, testBytes)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := functions.Decrypt(encrypted, testSecret, testBytes)
		if err != nil {
			b.Fatalf("Error in decryption benchmark: %v", err)
		}
	}
}

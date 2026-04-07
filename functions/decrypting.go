package functions

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"time"

	"encrypt-decrypt/telemetry"
	"go.opentelemetry.io/otel/attribute"
)

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func Decrypt(ctx context.Context, text, MySecret string, bytes []byte) (string, error) {
	// Create telemetry span
	_, span := telemetry.CreateSpan(ctx, "encryption.decrypt")
	defer span.End()

	// Add attributes
	telemetry.AddSpanAttributes(ctx,
		attribute.String("encryption.algorithm", "AES"),
		attribute.String("encryption.mode", "CFB"),
		attribute.Int("cipher_text.length", len(text)),
	)

	startTime := time.Now()

	// Add event for decryption start
	telemetry.AddSpanEvent(ctx, "decryption.started",
		attribute.String("timestamp", startTime.Format(time.RFC3339)),
	)

	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		telemetry.RecordError(ctx, err)
		return "", err
	}

	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	result := string(plainText)

	// Record metrics
	duration := time.Since(startTime)
	telemetry.AddSpanAttributes(ctx,
		attribute.Int("plain_text.length", len(result)),
		attribute.Float64("decryption.duration_ms", float64(duration.Nanoseconds())/1e6),
	)

	// Add event for decryption completion
	telemetry.AddSpanEvent(ctx, "decryption.completed",
		attribute.String("timestamp", time.Now().Format(time.RFC3339)),
		attribute.Float64("duration_ms", float64(duration.Nanoseconds())/1e6),
		attribute.String("plain_text.prefix", getTextPrefix(result)),
	)

	return result, nil
}

// getTextPrefix returns a safe prefix of text for logging
func getTextPrefix(text string) string {
	if len(text) <= 10 {
		return text
	}
	return text[:10] + "..."
}

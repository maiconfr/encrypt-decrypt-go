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

// Encrypt method is to encrypt or hide any classified text
func Encrypt(ctx context.Context, text, MySecret string, bytes []byte) (string, error) {
	// Create telemetry span
	_, span := telemetry.CreateSpan(ctx, "encryption.encrypt")
	defer span.End()

	// Add attributes
	telemetry.AddSpanAttributes(ctx,
		attribute.String("encryption.algorithm", "AES"),
		attribute.String("encryption.mode", "CFB"),
		attribute.Int("text.length", len(text)),
		attribute.String("text.prefix", getTextPrefix(text)),
	)

	startTime := time.Now()

	// Add event for encryption start
	telemetry.AddSpanEvent(ctx, "encryption.started",
		attribute.String("timestamp", startTime.Format(time.RFC3339)),
	)

	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		telemetry.RecordError(ctx, err)
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	result := Encode(cipherText)

	// Record metrics
	duration := time.Since(startTime)
	telemetry.AddSpanAttributes(ctx,
		attribute.Int("cipher_text.length", len(result)),
		attribute.Float64("encryption.duration_ms", float64(duration.Nanoseconds())/1e6),
	)

	// Add event for encryption completion
	telemetry.AddSpanEvent(ctx, "encryption.completed",
		attribute.String("timestamp", time.Now().Format(time.RFC3339)),
		attribute.Float64("duration_ms", float64(duration.Nanoseconds())/1e6),
	)

	return result, nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

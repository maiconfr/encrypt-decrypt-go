package main

import (
	"context"
	"encrypt-decrypt/functions"
	"encrypt-decrypt/telemetry"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel/attribute"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

var MySecret string

func main() {
	// Initialize OpenTelemetry
	cleanup, err := telemetry.InitProvider("encrypt-decrypt-service")
	if err != nil {
		log.Printf("Failed to initialize OpenTelemetry: %v", err)
	} else {
		defer cleanup()
	}

	// Handle graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Create main span
	ctx, span := telemetry.CreateSpan(ctx, "main.application")
	defer span.End()

	// Add application attributes
	telemetry.AddSpanAttributes(ctx,
		attribute.String("application.name", "encrypt-decrypt"),
		attribute.String("application.version", "1.0.0"),
		attribute.String("go.version", "1.25.0"),
	)

	err = godotenv.Load(".env")
	if err != nil {
		telemetry.RecordError(ctx, err)
		log.Println("Adicionar o arquivo de configuração .env \n", err.Error())
		panic(err)
	}

	MySecret = os.Getenv("MySecret")

	// Add event for configuration loaded
	telemetry.AddSpanEvent(ctx, "configuration.loaded",
		attribute.String("timestamp", time.Now().Format(time.RFC3339)),
		attribute.Bool("secret.loaded", MySecret != ""),
	)

	StringToEncrypt := "Encrypting this string"
	
	// To encrypt the StringToEncrypt
	encText, err := functions.Encrypt(ctx, StringToEncrypt, MySecret, bytes)
	if err != nil {
		telemetry.RecordError(ctx, err)
		fmt.Println("error encrypting your classified text: ", err)
	} else {
		fmt.Println(encText)
		
		// Add event for successful encryption
		telemetry.AddSpanEvent(ctx, "encryption.successful",
			attribute.String("encrypted_text", encText),
			attribute.String("original_text", StringToEncrypt),
		)
	}

	// To decrypt the original StringToEncrypt
	decText, err := functions.Decrypt(ctx, "Li5E8RFcV/EPZY/neyCXQYjrfa/atA==", MySecret, bytes)
	if err != nil {
		telemetry.RecordError(ctx, err)
		fmt.Println("error decrypting your encrypted text: ", err)
	} else {
		fmt.Println(decText)
		
		// Add event for successful decryption
		telemetry.AddSpanEvent(ctx, "decryption.successful",
			attribute.String("decrypted_text", decText),
		)
	}

	// Add final application event
	telemetry.AddSpanEvent(ctx, "application.completed",
		attribute.String("timestamp", time.Now().Format(time.RFC3339)),
	)
}

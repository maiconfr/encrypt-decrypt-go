# Encrypt-Decrypt Go

A simple Go project to demonstrate text encryption and decryption using AES (Advanced Encryption Standard) with CFB (Cipher Feedback) mode.

## Features

- **Encryption**: Encrypts text using AES-256 with CFB mode
- **Decryption**: Decrypts text back to its original format
- **Base64 Encoding**: Encodes encrypted result in Base64 for secure transport

## Project Structure

```
encrypt-decrypt-go/
├── main.go              # Main file with usage example
├── functions/
│   ├── encrypting.go    # Encryption functions
│   └── decrypting.go    # Decryption functions
├── go.mod              # Go module
├── go.sum              # Dependencies
└── .env                # Environment variables (not included in repository)
```

## Prerequisites

- Go 1.19 or higher
- `MySecret` environment variable configured in `.env` file

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd encrypt-decrypt-go
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file with your secret key:
```
MySecret=your-secret-key-here-must-be-32-bytes
```

## Usage

### Run the example

```bash
go run main.go
```

The program will:
1. Encrypt the string "Encrypting this string"
2. Display the encrypted text
3. Decrypt a hardcoded example
4. Display the decrypted text

### Use the functions in your code

```go
package main

import (
    "encrypt-decrypt/functions"
    "fmt"
)

func main() {
    // Secret key (32 bytes for AES-256)
    secret := "your-32-byte-secret-key-exactly-32"
    
    // IV (Initialization Vector) - should be unique for each encryption
    iv := []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
    
    // Encrypt
    original := "Secret text"
    encrypted, err := functions.Encrypt(original, secret, iv)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Encrypted: %s\n", encrypted)
    
    // Decrypt
    decrypted, err := functions.Decrypt(encrypted, secret, iv)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

## Technical Details

- **Algorithm**: AES-256 (Advanced Encryption Standard)
- **Mode**: CFB (Cipher Feedback)
- **Encoding**: Base64 for secure result representation
- **Key**: Requires exactly 32 bytes for AES-256
- **IV**: 16 bytes for AES in CFB mode

## Security

- The secret key must be exactly 32 characters/bytes
- The IV (Initialization Vector) should be unique for each encryption operation
- Never share the secret key or IV
- For production, consider generating random IVs for each encryption

## Compilation

To compile the executable:

```bash
go build -o encrypt-decrypt main.go
```

## License

This project is for educational and demonstration purposes only.

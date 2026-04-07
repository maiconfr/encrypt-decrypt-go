# Encrypt-Decrypt Go

A simple Go project to demonstrate text encryption and decryption using AES (Advanced Encryption Standard) with CFB (Cipher Feedback) mode.

## Features

- **Encryption**: Encrypts text using AES-256 with CFB mode
- **Decryption**: Decrypts text back to its original format
- **Base64 Encoding**: Encodes encrypted result in Base64 for secure transport
- **OpenTelemetry Integration**: Comprehensive observability with traces and metrics
- **Performance Monitoring**: Built-in timing and performance metrics
- **Comprehensive Testing**: 100% code coverage with unit tests and benchmarks
- **Performance Optimized**: Benchmarks showing 500K+ operations per second
- **Automation Workflows**: Git and testing automation for development workflow
- **Error Handling**: Robust error handling for invalid inputs and edge cases

## Project Structure

```
encrypt-decrypt-go/
├── main.go              # Main file with usage example
├── main_test.go         # Unit tests and benchmarks for main package
├── functions/
│   ├── encrypting.go    # Encryption functions with telemetry
│   ├── decrypting.go    # Decryption functions with telemetry
│   └── functions_test.go # Unit tests with 100% coverage
├── telemetry/
│   └── telemetry.go     # OpenTelemetry configuration and utilities
├── go.mod              # Go module
├── go.sum              # Dependencies
├── .env                # Environment variables (not included in repository)
├── .env.example        # Example environment variables
├── .gitignore          # Git ignore rules
└── .windsurf/
    └── workflows/      # Automation workflows
        ├── git-commit.md  # Git commit workflow
        └── run-tests.md   # Testing workflow
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
```bash
cp .env.example .env
# Edit .env with your configuration
```


## OpenTelemetry Observability

This application includes comprehensive OpenTelemetry integration for monitoring and observability.

### Telemetry Features

- **Distributed Tracing**: End-to-end tracing of encryption and decryption operations
- **Performance Metrics**: Automatic timing and duration measurements
- **Event Logging**: Detailed events for each operation step
- **Error Tracking**: Automatic error recording and span status updates
- **Attribute Enrichment**: Contextual metadata for better observability

### Configuration

The telemetry can be configured via environment variables:

```bash
# Disable OpenTelemetry (useful for production or testing)
OTEL_SDK_DISABLED=true

# Custom OTLP endpoint (for production monitoring)
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
OTEL_EXPORTER_OTLP_HEADERS=authorization=Bearer your-token
```

### Telemetry Output

By default, telemetry data is exported to stdout in JSON format. Each execution produces:

- **Spans**: Complete trace of operations with timing
- **Events**: Key moments during encryption/decryption
- **Attributes**: Metadata like algorithm, text lengths, durations
- **Metrics**: Performance measurements and error rates

Example output:
```json
{
  "Name": "encryption.encrypt",
  "SpanKind": "Internal",
  "StartTime": "2026-04-01T11:45:28.914883361-03:00",
  "EndTime": "2026-04-01T11:45:28.914920369-03:00",
  "Attributes": [
    {"Key": "encryption.algorithm", "Value": {"Type": "STRING", "Value": "AES"}},
    {"Key": "encryption.duration_ms", "Value": {"Type": "FLOAT64", "Value": 0.037008}}
  ]
}
```

### Integration with Monitoring Systems

For production use, configure OpenTelemetry to export to your monitoring system:

```bash
# Jaeger
OTEL_EXPORTER_JAEGER_ENDPOINT=http://jaeger:14250/api/traces

# Prometheus
OTEL_EXPORTER_PROMETHEUS_ENDPOINT=http://prometheus:9090

# Generic OTLP
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317
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

## Testing

This project includes a comprehensive test suite with unit tests and benchmarks.

### Run All Tests

```bash
# Run tests with verbose output
go test -v

# Run tests with coverage
go test -v -cover

# Run benchmarks
go test -bench=.

# Generate HTML coverage report
go test -coverprofile=coverage.out ./functions/ && go tool cover -html=coverage.out
```

### Test Results

**Unit Tests:**
- ✅ **Main Package**: 4 tests passing
  - `TestEncryptDecrypt`: Basic encryption/decryption functionality
  - `TestEncryptWithDifferentInputs`: Multiple input scenarios
  - `TestEncryptWithInvalidKey`: Error handling for invalid keys
  - `TestDecryptWithInvalidText`: Error handling for invalid base64

- ✅ **Functions Package**: 6 tests passing (100% coverage)
  - `TestEncrypt`: Encryption functionality
  - `TestDecrypt`: Decryption functionality
  - `TestEncode`: Base64 encoding
  - `TestDecode`: Base64 decoding
  - `TestEncryptWithInvalidKey`: Invalid key handling
  - `TestDecryptWithInvalidKey`: Invalid key handling for decryption

**Performance Benchmarks:**
- **Encryption**: 542,156 operations/sec (3,289 ns/op)
- **Decryption**: 508,561 operations/sec (2,350 ns/op)

**Code Coverage:**
- **Functions Package**: 100% statement coverage
- **Main Package**: Integration tests for workflow validation

### Test Scenarios

The test suite covers:
- ✅ Basic encryption and decryption workflows
- ✅ Multiple input types (empty, numbers, special characters, long text)
- ✅ Error handling for invalid keys (wrong size)
- ✅ Error handling for invalid base64 data
- ✅ Performance benchmarks for production readiness
- ✅ Edge cases and boundary conditions

### Automation Workflows

This project includes automated workflows for testing:

#### `/run-tests` Workflow
Executes complete test suite including:
- Dependency verification (`go mod tidy`)
- Unit tests with detailed output
- Performance benchmarks
- Code coverage analysis with HTML report

#### `/git-commit` Workflow
Automated Git operations with:
- Repository status checking
- Staging area management
- Commit with descriptive messages
- Optional push to remote repository

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

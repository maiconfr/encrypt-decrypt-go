# Encrypt-Decrypt Go

Um projeto Go simples para demonstrar criptografia e descriptografia de texto usando AES (Advanced Encryption Standard) com modo CFB (Cipher Feedback).

## Funcionalidades

- **Criptografia**: Criptografa texto usando AES-256 com modo CFB
- **Descriptografia**: Descriptografa o texto de volta ao seu formato original
- **Base64 Encoding**: Codifica o resultado criptografado em Base64 para transporte seguro

## Estrutura do Projeto

```
encrypt-decrypt-go/
├── main.go              # Arquivo principal com exemplo de uso
├── functions/
│   ├── encrypting.go    # Funções de criptografia
│   └── decrypting.go    # Funções de descriptografia
├── go.mod              # Módulo Go
├── go.sum              # Dependências
└── .env                # Variáveis de ambiente (não incluído no repositório)
```

## Pré-requisitos

- Go 1.19 ou superior
- Variável de ambiente `MySecret` configurada no arquivo `.env`

## Instalação

1. Clone o repositório:
```bash
git clone <repositório-url>
cd encrypt-decrypt-go
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Crie um arquivo `.env` com sua chave secreta:
```
MySecret=sua-chave-secreta-aqui-deve-ter-32-bytes
```

## Uso

### Executar o exemplo

```bash
go run main.go
```

O programa irá:
1. Criptografar a string "Encrypting this string"
2. Exibir o texto criptografado
3. Descriptografar um exemplo hardcoded
4. Exibir o texto descriptografado

### Usar as funções no seu código

```go
package main

import (
    "encrypt-decrypt/functions"
    "fmt"
)

func main() {
    // Chave secreta (32 bytes para AES-256)
    secret := "sua-chave-secreta-de-32-bytes-exato"
    
    // IV (Initialization Vector) - deve ser único para cada criptografia
    iv := []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
    
    // Criptografar
    original := "Texto secreto"
    encrypted, err := functions.Encrypt(original, secret, iv)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Criptografado: %s\n", encrypted)
    
    // Descriptografar
    decrypted, err := functions.Decrypt(encrypted, secret, iv)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Descriptografado: %s\n", decrypted)
}
```

## Detalhes Técnicos

- **Algoritmo**: AES-256 (Advanced Encryption Standard)
- **Modo**: CFB (Cipher Feedback)
- **Encoding**: Base64 para representação segura do resultado
- **Chave**: Requer exatamente 32 bytes para AES-256
- **IV**: 16 bytes para AES no modo CFB

## Segurança

- A chave secreta deve ter exatamente 32 caracteres/bytes
- O IV (Initialization Vector) deve ser único para cada operação de criptografia
- Nunca compartilhe a chave secreta ou o IV
- Para produção, considere gerar IVs aleatórios para cada criptografia

## Compilação

Para compilar o executável:

```bash
go build -o encrypt-decrypt main.go
```

## Licença

Este projeto é para fins educacionais e demonstrativos.

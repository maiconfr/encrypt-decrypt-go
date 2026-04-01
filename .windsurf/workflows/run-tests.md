---
description: Executar testes da aplicação Go
---

---
name: run-tests
description: Workflow para executar testes unitários, benchmarks e verificar cobertura de código
---

## Verificar dependências
1. Verifica se as dependências Go estão instaladas
2. Executa `go mod tidy` para garantir que as dependências estão atualizadas

## Executar testes unitários
1. Roda todos os testes com `go test -v`
2. Exibe resultado detalhado dos testes
3. Verifica se todos os testes passaram

## Executar benchmarks
1. Roda testes de performance com `go test -bench=.`
2. Exibe métricas de performance das operações de criptografia/descriptografia

## Verificar cobertura de código
1. Gera relatório de cobertura com `go test -cover`
2. Cria arquivo HTML de cobertura para visualização detalhada

## Comandos utilizados:
- `go mod tidy` - Atualizar dependências
- `go test -v` - Executar testes com detalhes
- `go test -bench=.` - Executar benchmarks
- `go test -cover` - Verificar cobertura de código
- `go test -coverprofile=coverage.out && go tool cover -html=coverage.out` - Gerar relatório HTML

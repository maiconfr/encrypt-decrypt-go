---
description: Realizar commit no Git
auto_execution_mode: 1
---

---
name: git-commit
description: Workflow para realizar commits no Git com verificação de status e mensagem padronizada
---

## Identificar idioma para o repositório
1. Pergunte qual idioma deseja realizar o comentário do commit

## Verificar status do repositório
1. Verifica o status atual do repositório Git
2. Lista arquivos modificados, novos ou deletados

## Adicionar arquivos ao staging
1. Adiciona todos os arquivos modificados ao staging area
2. Permite seleção específica de arquivos se necessário

## Criar commit
1. Solicita mensagem de commit descritiva
2. Executa o commit com a mensagem fornecida
3. Exibe confirmação do commit realizado

## Push (opcional)
1. Pergunta se deseja enviar as alterações para o repositório remoto
2. Executa push se confirmado

## Comandos utilizados:
- `git status` - Verificar status
- `git add .` - Adicionar todos os arquivos
- `git commit -m "mensagem"` - Criar commit
- `git push` - Enviar para repositório remoto
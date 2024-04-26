Claro, aqui está o README.md revisado para a sua solução do projeto:

---

# Cadastro e Listagem de Produtos - Desafio Técnico Oak Tecnologia

Este é o projeto desenvolvido como solução para o Desafio Técnico para a vaga de estágio em TI na Oak Tecnologia. A aplicação consiste em um sistema de cadastro e listagem de produtos, utilizando a linguagem de programação Go (Golang) com o framework web Gin e o banco de dados SQLite.

## Funcionalidades

- ### Cadastro de produtos com nome, descrição, preço e disponibilidade para venda.
![Captura de tela de 2024-04-26 04-31-00](https://github.com/marcuscarvalhodev/oak-tecnologia-desafio-tecnico/assets/135276762/0501faf1-c2a7-4b9e-bd37-8a4bad830685)
- ### Listagem de todos os produtos cadastrados.
![Captura de tela de 2024-04-26 04-31-16](https://github.com/marcuscarvalhodev/oak-tecnologia-desafio-tecnico/assets/135276762/43a9cbbb-c6a0-443e-96fa-1d4e9c9ef459)

## Executando a Aplicação

1. Clone este repositório:

```
git clone https://github.com/marcuscarvalhodev/oak-tecnologia-desafio.git
```

2. Navegue até o diretório do projeto:

```
cd oak-tecnologia-desafio
```

3. Execute o arquivo principal para iniciar o servidor:

```
go run cli.go
```
4. Ou apenas execute o arquivo executável 'cli' e ele irá rodar todas as dependencias do software.
```
./cli
```


## Estrutura do Projeto

- **`main.go`**: Arquivo principal que contém a inicialização do servidor e a configuração das rotas da API.
- **`database.db`**: Arquivo do banco de dados SQLite.
- **`view/`**: Pasta contendo os arquivos HTML e CSS para a interface do usuário.
- **`assets/`**: Pasta contendo os recursos estáticos, como imagens e estilos CSS.

## Melhorias Futuras

- Implementação de edição e exclusão de produtos cadastrados.
- Validação dos campos do formulário de cadastro na interface web.
- Paginação na listagem de produtos.
- Ordenação dos produtos por nome, preço, etc.
- Testes automatizados para a API e/ou interface web.

--- 

# 📘 API-Go-Registro-De-Usuarios

Uma API REST simples desenvolvida em Go para registrar, listar, atualizar e deletar usuários. Esse projeto tem como objetivo o aprendizado dos fundamentos da linguagem Go e da construção de APIs RESTful com o framework Gorilla Mux.

---

## 🚀 Funcionalidades

- ✅ Listar todos os usuários (`GET /listar`)
- ✅ Listar usuário por ID (`GET /listar/{id}`)
- ✅ Cadastrar novo usuário (`POST /cadastrar`)
- ✅ Atualizar usuário (`PATCH /atualizar/{id}`)
- ✅ Deletar usuário (`DELETE /deletar/{id}`)

---

## 🧠 Estrutura do Projeto

O projeto mantém todos os dados em memória (sem persistência em banco de dados), o que é ideal para fins de estudo e testes rápidos com APIs REST.

### 📄 Estrutura de Dados (JSON)

```json
{
  "nome": "João da Silva",
  "id": 1,
  "email": "joao@email.com",
  "done": false
}
```

---

## 📡 Rotas da API
<div align=center>

| Método | Rota               | Descrição                       |
|--------|--------------------|---------------------------------|
| GET    | `/listar`          | Lista todos os usuários         |
| GET    | `/listar/{id}`     | Lista um usuário pelo ID        |
| POST   | `/cadastrar`       | Cadastra um novo usuário        |
| PATCH  | `/atualizar/{id}`  | Atualiza um usuário existente   |
| DELETE | `/deletar/{id}`    | Remove um usuário pelo ID       |
  
</div>

---

## 🛠️ Como executar localmente

1. Clone o repositório:
   ```bash
   git clone https://github.com/napolisiqueira/API-Go-Registro-De-Usuarios.git
   cd API-Go-Registro-De-Usuarios
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Execute a aplicação:
   ```bash
   go run main.go
   ```

4. A API estará disponível em:  
   [http://localhost:3000](http://localhost:3000)

---

## ⚠️ Observações

- Os dados são armazenados apenas em memória, portanto são perdidos ao reiniciar a aplicação.
- Ainda há algumas melhorias a serem feitas, como:
  - Validação de entrada mais robusta
  - Suporte a banco de dados
  - Separação de camadas (handlers, services, models, etc.)

---

## 🧑‍💻 Autor

Desenvolvido por [Felipe Napoli Siqueira](https://github.com/napolisiqueira)

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.


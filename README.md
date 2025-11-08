<h1 align="center">📘 API-Go-Registro-De-Usuarios</h1>

<p align="center">
  <b>API REST simples em Go para registrar, listar, atualizar e deletar usuários.<br>
  Desenvolvido para aprendizado dos fundamentos da linguagem Go e de APIs RESTful com Gorilla Mux.</b>
</p>

<hr/>

<h2>🚀 Funcionalidades</h2>
<ul>
  <li>Listar todos os usuários (<code>GET /listar</code>)</li>
  <li>Listar usuário por ID (<code>GET /listar/{id}</code>)</li>
  <li>Cadastrar novo usuário (<code>POST /cadastrar</code>)</li>
  <li>Atualizar usuário (<code>PATCH /atualizar/{id}</code>)</li>
  <li>Deletar usuário (<code>DELETE /deletar/{id}</code>)</li>
</ul>

<hr/>

<h2>🧠 Estrutura e Dados</h2>
<ul>
  <li>Todos os dados são mantidos em memória (RAM), ideal para estudo e testes rápidos.</li>
  <li>Exemplo de estrutura JSON:
    <pre><code>{
  "nome": "João da Silva",
  "id": 1,
  "email": "joao@email.com",
  "done": false
}</code></pre>
  </li>
</ul>

<hr/>

<h2>📡 Rotas da API</h2>
<table>
  <thead>
    <tr>
      <th>Método</th>
      <th>Rota</th>
      <th>Descrição</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>GET</td><td>/listar</td><td>Lista todos os usuários</td></tr>
    <tr><td>GET</td><td>/listar/{id}</td><td>Lista um usuário pelo ID</td></tr>
    <tr><td>POST</td><td>/cadastrar</td><td>Cadastra um novo usuário</td></tr>
    <tr><td>PATCH</td><td>/atualizar/{id}</td><td>Atualiza um usuário existente</td></tr>
    <tr><td>DELETE</td><td>/deletar/{id}</td><td>Remove um usuário pelo ID</td></tr>
  </tbody>
</table>

<hr/>

<h2>🛠️ Como executar localmente</h2>
<ol>
  <li>Clone o repositório:
    <pre><code>git clone https://github.com/napolisiqueira/API-Go-Registro-De-Usuarios.git
cd API-Go-Registro-De-Usuarios
    </code></pre>
  </li>
  <li>Instale as dependências:
    <pre><code>go mod tidy</code></pre>
  </li>
  <li>Execute a API:
    <pre><code>go run main.go</code></pre>
  </li>
  <li>Acesse em: 
    <a href="http://localhost:3000" target="_blank">http://localhost:3000</a>
  </li>
</ol>

<hr/>

<h2>👨‍💻 Autor</h2>
<ul>
  <li><b>Felipe Napoli Siqueira</b></li>
  <li><a href="https://github.com/napolisiqueira" target="_blank">@napolisiqueira</a></li>
</ul>

<hr/>

<h2>📝 Licença</h2>
<p>Este projeto está sob licença MIT. Veja <a href="LICENSE">LICENSE</a> para detalhes.</p>

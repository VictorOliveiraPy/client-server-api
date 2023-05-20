# client-server-api

**Olá dev, tudo bem?**

*Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.*

## Requisitos

Os requisitos para cumprir este desafio são:

- O `client.go` deverá realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.

- O `server.go` deverá consumir a API contendo o câmbio de Dólar e Real no endereço: `https://economia.awesomeapi.com.br/json/last/USD-BRL` e em seguida deverá retornar no formato JSON o resultado para o cliente.

- Usando o package `context`, o `server.go` deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

- O `client.go` precisará receber do `server.go` apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package `context`, o `client.go` terá um timeout máximo de 300ms para receber o resultado do `server.go`.

- O `client.go` terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

- O endpoint necessário gerado pelo `server.go` para este desafio será: `/cotacao` e a porta a ser utilizada pelo servidor HTTP será a `8080`.





A pasta `internal` contém o código fonte interno do projeto. Dentro dela, temos as pastas `client` e `server`, onde colocamos os arquivos `client.go` e `server.go`, respectivamente.

O arquivo `README.md` contém as instruções e informações sobre o projeto.


## Execução


# Para executar o servidor:
``go run cmd/server/main.go``


Certifique-se de estar na raiz do projeto ao executar os comandos acima.

``Lembre-se de atualizar as dependências do projeto usando o comando `go mod tidy` antes de executá-lo.``

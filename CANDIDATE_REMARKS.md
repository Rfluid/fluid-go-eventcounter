# How to run

Eu adicionei uma regra no Makefile para rodar o código do `cmd/consumer/main.go`, incluindo inciar o RabbitMQ, publicar mensagens, etc. Para rodar sem estresse, basta `make dnwmt`.

Adicionei outra regra só para rodar o código em si, `make run`.

# Code patterns

Eu não entendi se o padrão do projeto é inglês ou português. Então eu fiz um pouco de cada. Se for necessário, eu posso traduzir tudo para inglês.

Inferi onde colocar cada coisa aqui. Adotei `pkg/` como uma pasta para colocar structs (e consequentemente interfaces) usadas em `cmd` e funções auxiliares para usar dentro dessas structs.

Dei uma pequena refatorada em `cmd/consumer/count.go` para usar uma função mais geral de escrever JSONs. Eu poderia ter adotado nesse package o padrão de projeto que usei para a conexão com o RabbitMQ e flags, mas achei que não era necessário e pensei que talvez fosse melhor manter o vosso código para que vocês tenham mais clareza nos testes por aí.

# New structs

Eu elaborei uma estrutura que se chama Mutex Swapper, que usa um Mutex global para selecionar e bloquear um mutex específico, que normalmente é relacionado a um id de usuário ou algo assim. Essa estrutura aumenta a concorrência do código, porque não estamos apenas tratando os EventTypes em paralelo, mas também os usuários dentro de cada EventType.

# Packages

- `exchange_listener`
  Cuida dos channels que vão receber as mensagens do RabbitMQ e passa para o handler do respectivo `EventType`. Faz todo o ciclo de vida da conexão com o RabbitMQ e dos channels (que precisam ser fechados já que são buffered).

- `out`
  Escreve outputs.

- `rabbitmq`
  Faz a conexão com o RabbitMQ.

- `config`
  Possui um único pacote que lê as flags, mas pode ser expandindo para ler um arquivo de configuração (como `.env`) ou algo assim.

---

σΔγ

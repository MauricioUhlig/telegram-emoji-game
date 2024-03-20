# telegram-emoji-game
Este é apenas um jogo que utiliza o bot de telegram que pode ser adicionado a um grupo
Ao mandar o comando `/new` ele devolve uma lista com `models.QTD_EMOJIS_ROOM` emojis

A ideia do jogo é que os jogadores formulem histórias que envolvam os emojis retornados aleatóriamente e com isso se divertirem

## Setup
Antes de iniciar é necessário criar um bot com o @botFather do telegram e pegar a API Key
Esta API Key deverá ser colocada no arquivo `configs.yaml` na seção `Telegram.key`
Depois disto basta dar o comando `go run main.go`, adicionar seu bot a algum grupo e inicar a interação com ele

## Personalização
É possível alterar a quantidade de emojis retornados a cada interação, basta alterar a constante `QTD_EMOJIS_ROOM` no arquivo `models/emojis.go` para o valor que você deseja

# Referências 
- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
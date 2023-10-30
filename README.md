# API de transferência de dinheiro como a do PicPay

Nesse desafio você tem a responsabilidade de criar um Restful API que provê formas de transferir dinheiro de uma pessoa para outra. Cada usuário no sistema tem um saldo que é atualizado em cada transferência. 

## Requisitos

- Crie um endpoint que recebe dois IDs de usuários, e um valor monetário representando a transferência entre eles. 
- Crie um endpoint que recebe um ID de usuário e retorna o saldo dele. 
- Valide se o usuário de origem tem saldo suficiente antes da transferência. 
- É preciso pensar na possibilidade de concorrência de transferência onde duas pessoas transferem dinheiro ao mesmo tempo para uma terceira. 
- Se uma transferência falhas, o saldo do usuário de origem deve ser restaurado. 
- Não é necessário endpoints para criar usuário, popule o banco de forma com que os dois usuário existam e que transferências possam ser feitas entre eles. 

## Desafio baseado no seguinte site da Devgym
Desafio de nível intermediário. 

[Devgym](https://app.devgym.com.br)

## Como rodar a aplicação

1. Clone do projeto
2. Rodar o projeto ```go run .```
3. Testar a rota home ``` localhost:8080\api\v1\home ```
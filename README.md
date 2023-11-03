# API de transferência de dinheiro como a do PicPay

Nesse desafio você tem a responsabilidade de criar um Restful API que provê formas de transferir dinheiro de uma pessoa para outra. Cada usuário no sistema tem um saldo que é atualizado em cada transferência. 

## O que precisa ter
- Endpoint para consultar o saldo, repassando o ID do usuário
- Endpoint para realizar uma transferência repassando os ID dos usuário e depois o valor monetário
- Validação de saldo, caso saldo insuficiente apresentar uma mensagem para o usuário
- Se uma transferencia falhar deverá retornar os saldos dos usúarios. 

## Tarefas
[x] Endpoint para consultar o saldo repassando o ID do usuário <br>
[x] Endpoint para realizar a transferência <br>
[x] Rollback quando o usuário não tiver saldo <br>
[] Pensar em uma forma de concorrência, caso dois usuário façam a transferência ao mesmo tempo para o mesmo usuário <br>
[] Criar uma rota para criar usuários <br> 
[] Adicionar testes para cada rota <br>
[x] Função para criar as tabelas necessárias <br>
[] Adicionar o docker<br>
[] Adicionar o swagger <br>
[] Endpoint para consultar todos os usuário <br>
[] Endpoint para cadastro de conta <br>
[] Endpoint para extrato <br>

## Desafio baseado no site da Devgym
 
[Devgym](https://app.devgym.com.br)
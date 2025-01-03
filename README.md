
# Order Processor

Um sistema distribuído para processamento de pedidos utilizando RabbitMQ, e implementando microsserviços com NestJS e Golang.

## Arquitetura

### API (NestJS)
- Recebe requisições para criar o pedido
- Valida as informações do pedido
- Publica o pedido no RabbitMQ para processá-lo

### Inventory Service (Golang)
- Verifica disponibilidade do produto no estoque
- Atualiza o estoque no banco de dados

### Payment Service (Golang)
- Processa o pagamento do pedido

### Notification Service (Golang)
- Envia notificação para o cliente conforme o resultado do processamento
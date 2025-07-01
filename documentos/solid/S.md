# Single Responsibility Principle (SRP)

## Definição Clássica
O SRP afirma que uma classe ou módulo deve ter **apenas uma razão para mudar**, ou seja, ela deve ser responsável por **uma única tarefa ou funcionalidade**.

## Importância e Aplicação

- **Coesão**: O SRP aumenta a **coerência interna** do código, pois classes e módulos que têm uma única responsabilidade tendem a ser mais coesos, o que melhora a **manutenibilidade** e a **compreensão** do código.
  
- **Flexibilidade**: Quando uma classe ou módulo tem uma única responsabilidade, ela pode ser **modificada** sem afetar outras partes do sistema, permitindo que a aplicação evolua com menos riscos de causar **efeitos colaterais**.
  
- **Testabilidade**: A separação de responsabilidades facilita **testes unitários**, pois é mais fácil escrever testes focados em uma única responsabilidade. Isso também melhora o **mocking** e a **injeção de dependências**.

## Erros Comuns de SRP

- **Micro Gerenciamento de Arquivos**: Um erro comum é **dividir demais**. Embora a separação de responsabilidades seja importante, criar uma classe ou módulo para cada pequena função pode resultar em **over-engineering**, tornando o sistema difícil de navegar. Um desenvolvedor sênior sabe balancear essa separação.

- **Excesso de Dependência**: Se uma classe ou função se torna excessivamente dependente de outras funções/módulos ao ponto de gerar **interdependência** (com muitos métodos ou dados externos), pode indicar uma violação do SRP.

## Exemplo Prático

### Violação do SRP (Exemplo de Classe com Múltiplas Responsabilidades)

Imagine uma classe responsável por **processar um pedido de pagamento**, mas ela também está fazendo **checagem de inventário** e **cálculo de total**.

```go
type PedidoProcessamento struct{}

func (pp PedidoProcessamento) ProcessarPagamento(pedido Pedido) {
    pp.checarInventario(pedido)
    pp.calcularTotal(pedido)
    pp.processarPagamento(pedido)
}

func (pp PedidoProcessamento) checarInventario(pedido Pedido) {
    fmt.Println("Verificando inventário...")
}

func (pp PedidoProcessamento) calcularTotal(pedido Pedido) {
    fmt.Println("Calculando total do pedido...")
}

func (pp PedidoProcessamento) processarPagamento(pedido Pedido) {
    fmt.Println("Processando pagamento...")
}
```

## Como Resolver Usando o SRP?

### Quebrar em Responsabilidades Menores
Cada responsabilidade deve ser movida para uma **classe separada**. Isso garante que cada classe tenha apenas uma razão para mudar, mantendo o código mais coeso e fácil de manter.

### Injeção de Dependência
Usar a **injeção de dependência (DI)** para passar as dependências para a classe `PedidoProcessamento`. Isso ajuda a manter a classe `PedidoProcessamento` desacoplada das outras responsabilidades, facilitando a manutenção e os testes.

### Exemplo de Código

```go
type ChecadorDeInventario struct{}

func (ci ChecadorDeInventario) Checar(pedido Pedido) {
    fmt.Println("Verificando inventário...")
}

type CalculadoraDeTotal struct{}

func (ct CalculadoraDeTotal) Calcular(pedido Pedido) {
    fmt.Println("Calculando total do pedido...")
}

type ProcessadorDePagamento struct{}

func (pp ProcessadorDePagamento) Processar(pedido Pedido) {
    fmt.Println("Processando pagamento...")
}

type PedidoProcessamento struct {
    checadorDeInventario    ChecadorDeInventario
    calculadoraDeTotal      CalculadoraDeTotal
    processadorDePagamento  ProcessadorDePagamento
}

func (pp PedidoProcessamento) Processar(pedido Pedido) {
    pp.checadorDeInventario.Checar(pedido)
    pp.calculadoraDeTotal.Calcular(pedido)
    pp.processadorDePagamento.Processar(pedido)
}
```

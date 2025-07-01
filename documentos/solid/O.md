# Open/Closed Principle (OCP)

## Definição Clássica
O **Open/Closed Principle (Princípio Aberto/Fechado)** afirma que **uma classe deve estar aberta para extensão, mas fechada para modificação**. Isso significa que é possível adicionar novos comportamentos a um sistema sem alterar o código já existente.

> **Aberto para Extensão:** É possível adicionar novas funcionalidades.
> 
> **Fechado para Modificação:** O código existente não deve ser alterado para suportar novas funcionalidades.

---

## Importância e Aplicação

- **Facilita a Escalabilidade:** O OCP permite que o sistema cresça sem impactar o que já foi implementado e testado.
- **Reduz Riscos:** Modificar código existente pode quebrar funcionalidades já validadas. O OCP protege o sistema disso.
- **Melhora a Manutenção:** Extensões podem ser feitas com novas classes ou implementações, mantendo o código organizado.
- **Estimula o Uso de Polimorfismo:** Estratégias como interfaces e herança são fundamentais para aplicar corretamente o OCP.

---

## Ligação com o Strategy Pattern

O OCP é frequentemente aplicado em conjunto com o **Strategy Pattern**, que permite encapsular algoritmos intercambiáveis em diferentes classes.

Com Strategy, você pode passar diferentes "estratégias" para uma classe principal, sem precisar alterá-la. Esse padrão viabiliza a extensão do comportamento sem modificar a lógica existente.

---

## Problema: Solução Sem OCP

### Exemplo
Uma classe de pagamento onde é necessário alterar o código sempre que se adiciona um novo método de pagamento.

```go
type PaymentService struct{}

func (p PaymentService) Pay(method string) {
    if method == "credit" {
        fmt.Println("Processing credit card payment")
    } else if method == "paypal" {
        fmt.Println("Processing PayPal payment")
    }
}
```

## Problema: Toda vez que um novo método de pagamento for adicionado, será necessário alterar a lógica existente, o que viola o OCP.

# Solução Aplicando OCP
### Exemplo com Interface e Strategy Pattern

```go
type PaymentProcessor interface {
    Process(amount float64)
}

type CreditCard struct{}

func (c CreditCard) Process(amount float64) {
    fmt.Println("Processing credit card payment")
}

type PayPal struct{}

func (p PayPal) Process(amount float64) {
    fmt.Println("Processing PayPal payment")
}

type PaymentService struct {
    processor PaymentProcessor
}

func (ps PaymentService) Pay(amount float64) {
    ps.processor.Process(amount)
}
```

# Benefício:
### Para adicionar novos métodos de pagamento, basta criar novas structs que implementam a interface PaymentProcessor, sem precisar alterar a classe PaymentService.


## 3. LSP – Liskov Substitution Principle

**Definição:**  
Se voce começar a herdar de uma clase qualquer, voce tem que ser muito cuidadoso, quando voce muda a pre e pos condição, que pode fazer o sistema ter comportamentos erroneos nas subclasses.

Classes filhas ou classes derivadas **nunca devem infringir comportamentos e definições de tipo** da classe base ou da interface que estão implementando.

> _Se parece com um pato, tem som de pato, mas precisa de bateria para funcionar, provavelmente tem algo errado com sua abstração._

---

### Conceitos Importantes

#### Pre-condição
- 🔒 **Ligada às regras estabelecidas na classe base.**
- Uma subclasse **não pode exigir mais** do que a classe base exigia.
- Se sobrescrever uma pré-condição numa herança, **vai quebrar o código cliente.**

#### Pós-condição
- 🎯 **Ligada diretamente ao retorno dos objetos.**
- Uma subclasse **não pode reduzir as garantias** fornecidas pela classe base após a execução do método.

#### Invariância
- ⚙️ **Regra que não pode ter variação.**
- A subclasse **não pode alterar condições internas** que a classe base mantinha constante.

---

## Problema: Substituição quebra o sistema

```go
type Bird interface {
    Fly()
}

type Duck struct{}

func (d Duck) Fly() {
    fmt.Println("Duck flying!")
}

type Ostrich struct{}

func (o Ostrich) Fly() {
    panic("Ostrich can't fly!") // Problema
}
```

## 👉 Ostrich quebra o programa ao ser usada no lugar de Duck.

# Solução: Ajustar a abstração

```go
type Bird interface {
    Move()
}

type Duck struct{}

func (d Duck) Move() {
    fmt.Println("Duck flying!")
}

type Ostrich struct{}

func (o Ostrich) Move() {
    fmt.Println("Ostrich running!")
}
````
 Agora:

Qualquer Bird pode ser substituído por outro sem quebrar o programa.
# Polimorfismo em Go

## O que é Polimorfismo?
Polimorfismo permite que você use o mesmo nome para métodos ou funções, mas com comportamentos diferentes dependendo do tipo do objeto.

Polimorfismo é a capacidade de tratar diferentes tipos de dados de forma uniforme através de uma **interface comum**.  
Em Go, **polimorfismo é alcançado exclusivamente via interfaces**, não através de herança (Go não tem herança tradicional).

---

## Polimorfismo em Go: O Conceito-Chave

Em Go:
- **Não há herança.**
- O polimorfismo é **estritamente baseado em comportamento**, via **interfaces**.
- A implementação é **implícita** (não é necessário declarar que um tipo implementa uma interface).

---

## Exemplo Básico

```go
package main

import "fmt"

// Interface
type Forma interface {
    Area() float64
}

// Tipo 1: Retângulo
type Retangulo struct {
    Largura, Altura float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

// Tipo 2: Círculo
type Circulo struct {
    Raio float64
}

func (c Circulo) Area() float64 {
    return 3.14 * c.Raio * c.Raio
}

func MostrarArea(f Forma) {
    fmt.Printf("Área: %.2f\n", f.Area())
}

func main() {
    r := Retangulo{Largura: 10, Altura: 5}
    c := Circulo{Raio: 7}

    MostrarArea(r)
    MostrarArea(c)
}
```

# Pontos importantes sobre Polimorfismo em Go

- **`Forma` é a interface.**
- **`Retangulo` e `Circulo` são tipos concretos que implementam a interface.**
- **A função `MostrarArea` aceita qualquer tipo que satisfaça a interface `Forma`.**


## Benefícios do Polimorfismo em Go

- **Desacoplamento:** O código depende de comportamentos, não de implementações concretas.
- **Flexibilidade:** É fácil adicionar novos tipos sem alterar código existente.
- **Composição > Herança:** Go favorece composição, permitindo criar estruturas flexíveis sem hierarquias profundas.

## Diferenças para Linguagens OO Clássicas

| Característica    | Go          | Java/C#     |
|-------------------|-------------|-------------|
| Herança           | ❌          | ✅           |
| Interfaces        | ✅          | ✅           |
| Implementação     | Implícita   | Explícita    |
| Sobrecarga        | ❌          | ✅           |
| Sobrescrita       | ✅          | ✅           |

- Go **não suporta sobrecarga de métodos** (mesmo nome, parâmetros diferentes).
- Go **não possui herança de estruturas ou métodos.**


## Polimorfismo + Type Assertion

Go permite verificar o tipo concreto através de **type assertions** ou **type switches**.

```go
func IdentificarForma(f Forma) {
    switch v := f.(type) {
    case Retangulo:
        fmt.Println("É um retângulo:", v)
    case Circulo:
        fmt.Println("É um círculo:", v)
    default:
        fmt.Println("Forma desconhecida")
    }
}
```

# Pontos Avançados sobre Polimorfismo em Go

## Nota Importante

O **uso excessivo de `type assertion` pode indicar quebra do polimorfismo**.  
Sempre que possível, **trabalhe no nível da abstração**.

---

## Quando um Desenvolvedor Sênior Precisa Saber

- **Interfaces pequenas são melhores.**  
  (Princípio da Interface Segregation)
  
- **Composição é preferida.**  
  Go incentiva o uso de composição no lugar de herança.

- **Evitar criar interfaces prematuramente.**  
  No Go, interfaces devem nascer pelo **consumo**, não pela definição.

- **Domínio das interfaces padrão.**  
  Saber trabalhar com `io.Reader`, `io.Writer`, `http.Handler` e outras interfaces do pacote padrão que são exemplos poderosos e idiomáticos de polimorfismo na linguagem.

- **Conhecer os riscos de `nil` em interfaces.**  
  No Go:
  ```go
  var f Forma = nil
  if f == nil {
      fmt.Println("Forma é nil") // Isso pode não funcionar como esperado
  }
```

## Interfaces com Ponteiro Nulo: Armadilha Comum

Quando uma interface contém um **ponteiro nulo para um tipo concreto**, a interface **não é nil**, apenas o valor dentro dela é.

Esse é um dos erros mais comuns em Go para quem está começando a trabalhar com interfaces.

---

## Conclusão

Um desenvolvedor **sênior em Go** deve ser capaz de:

✅ Explicar polimorfismo via interfaces com clareza.  
✅ Usar polimorfismo para criar código desacoplado e extensível.  
✅ Reconhecer padrões idiomáticos de polimorfismo no ecossistema Go (ex.: `io.Reader`, `http.Handler`).  
✅ Saber **quando não usar interfaces** (criar interfaces sem necessidade é um erro comum).  
✅ Compreender profundamente as armadilhas envolvendo:
- Interfaces nulas.
- Uso inadequado de **type assertion**.
- O impacto de **composição vs herança**.

---

## Resumo

O polimorfismo no Go é **poderoso, simples e eficaz quando utilizado de forma idiomática.**  
Saber usar bem as interfaces é um **diferencial essencial no perfil de um sênior Go.**



# ISP – Interface Segregation Principle

## Problema
O problema central do ISP é a criação de **interfaces muito grandes**, que forçam as classes a implementarem métodos desnecessários ou irrelevantes para o seu funcionamento.

### Exemplo de Código sem ISP

```go
type Worker interface {
    Work()
    Eat()
}

type Human struct{}

func (h Human) Work() { fmt.Println("Working...") }
func (h Human) Eat()  { fmt.Println("Eating...") }

type Robot struct{}

func (r Robot) Work() { fmt.Println("Working...") }
func (r Robot) Eat()  { /* ??? Robots don't eat */ }
```

## 👉 Robot é forçado a implementar Eat, que não faz sentido.

```go
type Workable interface {
    Work()
}

type Eatable interface {
    Eat()
}

type Human struct{}

func (h Human) Work() { fmt.Println("Working...") }
func (h Human) Eat()  { fmt.Println("Eating...") }

type Robot struct{}

func (r Robot) Work() { fmt.Println("Working...") }
```

## ✅ Agora:

## Cada struct implementa apenas as interfaces que fazem sentido.
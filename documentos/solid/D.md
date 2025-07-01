# DIP – Dependency Inversion Principle

## Problema
O problema central é que a classe depende diretamente de uma implementação concreta. Este princípio se alinha bem com o padrão de projeto *Adapter*, ajudando a melhorar a flexibilidade, manutenção do código e, principalmente, a testabilidade.

## Objetivo
O objetivo do DIP é garantir que o código seja sempre escrito para interfaces, abstrações, e nunca para implementações concretas. Isso melhora a flexibilidade do sistema, permitindo que a implementação da interface seja alterada sem afetar as classes que dependem dela.

Este princípio é um dos pilares que norteiam arquiteturas como **Arquitetura Hexagonal** e **Arquitetura em Camadas**.

## Exemplo de Código sem DIP

```go
type MySQLDatabase struct {}

func (db MySQLDatabase) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}

type UserService struct {
    database MySQLDatabase
}

func (us UserService) SaveUser(data string) {
    us.database.Save(data)
}
```

## 👉 UserService está acoplado ao MySQLDatabase.


```go
type Database interface {
    Save(data string)
}

type MySQLDatabase struct{}

func (db MySQLDatabase) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}

type UserService struct {
    database Database
}

func NewUserService(db Database) *UserService {
    return &UserService{database: db}
}

func (us UserService) SaveUser(data string) {
    us.database.Save(data)
}
```

## Agora:

## UserService depende de uma abstração (interface).

### Você pode trocar MySQLDatabase por outra implementação sem mudar UserService.
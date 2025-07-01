 DIP – Dependency Inversion Principle

 Problema: Classe depende diretamente de implementação concreta. Se ajusta bem com padrao projeto adpter
 Melhorar a flexibilidade , manutenção do codigo e principalmente a testabilidade
 E o principio que norteia arquitetura hexagonal, arquitetura em camadas

 Nosso codigo e sempre bom programar para interfaces, classes abstratam, nunca para implementações concretas

type MySQLDatabase struct{}

func (db MySQLDatabase) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}

type UserService struct {
    database MySQLDatabase
}

func (us UserService) SaveUser(data string) {
    us.database.Save(data)
}


👉 UserService está acoplado ao MySQLDatabase.


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

 Agora:

UserService depende de uma abstração (interface).

Você pode trocar MySQLDatabase por outra implementação sem mudar UserService.
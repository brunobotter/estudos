3. LSP – Liskov Substitution Principle

classes filhas ou classes derivadas nunca devem infrigir comportamentos e definições de tipo de classes base ou da interface a qual essas classes implementam
Se parece com um pato, tem som de pato, mas precisa de bateria para funcionar, provavelmente tem algo errado com sua abstração

Pre-condição:( ligada as regras estabelecidas na classe base)
uma subclasse nao pode exigir mais do que a classe base exigia. Se sobrescrever uma pre condição numa herança, vai quebrar o codigo cliente

Pós-condição: (ligada diretamente ao retorno dos objetos)
Uma subclasse nao pode reduzir as garantias fornecidas pela classe base apos a execução do metodo

Invariancia: (Regra que nao pode ter variação)
A subclasse nao pode alterar condições internas que a classe base mantinha constante

Problema: Substituir uma implementação quebra o sistema.

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

👉 Ostrich quebra o programa ao ser usada no lugar de Duck.

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

 Agora:

Qualquer Bird pode ser substituído por outro sem quebrar o programa.
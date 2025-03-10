# linked list

Uma estrutura que guarda nos ligados:
Contem um sequencia de nos
Um no tem uma informação valor e guarda o endereço do proximo no
O primeiro no e o HEAD
O ultimo no aponta para null

type Node struct{
    Data string //pode ser qualquer valor
    Next *Node
}

Ele nao e literalmente uma lista e sim um objeto que segue a ideia de uma lista
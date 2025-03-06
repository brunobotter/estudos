# Estruturas de dados

# Dados lineares

Estruturas de dados lineares são aquelas que armazenam dados em uma ordem sequencial, onde cada elemento tem um predecessor e sucessor exclusivos. 

1. linked list 
2. Array / Slice
3. Queue
4. Stack

# Dados não lineares

Estruturas de dados não lineares são aquelas que armazenam dados em uma ordem hierárquica ou em rede, onde cada elemento pode ter vários predecessores e sucessores. 

1. Tree
2. Graph


# Complexidade do tempo

E a forma de ter uma noção de quanto tempo seu algoritimo levaria para terminar sua execução.

Exemplo de um input recebido nos ajuda a ter uma ideia de como nosso algoritimo vai performar

// mais facil porem nao performatico, pois se for um numero muito grande tera que rodar o for o numero de vezes
func naturalNumber(n int) (sum int){
    for i := 1; i <= n; i++{
        sum += i
    }
    return
}

// modo mais performatico com mesmo resultado
func naturalNumber(n int) int{
    return n * (n + 1) / 2
}

# Complexidade de espaço

E a medida do espaço em memoria necessario para executar um algoritimo, e frequentimento expresso como uma função do tamanho de entrada

Exemplo
O(1) significa que o espaço usado e constante, nao importa o tamanho da entrada.

O(n) significa que o tamanho do espaço aumenta linearmente com tamanho da entrada.

Imagine uma lista de numeros, se voce copiar todos esses numeros em uma nova lsita, estara usando espaçõ adicional igual o tamanho da lista original, isso teria uma complexidade de espaçõ de O(n).

# Notações assintoticas

Sao formas matematicas de descrever o tempo de execução de um algoritimo
E uma tecnica utilizada para avaliar o desempenho de algoritimos a medida que o tamanho da entrada cresce para o infinito

# Notação BIG O

Forma de definir o tempo maximo que o algoritimo pode executar
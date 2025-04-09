# 3 Pilares da Escalabilidade

## 1. Caching
Remove sobrecargas de sistemas e estruturas que não têm condições de suportar uma grande carga, como o banco de dados (DB).

## 2. Assíncrono (Async)
Postergar operações para processá-las mais tarde, utilizando mensageria e filas.

## 3. Load Balancing
Distribuição de carga entre várias máquinas.

---

# Microserviços

Microserviços envolvem múltiplas máquinas executando diversos processos e trocando mensagens através de uma rede, que é instável por natureza.

> **Importante:** Microserviços não são uma decisão técnica. Eles existem para escalar **equipes**, e não apenas para escalar **aplicações**.

---

## As 8 Mentiras dos Sistemas Distribuídos

1. A rede é confiável  
2. A latência é zero  
3. A largura de banda é infinita  
4. A rede é segura  
5. A topologia nunca muda  
6. Só existe um administrador  
7. O custo de transporte é zero  
8. A rede é homogênea  

---

# Tipos de Carga

## CPU-Bound
Processo com alto uso de CPU e baixo uso de I/O.

## IO-Bound
Processo com baixo uso de CPU e alto uso de I/O.

---

# Banco de Dados e Performance

- Boa parte do tempo de resposta de uma transação está na comunicação com o banco de dados.  
- Quanto menor o tempo de resposta da transação, maior o **throughput**.  
- Criar conexões com o banco é um recurso caro.  
  → Toda aplicação que se conecta a um DB deve utilizar um **pool de conexões**, com um tamanho de pool apropriado.

---

## Performance de Queries

- Queries entre **0 e 10ms**: excelente  
- Queries entre **10 e 100ms**: podem ser otimizadas  
- Queries acima de **100ms**: precisam ser otimizadas

> Utilizar **stored procedures** pode ajudar na otimização das queries.

---

# Boas Práticas

- Resolver latência e sobrecarga do banco de dados com **caching**.  
- Tudo que puder ser processado **assincronamente** deve ser, para melhorar o **throughput**.

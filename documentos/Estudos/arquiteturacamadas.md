# Arquitetura em Camadas

Uma funcionalidade que estou implementando nesse ponto do meu código conversa com várias outras partes do sistema: banco de dados, serviços externos, bibliotecas, etc.

## 💡 Vantagem

Por exemplo, quero implementar um serviço de e-mail, mas ainda não defini qual biblioteca de envio usar. Posso criar uma interface (casca) que representa esse serviço. Quando decidir qual lib usar, apenas implemento essa interface com a biblioteca escolhida.

## 🎯 Objetivo das Arquiteturas

Todas as arquiteturas modernas buscam o mesmo objetivo: proteger o domínio da aplicação (regra de negócio), deixando-o o mais estável e isolado possível.

### Requisitos de uma boa arquitetura:

1. Independente de framework → mudanças no framework não afetam o domínio.
2. Testável.
3. Independente de UI.
4. Independente de banco de dados.
5. Independente de agentes externos.

> **Regra de dependência**:  
> Camadas externas podem depender das internas, mas nunca o contrário.

---

# 🧭 Diferença entre Arquitetura Hexagonal e Clean Architecture

### Hexagonal Architecture (Ports and Adapters)

- 🧩 O mundo externo conversa com a aplicação através de **portas** (interfaces) e **adaptadores**.
- 🎯 Foco: **isolar** o core da aplicação das entradas e saídas.
- 🗣️ Metáfora:  
  > *"Minha aplicação tem portas para o mundo. Quem quiser falar comigo precisa passar por elas."*

### Clean Architecture

- 🎯 Foco: **centralizar** a regra de negócio, tornando-a independente de frameworks, banco e UI.
- 🔁 Organizada em círculos concêntricos com dependências sempre apontando para dentro.
- 🏰 Metáfora:  
  > *"Minha aplicação é um castelo com muralhas. As regras mais importantes estão no centro."*

---

## 🏗 Arquitetura em Camadas

Separação de responsabilidades em camadas hierárquicas, onde cada camada interage apenas com as camadas adjacentes.

### 📌 Exemplo prático:

Estou implementando uma funcionalidade que integra banco de dados, API de pagamento, libs internas e externas. A lógica de orquestração conversa com várias partes, cada uma encapsulada em sua camada.

### 💡 Vantagem:

Se eu quiser trocar a biblioteca de envio de e-mail (gomail por mailgun), apenas substituo a implementação da interface. A regra de negócio não precisa mudar.

---

## 🔄 Princípios fundamentais

| Princípio                       | Descrição |
|-------------------------------|-----------|
| **Independência de framework** | Trocar `gin` por `echo` não afeta o core. |
| **Testabilidade**              | Separação clara facilita mocks e testes unitários. |
| **Independência de UI**        | Regras funcionam com REST, gRPC, CLI ou GraphQL. |
| **Independência de DB**        | Trocar Postgres por Mongo sem quebrar o domínio. |
| **Independência de agentes externos** | O domínio não sabe nada de Stripe, AWS etc. |

---

## 🔍 Diferenças práticas entre Clean e Hexagonal

| Aspecto              | Clean Architecture       | Hexagonal Architecture         |
|----------------------|--------------------------|--------------------------------|
| Organização          | Círculos concêntricos    | Núcleo com portas e adaptadores |
| Comunicação externa  | Camadas (ex: controller → usecase) | Portas (interfaces)      |
| Foco                 | Centralizar o domínio    | Isolar o core com portas       |
| Visual mental        | Cebola com camadas       | Núcleo com plug-ins em volta   |

---

## 🧠 Aplicando no dia a dia (Go)

- Crie interfaces no domínio (ex: `EmailSender`, `OrderRepository`, `PaymentGateway`)
- Use injeção de dependência (via construtores)
- Mantenha o domínio limpo e testável
- Infraestrutura (PostgreSQL, Redis, Stripe) fica fora e implementa apenas interfaces

---

## ✅ Pergunta clássica de entrevista

**"Por que seguir essa arquitetura?"**

> Porque sistemas mudam. Frameworks e libs mudam. O que não pode mudar é a regra de negócio. Se protegermos ela, podemos evoluir o resto com segurança.

---

# 🔷 Arquitetura Hexagonal (Ports and Adapters)

### 📌 O que é?

Criada por **Alistair Cockburn**, visa isolar o núcleo da aplicação das formas como ela é usada (UI, banco, APIs etc.).

> *"Queremos que a lógica de negócio não saiba nada sobre bancos, HTTP, filas, arquivos ou bibliotecas externas."*

### 🎯 Objetivos

- Domínio **puro**, sem dependências técnicas
- Facilidade para **trocar adaptadores**
- Testabilidade e modularidade

---

### 🔄 Conceitos principais

#### 🔸 Core da aplicação

- Contém **regras de negócio**
- Não conhece frameworks, DBs ou APIs
- Composto por entidades, serviços e **interfaces (ports)**

#### 🔸 Ports (Portas)

Interfaces de entrada e saída:

| Tipo          | Direção     | Exemplo                            | Quem chama quem                 |
|---------------|-------------|------------------------------------|----------------------------------|
| Driving Port  | Entrada     | `PlaceOrder(input OrderRequest)`   | UI → Domínio                    |
| Driven Port   | Saída       | `EmailSender.Send(to, subject...)` | Domínio → Infraestrutura       |

#### 🔸 Adapters (Adaptadores)

Implementações concretas das portas:

- HTTP, banco, Kafka, CLI, gRPC, etc.
- Ficam **fora** do core

---

## 📦 Vantagens da Arquitetura Hexagonal

| Benefício            | Descrição                                      |
|----------------------|------------------------------------------------|
| 🔌 Isolamento         | Domínio desacoplado da infraestrutura          |
| 🔁 Facilidade de troca| Troca de libs sem mexer no core                |
| 🧪 Testabilidade      | Testa o domínio com mocks                      |
| 📦 Modularidade       | Separação clara entre domínio e infraestrutura |
| 🛠️ Flexibilidade      | Infra pode evoluir sem afetar o negócio        |

---

## 🤔 Quando usar?

- O domínio precisa ser **protegido**
- Há múltiplas entradas/saídas: REST, gRPC, CLI, filas...
- Quer **testes automatizados** sem rodar banco, servidor, etc.

---

## 🧱 Estrutura de projeto Go (exemplo)


/cmd                      <- entrypoint (main.go)
/internal
  /order
    /application
      - create_order.go  <- casos de uso (driving port)
    /domain
      - order.go         <- entidades, regras
      - service.go
      - port.go          <- interfaces (driven ports)
    /adapter
      /email
        - smtp.go        <- driven adapter
      /db
        - pg_repo.go     <- driven adapter
      /http
        - handler.go     <- driving adapter


Frases de impacto para entrevista
“A arquitetura hexagonal me permite testar meu domínio sem rodar banco, sem subir servidor.”

“No core da aplicação, a única coisa que existe é lógica de negócio pura, sem saber se o dado vem de HTTP, CLI ou Kafka.”

“Adaptadores podem ser trocados sem que o core perceba — é uma questão de contrato.”

“Quando o domínio muda, o core muda. Quando a tecnologia muda, só o adaptador muda.”


# 🧭 Diferença entre Arquitetura Hexagonal e Clean Architecture

A ideia e quase a mesma, independente de qual direcionamente voce quer se proteger de alterações, postegar eventuais mudanças, quer transformar codigo mais intaveis em estaveis, mexendo o minimo possivel no codigo, todas querem resolver o mesmo problema.

1. O que é Clean Architecture e por que ela é importante?
Resposta esperada:

Clean Architecture é uma abordagem de arquitetura proposta por Robert C. Martin, que visa isolar o núcleo da aplicação (domínio e regras de negócio) de detalhes externos como frameworks, bancos de dados e interfaces de usuário.

Os principais círculos são:

Entidades: regras de negócio puras;

Casos de uso: aplicação das regras para resolver problemas reais;

Interface Adapters: adaptação de entrada/saída (DTOs, controllers, presenters);

Frameworks & Drivers: banco, HTTP, UI, etc.

2. Como você aplica inversão de dependência em Clean Architecture ou Hexagonal Architecture?
Resposta esperada:

Eu aplico a inversão de dependência através de interfaces definidas dentro do core da aplicação (domínio ou camada de caso de uso), e as implementações ficam nas camadas mais externas, como infraestrutura.

Por exemplo, se um caso de uso precisa salvar algo no banco, ele depende de uma interface UserRepository, definida no domínio ou na aplicação. A implementação concreta (PostgreSQL, Mongo, etc) é injetada de fora, normalmente no ponto de entrada (main ou controller).

Isso permite que eu teste o caso de uso isoladamente, com um mock ou stub, sem acoplar à tecnologia usada no banco.

3. Qual a diferença entre a arquitetura em camadas tradicional e Clean Architecture?
Resposta esperada:

A arquitetura em camadas tradicional é baseada em camadas técnicas — por exemplo, apresentação, serviço, repositório, banco de dados. A dependência costuma ser sempre da camada superior para a inferior, o que gera acoplamento, especialmente com a infraestrutura.

Na Clean Architecture, a principal diferença é que as regras de negócio não dependem de nada externo. As dependências são invertidas: frameworks, banco, e interfaces de usuário é que dependem do domínio.

Isso deixa o core isolado, mais testável e resistente a mudanças tecnológicas.

4. Em um projeto real, como você estruturaria um sistema usando Hexagonal Architecture?
Resposta esperada:

Na Arquitetura Hexagonal, a ideia central é isolar o core da aplicação — composto pelas entidades e casos de uso — das interações externas. Para isso, usamos ports (interfaces) e adapters (implementações).

As ports de entrada definem como o mundo externo pode interagir com a aplicação — por exemplo, um UserInputPort com métodos como CreateUser. Um adapter de entrada seria um controller HTTP que recebe a requisição e chama esse port.

As ports de saída definem como o core interage com o mundo externo — como UserRepository ou EmailService. Os adapters de saída são implementações concretas, como acesso ao banco de dados ou envio de e-mails.

Toda a injeção de dependência acontece nos adapters de entrada ou na composição inicial da aplicação.

Em um projeto real, implementei isso num sistema de pagamentos. Tínhamos:

Um adapter de entrada REST que chamava o caso de uso ProcessPayment.

Um adapter de saída que implementava PaymentGateway para se comunicar com o Stripe.

O core era totalmente isolado e testável com mocks dessas interfaces.

5. Em que situações você não usaria Clean Architecture ou Hexagonal Architecture?
Resposta esperada:

Em projetos pequenos, com baixo domínio de negócio ou vida útil curta (como MVPs simples), pode não compensar o esforço de organizar todas as camadas e abstrações.

Nesses casos, eu uso uma estrutura mais simples, mas ainda mantenho alguns princípios — como separação de responsabilidades e uso de interfaces onde faz sentido.

Sempre avalio o custo-benefício de aplicar uma arquitetura robusta versus a complexidade do domínio e o tempo de entrega.

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


Na Clean Architecture, o foco é isolar o domínio da aplicação de qualquer dependência externa como banco de dados, frameworks ou protocolos. Isso torna o sistema mais testável, flexível e resiliente a mudanças.

Em projetos reais, costumo estruturar os pacotes começando com o internal/core, onde ficam as entidades e os casos de uso (usecases). As interfaces (contracts) dos repositórios e serviços externos também ficam ali, para garantir que o core dependa apenas de abstrações.

Fora do core, crio os adapters, por exemplo:

internal/adapter/http com handlers e rotas REST (entrada);

internal/adapter/db com implementações dos repositórios (saída);

Os casos de uso recebem as interfaces como dependência, e no ponto de entrada da aplicação (ex: main.go), eu conecto tudo usando injeção de dependência.

Também costumo isolar código utilitário, DTOs (requests/responses) e middlewares em pacotes separados, sem cruzar com o core.

Essa estrutura ajuda a manter a aplicação organizada e permite escalar ou trocar tecnologias sem tocar no domínio.
# 📦 Padrões de Projeto Usados em Go (para Entrevista Sênior)

---

## 🧪 Test Builder Pattern

### 🧠 Possível pergunta:
> "Que padrão de projeto você vê nesse código de teste?"

### ✅ Resposta sugerida:
"Estamos usando um **Test Builder**, combinado com o padrão de **Mock Object**. O `Setup` atua como um builder fluente para configurar as dependências dos testes, e os mocks usam a lib `testify/mock` para simular chamadas externas e permitir testes isolados e previsíveis."

### Para mostrar maturidade:
"Esse padrão é ótimo para testes porque reduz o acoplamento e melhora a legibilidade, especialmente em serviços com muitas dependências. Também garante que só inicializo o que preciso — o que melhora performance e clareza nos testes."

---

## 🧱 Builder Pattern

- Uso comum com métodos `WithX()` encadeáveis.
- Ideal para objetos complexos com campos opcionais.
- Reduz risco de erro e melhora legibilidade.

### 📌 Perguntas e respostas comuns:

**Como você estrutura objetos complexos com muitos campos opcionais?**  
Uso o padrão Builder com métodos `WithX()` encadeáveis. Isso torna a construção clara, reduz erros e evita construtores gigantes.

**Você já usou WithX() para montar objetos em etapas?**  
Sim, especialmente em testes e mocks. Também usei para construir configs ou estruturas com muitos campos opcionais.

**Qual a diferença entre Builder Pattern e usar structs literais?**  
Com `struct` literals, posso errar a ordem ou deixar campos nulos. O Builder me dá uma API fluente e segura, além de permitir validações na construção.

**Quando usar Builder ao invés de Factory?**  
Uso Builder quando há muitos campos opcionais ou quando quero construir em etapas. Factory é mais direto, para construção simples e imutável.

---

## 🏭 Factory Pattern

- Encapsula regras futuras de criação, se necessário.
- Oculta detalhes de implementação, retornando uma interface.
- Muito usado em Go com funções `NewX()`.

### 🧠 Possível pergunta:
> "Você está usando algum padrão aqui ao criar esse service?"

**Resposta sugerida:**  
"Sim, estou usando o padrão **Factory Function**, que é uma forma idiomática do Factory Pattern em Go. Ele me permite centralizar a criação da struct `categorySerice`, injetando suas dependências e retornando a interface `CategoryService`, o que me dá desacoplamento e facilita testes e substituições futuras."

### Outro exemplo de resposta:
"Percebi que minhas instâncias de repositórios e serviços são criadas com funções do tipo `NewX`, e o que retorno sempre são interfaces. Isso segue o padrão Factory Function e ajuda a manter o domínio desacoplado das implementações concretas."

### 📌 Perguntas e respostas comuns:

**Você costuma usar funções NewX()? Por quê?**  
Sim. Elas encapsulam a criação do objeto, escondem a implementação concreta e permitem aplicar validações ou configurar dependências internas.

**O que te faz esconder a implementação concreta por trás de uma interface?**  
Para desacoplar quem usa da implementação, permitindo mocking, substituição e testes isolados. Também me dá flexibilidade para mudar o backend sem afetar o consumidor.

**Como você lidaria com criação condicional de objetos?**  
Com um Factory Function que recebe uma config ou tipo, e retorna a implementação adequada via interface.

**Factory Function e Factory Pattern são a mesma coisa em Go?**  
Não exatamente. Factory Function é a versão simples (ex: `NewXxx()`), enquanto Factory Pattern pode envolver uma struct `Factory` com lógica mais complexa.

---

## 🤖 Strategy Pattern

- Permite trocar algoritmos ou comportamentos em runtime.
- Evita múltiplos `if`/`switch`.
- Muito usado com interfaces em Go.

### 📌 Perguntas e respostas comuns:

**Quando você usaria o padrão Strategy?**  
Quando preciso encapsular algoritmos ou comportamentos intercambiáveis, como validação, serialização ou escolha de backends.

**Qual a diferença entre Strategy e um simples if/switch?**  
`if/switch` embutem lógica de decisão no código. Strategy delega isso para o momento da injeção via interface.

**Como o Go, que não tem herança, implementa o Strategy Pattern?**  
Com interfaces e composição. Cada struct implementa uma estratégia, e posso injetar dinamicamente.

**Já usou interfaces para encapsular comportamentos intercambiáveis?**  
Sim. Ex: `Notifier` com implementações como `EmailNotifier`, `SMSNotifier`, `PushNotifier`.

**Pode me mostrar um exemplo onde trocou a lógica de execução dinamicamente?**  
Sim. Em integração com REST e gRPC, escolhia a estratégia conforme o ambiente, configurado no boot.

---

## 🔗 Chain of Responsibility

- Permite encadear objetos que tentam processar uma requisição.
- Usado em middlewares ou validações.

### 📌 Vantagens:
- Extensível sem alterar handlers existentes
- Permite interceptar ou interromper o fluxo

### 📌 Perguntas e respostas comuns:

**O que é Chain of Responsibility e onde usou?**  
Encadeamento de handlers que processam requisições. Usei em middlewares HTTP e validações encadeadas.

**Qual a diferença entre Chain of Responsibility e Decorator?**  
Decorator sempre chama o próximo. Chain pode interromper o fluxo (ex: autenticação que falha).

**Como aplicaria esse padrão num pipeline de validações?**  
Crio uma interface `Validator` e encadeio validadores (ex: `RequiredValidator`, `EmailValidator`).

**Como middleware HTTP implementa esse padrão?**  
Cada middleware recebe o próximo handler e decide se passa adiante ou não.

---

## 🎭 Decorator Pattern

- Adiciona funcionalidades extras sem alterar o código original.
- Muito útil para logging, retry, cache, etc.

### 📌 Perguntas e respostas comuns:

**Como diferenciar Decorator de Chain of Responsibility?**  
Decorator sempre chama o próximo. Chain pode parar o fluxo. Decorator estende, Chain decide.

**Quando você usaria Decorator ao invés de estender uma interface?**  
Quando quero adicionar funcionalidades (log, cache, retry) sem mudar o comportamento base.

**Como Decorator ajuda a manter o código Open/Closed?**  
Permite adicionar novos comportamentos empilhando decoradores, sem modificar código existente.

**Qual seria a aplicação desse padrão em serviços de notificação, cache ou logs?**  
- Logs: `LoggerDecorator`
- Cache: `CacheDecorator`
- Retry: `RetryDecorator`

---

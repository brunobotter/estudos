# 📚 Domain-Driven Design (DDD), Clean Architecture e Hexagonal Architecture

---

## 🧠 Domain-Driven Design (DDD)

O DDD propõe que o **código reflita profundamente as regras e conceitos do negócio**, promovendo colaboração entre engenheiros e especialistas do domínio para modelar essas regras corretamente.

### 🎯 Objetivo do DDD

Reduzir a complexidade de sistemas colocando o **negócio no centro do design**, criando modelos que sejam:

- ✅ Claros para o negócio
- ✅ Manuteníveis e escaláveis
- ✅ Flexíveis a mudanças nas regras

---

### 🧩 Componentes Principais do DDD

#### 1. **Ubiquitous Language (Linguagem Ubíqua)**
- Linguagem comum usada por devs e especialistas do negócio
- Termos como `Cliente`, `Pedido`, `Entrega`, `Fatura` usados:
  - No **código**
  - Nas **conversas**
  - Nos **testes**

#### 2. **Modelagem do Domínio**
- O **domínio** é a área do problema que o sistema resolve (ex: logística, finanças, educação)
- Modelos centrais:
  - **Entidades**
  - **Serviços de Domínio**
  - **Agregados**
  - **Objetos de Valor**

---

## 🧼 Clean Architecture

Proposta por **Robert C. Martin (Uncle Bob)**, com foco em:

- Separação de responsabilidades por camadas
- Isolamento do domínio das dependências externas
- Inversão de dependências: o **núcleo não conhece o mundo externo**

---

### 🧱 Camadas da Clean Architecture

| Camada | Papel | Conteúdo |
|-------|--------|----------|
| **Entities (Domínio)** | Regras de negócio puras | Entidades, objetos de valor, validações |
| **Use Cases (Application)** | Orquestra regras de negócio | Serviços de aplicação, coordenam entidades |
| **Interface Adapters** | Traduz dados entre camadas | Controllers, DTOs, repositórios concretos |
| **Frameworks & Drivers** | Interfaces técnicas | HTTP, banco de dados, CLI, etc. (`/cmd`) |

---

## 🛸 Hexagonal Architecture (Ports and Adapters)

Proposta por **Alistair Cockburn**, também chamada de "Ports and Adapters". Seu objetivo é:

- **Isolar o domínio das implementações técnicas**
- **Facilitar testes**
- Tornar o sistema **orientado a comportamentos**, não a tecnologias

### 📌 Componentes

- **Portas (Ports)**: Interfaces de entrada e saída (ex: `Notifier`, `UserRepository`)
- **Adaptadores (Adapters)**: Implementações das interfaces (ex: HTTP handler, gRPC client, repos DB)

---

## 🗂️ Estrutura de Pastas (Resumo)

| Pasta | Papel | Arquitetura | Relacionado a |
|-------|-------|-------------|----------------|
| `/cmd` | Entrada da aplicação (`main.go`) | Clean | **Framework Layer** |
| `/internal/order/domain` | Regras do negócio | DDD + Clean | **Domain Layer** |
| `/internal/order/application` | Orquestradores de regras | Clean | **Use Case Layer** |
| `/internal/order/infrastructure` | Implementações técnicas (DB, HTTP, etc.) | Clean + Hexagonal | **Adapter Layer** |

---

## ✅ Conclusão

Essas arquiteturas e práticas se complementam:

- **DDD** dá o modelo e a linguagem para representar o negócio.
- **Clean Architecture** organiza as responsabilidades em camadas claras.
- **Hexagonal Architecture** promove isolamento e testabilidade, especialmente na integr

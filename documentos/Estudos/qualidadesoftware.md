# 📘 Guia Prático de Cognitive Driven Design (CDD)

CDD (Cognitive Driven Design) é uma abordagem para reduzir a complexidade cognitiva do código, tornando-o mais simples de entender, testar, manter e evoluir. Este guia cobre práticas para medir, controlar e melhorar a qualidade do código utilizando princípios do CDD.

---

## 🧠 O que é Complexidade Cognitiva?

Complexidade cognitiva é o esforço mental necessário para entender o fluxo de um código. Códigos com alta complexidade exigem mais atenção, dificultam manutenção e aumentam a chance de erros.

---

## 📐 Métrica de ICP (Intrinsic Complexity Points)

### 🎯 Objetivo

Quantificar a complexidade para facilitar sua gestão.

### 📊 Pontuação

| Elemento                                          | Pontos |
|--------------------------------------------------|--------|
| `if / else`, `switch`, operador ternário         | +1     |
| Loops (`for`, `while`)                           | +1     |
| `try/catch/finally`                              | +1     |
| Funções com múltiplos argumentos                 | +1     |
| Acoplamento com classes específicas do projeto   | +1     |

### 🧪 Exemplo

```ts
function processPayment(payment) {
  if (!payment.isValid) return false; // +1
  try {
    paymentService.charge(payment);   // +1 (try block)
  } catch (e) {
    logger.error(e);                  // +1 (log de erro)
    return false;
  }
  return true;
}
// ICP total: 3
```

---

## 📉 Avaliando ICPs

- **Projetos novos**: limite inicial de 10–15 ICPs por classe.
- **Projetos legados**: analise as 5 piores classes e defina um limite médio.
- **Por camada**:
  - `Service`: até 15
  - `Controller`: até 10
  - `Repository`: até 5

> 🔍 **Dica**: Ferramentas como SonarQube ou ESLint com plugins personalizados podem ajudar a medir automaticamente.

---

## ✅ Guia de Escrita de Testes Automatizados

### Regras Práticas

- ✅ **Cobertura mínima**: 90%
- 🔗 **Prefira testes integrados**
- 🕐 **Otimize o tempo de execução**
- 🧪 Aplique **MC/DC (Modified Condition/Decision Coverage)**
- ⛳ Use **Boundary Testing**
- 🔁 Explore **Property-Based Testing**
- 📏 Avalie a complexidade dos testes com uma versão de ICP adaptada

### 🧪 Exemplo com MC/DC e Boundary

```ts
it("should reject payment if amount exceeds limit", () => {
  const result = processPayment({ amount: 10001 }); // limite: 10000
  expect(result).toBe(false);
});
```

---

## 🧾 Logging Sistemático

### Quando logar

| Evento                                 | Nível de Log |
|----------------------------------------|--------------|
| Alterações de estado                   | `info`       |
| Chamadas de serviços externos          | `info`       |
| Erros recuperáveis                     | `error`      |

### Exemplo

```ts
logger.info("Iniciando persistência no banco", { userId });
await repository.save(data);
logger.info("Persistência concluída", { userId });
```

> 🛡 **Importante**: Use a biblioteca padrão da empresa que formata e enriquece o log.

---

## 🔗 Maximize a Coesão

- Agrupe funcionalidades relacionadas.
- Métodos que acessam atributos devem estar dentro da mesma classe.

### Exemplo

```ts
class Pedido {
  dataEntrega;

  foiEntregueAntes(outraData) {
    return this.dataEntrega < outraData;
  }
}
```

> 🚫 Evite classes “Frankenstein” com responsabilidades desconexas.

---

## 🛑 Postergue Generalizações

> ⚠️ **Não generalize antes da hora.** Espere pelo menos **3 casos concretos** antes de criar abstrações.

### Exemplo de má prática:

```ts
// Generalização precoce
class BaseProcessor {
  process() {}
}
```

---

## 🎯 Patterns Recomendados

- `Controllers` 100% coesos
- `Services` 100% coesos
- `UseCases` com uma única responsabilidade
- `Form/Request/DTO`: apenas transporte de dados

---

## 📌 Checklist CDD para Pull Requests

- [ ] A classe tem ICP ≤ limite da camada?
- [ ] Os testes cobrem >90% com variedade de entradas?
- [ ] A lógica está coesa e autocontida?
- [ ] Há logging adequado?
- [ ] Há abstrações desnecessárias ou prematuras?

---

## 💬 Conclusão

Este guia é um ponto de partida para reduzir complexidade e melhorar a saúde do código. CDD é uma abordagem incremental, prática e altamente eficaz para manter a qualidade em equipes de desenvolvimento.

> 💡 **Recomendações futuras**:
> - Ferramenta para análise automática de ICPs
> - Integração de CDD ao CI/CD
> - Exemplos antes/depois reais de refatoração


# 📘 Guia Prático de Cognitive Driven Design (CDD)

O CDD (Cognitive Driven Design ou Cognitive-Driven Development) é uma abordagem criada pela Zup Innovation para ajudar desenvolvedores a manter o código simples, compreensível e sustentável, reduzindo a complexidade cognitiva.


---

## 🧠 O que é Complexidade Cognitiva?

É o esforço mental necessário para entender um trecho de código. Ela vai além da quantidade de linhas: depende da forma como o código é escrito, como as decisões estão organizadas e da presença de estruturas que exigem mais raciocínio.

Ex: if, for, try-catch, muitas chamadas encadeadas, funções longas, etc.

---

## 🎯 Fundamentos do CDD

O CDD busca oferecer visibilidade sobre decisões e dificuldades técnicas ao longo do desenvolvimento. Isso é feito por meio de **eventos de marcação cognitiva** (quando o código atinge certos critérios) e **ICP (Intrinsic Complexity Points)**.

- **Marcações cognitivas**: pontos do código onde a complexidade aumenta.
- **ICP**: métrica numérica que representa a complexidade.

Essas práticas incentivam comportamentos como:

- Refatoração precoce
- Discussão técnica baseada em dados
- Simplicidade no design

---

## 📐 Métrica de ICP (Intrinsic Complexity Points)

### Como pontuar?

| Elemento                                          | Pontos |
|--------------------------------------------------|--------|
| `if / else`, `switch`, operador ternário         | +1     |
| Loops (`for`, `while`)                           | +1     |
| `try/catch/finally`                              | +1     |
| Funções com múltiplos argumentos                 | +1     |
| Acoplamento com classes específicas do projeto   | +1     |

> 🔍 A Zup recomenda que desenvolvedores marquem os pontos ICP manualmente com comentários ou anotações nos PRs, usando ferramentas como plugins para VSCode ou CI scripts personalizados.

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

> 💡 A Zup sugere usar os ICPs também em arquivos de teste e em PRs como critério de revisão e melhoria contínua.

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

---

## 🧾 Logging Sistemático

### Quando logar

| Evento                                 | Nível de Log |
|----------------------------------------|--------------|
| Alterações de estado                   | `info`       |
| Chamadas de serviços externos          | `info`       |
| Erros recuperáveis                     | `error`      |

> 🛡 **Importante**: Use a biblioteca padrão da empresa que formata e enriquece o log.

---

## 🔗 Maximize a Coesão

- Agrupe funcionalidades relacionadas.
- Métodos que acessam atributos devem estar dentro da mesma classe.

---

## 🛑 Postergue Generalizações

> ⚠️ **Não generalize antes da hora.** Espere pelo menos **3 casos concretos** antes de criar abstrações.

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

## 📈 Benefícios Comprovados

Segundo experimentos conduzidos pela Zup, o uso de CDD:

- Reduz o tempo de revisão de PRs
- Aumenta a clareza do código para novos membros
- Cria ambiente técnico mais colaborativo
- Estimula refatorações mais seguras e frequentes

---

## 💬 Conclusão

Este guia é um ponto de partida para reduzir complexidade e melhorar a saúde do código. CDD é uma abordagem incremental, prática e altamente eficaz para manter a qualidade em equipes de desenvolvimento.

> 💡 **Recomendações futuras**:
> - Ferramenta para análise automática de ICPs
> - Integração de CDD ao CI/CD
> - Exemplos antes/depois reais de refatoração
> - Estudos de caso do impacto de CDD em times ágeis

---

📚 Fonte: [Zup - Cognitive Driven Development (CDD)](https://zup.com.br/blog/cognitive-driven-development-cdd)

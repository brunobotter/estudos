
# Arquiteturas de Sistemas Enterprise

## Desafios

- Muitos dados  
- Muitos usuários  
- Alta performance  
- Diversidade de produtos  
- Múltiplas equipes  
- Sistemas legados  
- Muitos regulamentos  
- Privacidade e segurança  

## Exemplo de um gateway de pagamentos

- Os pagamentos precisam acontecer o mais rápido possível  
- Não pode haver erros de dados, pagamentos duplicados, etc.  
- Os comerciantes podem ajustar seu pagamento ao extremo; sua configuração é necessária para executar cada transação  
- Há muita comunicação com partes externas (bancos, adquirentes), que podem falhar  
- Há picos durante algumas temporadas  
- Suporte a meios de pagamento de todo o mundo  
- Recursos avançados de relatórios  

## Tópicos

- Escalabilidade  
- Autonomia  
- Infraestrutura e plataforma  
- Qualidade  

---

## Escalabilidade

- Ter um único banco de dados para tudo pode ser difícil  
- Pode ser necessário usar réplicas para aliviar a carga do banco principal  
- É importante aprender a lidar com **consistência eventual**

### Caches

- O acesso ao banco tem custo; se precisamos de velocidade, muitas vezes não podemos consultar o banco toda vez  
- Cache em memória ou distribuído  
- Pode haver inconsistência entre banco de dados e cache  
- É essencial ajustar a estratégia de invalidação de cache  

### Escalabilidade horizontal

- É possível criar mais máquinas e escalar de forma regional para que a máquina fique mais próxima do usuário  

### Processamento assíncrono

- Fazer tudo de forma síncrona pode ser inviável se o processo for caro  
- O usuário tem a sensação de que o processo foi finalizado, mesmo que ainda esteja em andamento  
- Funciona bem com escalabilidade horizontal  
- Você tem uma requisição e uma resposta rápidas  
- É preciso garantir que o item não será processado mais de uma vez  

### Bancos de leitura e escrita separados

- Relatórios podem ser bastante custosos para o banco principal  
- Use um **ETL** (Extrair, Transformar, Carregar)  
- Transmissão de dados para banco exclusivo de relatórios reduz a carga do banco principal  
- Relatórios podem ser gerados em segundo plano  

---

## Autonomia

### Arquitetura baseada em eventos

- Em projetos grandes, há dificuldade de coordenação entre equipes. Uma modificação precisa ser comunicada a várias equipes, exigindo alterações em sprints, tratamento de erros e novas tentativas em diversos pontos  
- A arquitetura baseada em eventos facilita a interação entre serviços, separando produtores e consumidores  
  - Exemplo: uso de filas com vários listeners  
- Desvantagens:
  - Difícil de depurar  
  - Monitoramento mais complexo  

- Do ponto de vista da arquitetura:
  - É necessário planejar o escalonamento das filas  
  - As filas devem estar sempre disponíveis  
  - Deve-se prever o que fazer com itens não processados (DLQ - Dead Letter Queue)

### Arquitetura de microserviços

- Quando há crescimento em complexidade e em número de pessoas, é necessário estabelecer limites claros  
- Microserviços não são uma solução técnica, mas uma solução organizacional para equipes  
- Promovem independência e autonomia da equipe  
- Dependências entre times e sistemas custam tempo e dinheiro  
- **Desvantagem**: distribuir serviços é tecnicamente difícil  

---

## Infraestrutura e Plataforma

- É essencial garantir que todas as aplicações sejam **cloud native** desde o primeiro dia  
- As aplicações devem:
  - Estar prontas para nuvem  
  - Escalar horizontalmente  
  - Ser independentes de localização (ex.: mudança de região)  
  - Ser efêmeras (se o container morrer e reiniciar, não pode haver perda de dados ou estado)  
  - Ser facilmente replicáveis  

---

## Qualidade

- Reuso de código é uma das tarefas mais difíceis de acertar  
  - Compartilhar pouco gera repetição e bugs  
  - Compartilhar demais gera acoplamento  
- Crie bibliotecas e frameworks adequados, com equilíbrio entre reutilização e independência

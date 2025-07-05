Modelo de Documento Formal – Submissão para Promoção

Pod: Horizon
Squad: Jornada Web Cartões

Resumo Executivo
Atuo atualmente como desenvolvedor backend em go no contexo Cartões para mar aberto, com foco em infraestrutura, automação, qualidade de software e resposta rápida a incidentes. Tenho liderado iniciativas com impacto técnico e organizacional, incluindo migração para ambiente governado, evolução de padrões de teste e atuação crítica em produção. Utilizo ativamente ferramentas da StackSpot e agentes de IA internos da Zup para acelerar entregas, garantir boas práticas e controlar custos.

Cases de Impacto


✅ Case 1 – Criação de Infraestrutura no Ambiente Governado com StackSpot e IA


Desafio
Realizar a migração parcial do projeto de Cartões do ambiente auto-governado para o governado do Itaú (JM7), garantindo governança, segurança e continuidade. O desafio envolvia entender o que deveria ser migrado e o que permaneceria no auto-governado, sem impacto nas operações existentes.

Minha Atuação
Conduzi todo o levantamento técnico sobre os componentes que deveriam migrar e os que deveriam permanecer no auto-governado.

Desenhei a infraestrutura sozinho, apresentei ao arquiteto do Itaú e o desenho foi aprovado sem necessidade de ajustes.

Provisionei recursos essenciais como ECS, Fargate, Redis, Secret Manager e Bastion.

Utilizei agentes de IA da Zup para acelerar a escrita do IaC e automatizar validações, aumentando a velocidade e a confiança no provisionamento.

Configurei pipelines no GitHub Actions com StackSpot, incluindo:

Análise de PR automatizada

Análises de segurança

Monitoramento de FinOps AWS

Impacto
A infraestrutura criada pode se tornar referência para outros squads que também irão migrar para o governado, pois foi construída com foco em escalabilidade e boas práticas.

Reduzi significativamente o tempo de entrega: o planejamento inicial previa 2 sprints, mas concluí a entrega em menos de 1 sprint.

Garante governança, segurança e compliance, alinhados com os padrões exigidos no Itaú.

Nível de Participação
Atuei com ownership técnico de ponta a ponta, desde o levantamento, desenho, provisionamento, automação e entrega final.



✅ Case 2 – Evolução e Padronização de Testes com StackSpot AI

Desafio: Resolver gargalo de produtividade e manutenibilidade na base de testes do projeto de Cartões no auto-governado.

Minha Atuação:

Identifiquei o problema de acoplamento e código duplicado (~7.000 linhas)

Adotei o padrão Builder para estruturar os testes.

Refatorei cenários, utilizando StackSpot AI para acelerar o desenvolvimento.

Criei quick commands para facilitar a escrita de novos testes.

Reaproveitei o padrão no novo projeto migrado para o ambiente governado.

Impacto:

Redução expressiva de código duplicado.

Aumento da produtividade e qualidade dos testes.

Estrutura replicável, adotada no novo projeto no governado.

Nível de Participação: Protagonismo técnico e disseminação de boas práticas.


✅ Case 3 – Resolução de Incidente Crítico com reCAPTCHA e Feature Flags

Desafio: Resolver um incidente em produção onde o reCAPTCHA bloqueava indevidamente clientes legítimos.

Minha Atuação:

Diagnóstico rápido através de logs e telemetria.

Realizei ajuste dinâmico do threshold via feature flag.

Implementei eventos de history para monitorar o comportamento do reCAPTCHA.

Atuei em colaboração com o time LATAM e documentei as decisões.

Impacto:

Incidente mitigado sem rollback.

Segurança mantida com ajuste calibrado.

Maior observabilidade para futuras análises.

Nível de Participação: Responsável técnico pela mitigação e solução do incidente.

Ferramentas e Tecnologias Utilizadas
Linguagem: Go

Cloud: AWS (ECS, Fargate, Redis, Secrets, Bastion)

CI/CD: GitHub Actions

StackSpot: Quick Commands, AI Agents


Conclusão
Acredito que minhas entregas demonstram evolução técnica, atuação com ownership, protagonismo e impacto positivo tanto no projeto quanto no time e na organização. 

🎤 Roteiro de Apresentação Oral (3-5 min)
Estrutura Recomendada:
1. Introdução (30 seg)
"Olá, eu sou [Seu Nome], desenvolvedor backend na squad de BPMN – Cartões. Venho buscando a progressão para sênior, e gostaria de compartilhar três cases que demonstram meu protagonismo, impacto técnico e uso eficiente das ferramentas da Zup, como a StackSpot."

2. Case 1 – Infraestrutura no Governado (1 min)
"No primeiro case, atuei na migração parcial do projeto de Cartões para o ambiente governado do Itaú. O desafio era entender o que deveria migrar ou permanecer no auto-governado.
Conduzi o desenho da infraestrutura junto com o arquiteto do Itaú, fiz o provisionamento completo e utilizei agentes de IA e StackSpot para acelerar a escrita de IaC e automação de pipelines.
Essa estrutura se tornou referência para outros squads e reduziu significativamente o tempo de provisionamento."

3. Case 2 – Evolução de Testes com StackSpot AI (1 min)
"O segundo case foi a evolução da base de testes no projeto de Cartões, no auto-governado, onde havia mais de 10.000 linhas de código duplicado.
Adotei o padrão Builder, refatorei os cenários e utilizei StackSpot AI para criar comandos rápidos que facilitaram a escrita dos testes.
Esse padrão foi levado e reaproveitado no novo projeto no governado, mantendo a qualidade e produtividade."

4. Case 3 – Resolução de Incidente em Produção (1 min)
"No terceiro case, lidei com um incidente crítico em produção: o reCAPTCHA estava bloqueando clientes legítimos.
Realizei o diagnóstico rápido via telemetria, ajustei o threshold usando feature flag sem precisar de rollback e implementei eventos para monitorar com precisão.
Essa atuação foi essencial para restaurar o fluxo normal e manter a segurança calibrada."

5. Fechamento (30 seg)
"Esses cases demonstram meu protagonismo, minha capacidade de resolver problemas complexos e meu alinhamento com as práticas de qualidade, segurança e eficiência da Zup. Estou confiante de que estou pronto para assumir o papel de sênior."
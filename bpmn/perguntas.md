🎯 BPMN 1 — Desenvolvimento da infraestrutura no ambiente governado
🔧 Técnicas
Pergunta: Como foi o processo de levantar e provisionar infraestrutura no ambiente governado?

Resposta:
Comecei levantando os requisitos junto ao time de arquitetura e compliance, entendendo o que poderia ou não ser criado dentro do ambiente governado. Em seguida, desenhei a solução com foco em segurança e escalabilidade, envolvendo ECS com Fargate, Redis, Secret Manager e Bastion. Utilizei agentes de IA da Zup para acelerar o provisionamento e garantir que tudo estivesse de acordo com os padrões da empresa.

Pergunta: Quais desafios você enfrentou nessa entrega?

Resposta:
O maior desafio foi equilibrar velocidade de entrega com as restrições de segurança do ambiente governado. Muitas validações manuais e processos burocráticos poderiam atrasar, mas com a ajuda dos agentes de IA, consegui acelerar a validação de alguns recursos e automatizar partes do provisionamento.

💬 Comportamentais
Pergunta: Como você garantiu alinhamento com diferentes times durante a criação dessa infraestrutura?

Resposta:
Marquei rituais de alinhamento com arquitetura e segurança logo no início. Documentei todo o processo, deixei claro o que seria entregue em cada fase, e usei canais assíncronos para garantir fluidez na comunicação. Esse alinhamento evitou retrabalho e garantiu que todos estivessem na mesma página.

🎯 BPMN 2 — Evolução de testes com Stackspot AI
🔧 Técnicas
Pergunta: Que melhorias você trouxe para os testes do projeto?

Resposta:
Implementei o padrão de projeto Builder nos testes, principalmente na instanciação dos services. Isso eliminou redundância e melhorou a legibilidade. Refatorei uma suíte de testes complexa do fluxo One Itaú, reduzindo 2800 linhas de código. Isso impactou diretamente na produtividade e velocidade de deploys.

Pergunta: Como o Stackspot AI ajudou nesse processo?

Resposta:
Usei o Stackspot AI para gerar scaffolds de testes e converter parte da documentação e exemplos para código real. Com isso, padronizamos o formato dos testes e conseguimos escalar o modelo para o restante do time com mais facilidade.

💬 Comportamentais
Pergunta: Como você convenceu o time a adotar essas melhorias?

Resposta:
Mostrei os ganhos em números: redução de linhas de código, aumento da cobertura e facilidade de manutenção. Fiz um pair programming com alguns devs e transformei os aprendizados em documentação com Stackspot, o que facilitou a adoção em outras squads também.

🎯 BPMN 3 — Recaptcha + Threshold + Resolução de Incidente
🔧 Técnicas
Pergunta: Que medidas técnicas você tomou ao implementar o Recaptcha?

Resposta:
Implementei um threshold de 0.1 para bloquear bots, o que protegeu nossos formulários de spam. Além disso, adicionei eventos de metrificação no sistema de history para acompanhar os casos de sucesso, erro e início de execução do Recaptcha. Isso nos deu mais visibilidade e rastreabilidade.

Pergunta: Como você lidou com o incidente em produção relacionado a essa entrega?

Resposta:
Assim que foi reportado o comportamento inesperado, ajudei a isolar o problema rapidamente, identifiquei que o threshold estava bloqueando alguns usuários válidos e ajustei o valor. Mantive a comunicação constante com stakeholders e escrevi um post-mortem com aprendizados e próximos passos.

💬 Comportamentais
Pergunta: Como você demonstrou senioridade ao lidar com esse incidente?

Resposta:
Atuei com calma, organizei o time em torno da contenção, assumi a frente na análise técnica e garanti a comunicação com os stakeholders. Após a resolução, documentei o incidente, propus alertas proativos e automatizei parte da validação para evitar recorrência.
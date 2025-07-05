Arquiterura em camadas

Uma funcionalidade que to implementando nesse ponto do meu codigo, e nesse ponto do codigo eu converso com varias outras partes do meu sistema, bd, serviços externos, libs e etc.

Vantagem que por exemplo quero implementar um serviço de email mas ainda nao defini qual lib de email vou usar, eu posso criar uma casca que vai implementar esse serviço

Mas no final todas as arquiteturas querem resolver a mesma coisa , tudo que e relativo ao dominio, negocio querem proteger, deixar o mais estavel

1° Tem que ser indepemdente de framework, qualquer mudança no framework nao deveria impactar seu dominio
2° Tem que ser testavel
3° Independente de UI
4° Independente de banco de dados
5° Independente de qualquer agente externo

As camadas externas podem depender das camadas internas, mas nunca ao contrario

Diferença entre arquitetura hexagonal com clean

Hexagonal:
👉 O mundo externo conversa com a aplicação através de portas (interfaces) e adaptadores.
👉 Foco principal: isolar o core da aplicação das entradas e saídas.

Clean Architecture:
👉 O sistema é organizado em círculos concêntricos com dependências que sempre apontam para dentro.
👉 Foco principal: centralizar a regra de negócio e tornar tudo o mais independente possível de infraestrutura.

Hexagonal: "Minha aplicação tem portas para o mundo. Quem quiser falar comigo precisa passar por elas."

Clean: "Minha aplicação é um castelo com muralhas. As regras mais importantes estão no centro, e tudo ao redor existe apenas para servir o core."
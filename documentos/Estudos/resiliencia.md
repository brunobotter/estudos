Resiliencia

Resiliência em código é a capacidade do software continuar funcionando (ou se recuperar rapidamente) mesmo quando ocorrem falhas, problemas ou situações inesperadas.

Exemplos práticos de resiliência no código:
Situação	Código não resiliente	Código resiliente
API externa falha	Lança exceção e derruba o sistema	Faz retry, fallback ou ignora com log
Banco de dados indisponível	Crasha	Usa cache temporário, reprocessa depois
Input inválido do usuário	Quebra	Valida e dá mensagem de erro amigável
Picos de tráfego	Estoura recursos	Usa fila, balanceamento, circuit breaker

Se deve suportar resiliencia logo no inicia, logo na decisão de arquitetura
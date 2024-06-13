# Projeto de Servidor de Autenticação e Autorização

Este documento descreve os requisitos funcionais e não funcionais para o desenvolvimento de um 
servidor de autenticação e autorização. O objetivo do sistema é fornecer um 
serviço seguro e escalável para autenticar usuários e gerenciar suas permissões em um ambiente de nuvem. 
O sistema será desenvolvido em Go e seguirá os princípios SOLID e o padrão Port and Adapters.

## Requisitos Funcionais

1. Garantir segurança no acesso a recursos.
2. Tipos de usuários:
    - Administrador
    - Usuário
3. Funcionalidades específicas:
    - Administrador: criar usuários, gerenciar permissões, monitorar atividade do usuário, gerenciar grupos de usuário, recuperar e revogar contas.
    - Usuário: monitorar suas próprias atividades, alterar dados pessoais.
4. Métodos de autenticação: OAuth usando Google e GitHub. 

5. Informações de usuário: username, email, tipo de autenticação escolhida, data de criação de usuário, atualização e deleção, grupo de usuário, histórico de logins, se usuário está logado, status da conta, lista de permissões.

6. Operações dos usuários: login, logout, atualizar informações do perfil, visualizar permissões, acessar recursos autorizados.

7. Processo de cadastro e login: O usuário passará um username e um método de autenticação. No login, basta o usuário acessar pelo método de autenticação escolhido no cadastro.

8. Integrações previstas: Deve ser integrado a um servidor de chat, comunicando-se através de REST e gRPC.

## Requisitos Não Funcionais

1. Desempenho esperado: Altamente escalável conforme a quantidade de usuários.

2. Requisitos de segurança: Sem senhas, somente OAuth. Conexão com certificado SSL/TSL.

3. Disponibilidade e confiabilidade: Disponível 24/7, com um tempo de inatividade mínimo para manutenção. Meta de disponibilidade de 99,9%.

4. Escalabilidade: Containerizado, seguindo os princípios SOLID e o padrão Port and Adapters.

5. Manutenção e suporte: CI/CD com code review e testes. Commits pequenos e contínuos.

6. Conformidade ou regulamentação: Emails encriptados para conformidade com GDPR e LGPD.

7. Arquitetura geral: Microserviço com Ports and Adapters, API em REST e gRPC, PostgreSQL como banco de dados principal, Redis como cache.

8. Usabilidade ou experiência do usuário: Baseado em terminal.

9. Restrições tecnológicas ou de infraestrutura: Desenvolvido em Go, implantado em AWS ou GCP. Deve ser containerizado.

10. Backup e recuperação de dados: PostgreSQL com backup diário para um file storage.
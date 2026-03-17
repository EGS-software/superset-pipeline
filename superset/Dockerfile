# Usa a imagem oficial como base
FROM apache/superset:latest

# Muda para o usuário root para ter permissão de instalar no sistema
USER root

# Injeção Cirúrgica: Instala o driver diretamente DENTRO do cofre do Ambiente Virtual do Superset
RUN /usr/local/bin/pip install --target /app/.venv/lib/python3.10/site-packages psycopg2-binary

# Volta para o usuário padrão por segurança
USER superset
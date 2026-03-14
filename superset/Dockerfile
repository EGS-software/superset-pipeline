# Usa a imagem oficial como base
FROM apache/superset:latest

# Muda para o usuário root para poder instalar pacotes
USER root

# Instala o driver do PostgreSQL (ou outro banco que o Go for alimentar)
RUN pip install --no-cache-dir psycopg2-binary

# Volta para o usuário padrão por segurança
USER superset
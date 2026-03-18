# 📊 Superset Data Pipeline | Pokémon Analytics

![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/postgresql-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Apache Superset](https://img.shields.io/badge/apache%20superset-%2320A6D8.svg?style=for-the-badge&logo=apache%20superset&logoColor=white)

Um pipeline de dados completo, conteinerizado e automatizado para extração, armazenamento e visualização de dados do universo Pokémon. Construído como projeto de apresentação em Business Intelligence e Engenharia de Dados.

## 🏗️ Arquitetura do Projeto

A infraestrutura foi desenhada utilizando o conceito de **Infrastructure as Code (IaC)** através do Docker Compose, garantindo que o ambiente seja 100% isolado, reprodutível e compatível com múltiplas arquiteturas.

O fluxo de dados é dividido em 3 camadas principais:

1. **Extract, Transform, Load (ETL):** Uma aplicação escrita em **Go** responsável por buscar e processar os status e tipos dos Pokémons e realizar a carga no banco de dados.
2. **Data Warehouse:** Um banco **PostgreSQL** (`pokemon_dw`) dedicado ao armazenamento persistente dos dados estruturados.
3. **Business Intelligence (BI):** O **Apache Superset** conectado ao Data Warehouse, responsável por fornecer a interface visual, execução de queries SQL e criação de Dashboards interativos.

### 🌟 Destaques da Arquitetura e Engenharia

* **ETL Idempotente (Upsert):** A rotina de carga em Go foi construída utilizando comandos `ON CONFLICT DO UPDATE`. Isso garante a idempotência do pipeline: o script pode ser executado múltiplas vezes de forma segura, atualizando registros existentes sem gerar duplicidade no Data Warehouse.
* **Otimização Extrema com Multi-stage Builds:** O serviço de extração utiliza *Multi-stage builds* no Docker (`golang:1.25-alpine` -> `alpine:latest`), separando o ambiente de compilação do ambiente de execução. O resultado é uma imagem final ultra-leve e segura, contendo apenas o binário estático.
* **Resiliência e Rate Limiting:** A aplicação foi desenhada para lidar com falhas de rede. Implementa o padrão *Retry (Wait-for-it)* para aguardar a inicialização do PostgreSQL e aplica *Throttling* (pausas intencionais em milissegundos) nas requisições à PokéAPI para evitar bloqueios por excesso de tráfego.
* **Isolamento de Ambiente (Python/Superset):** O `Dockerfile` do BI foi customizado cirurgicamente para injetar o driver `psycopg2-binary` direto no ambiente virtual (`/app/.venv`) da imagem oficial, blindando o projeto contra falhas de compilação em processadores ARM/Apple Silicon.
* **Persistência Estratégica:** Configuração de Docker Volumes locais isolados. Mesmo que os contêineres sofram *down*, o banco de dados principal (`pokemon_dw`) e o banco de metadados do Superset (SQLite interno) permanecem intactos.
---

## 🚀 Como Executar o Projeto

### Pré-requisitos
* [Docker](https://docs.docker.com/get-docker/) e Docker Compose instalados.

### Passo 1: Subir a Infraestrutura
Clone o repositório e inicie os contêineres na raiz do projeto:

```bash
docker compose up --build -d
```
*Aguarde alguns segundos para que o script em Go finalize a carga de dados no PostgreSQL.*

### Passo 2: Inicializar o Apache Superset
Como os dados são persistidos via volume local, na **primeira execução** é necessário configurar o banco interno do Superset. Execute os comandos abaixo no seu terminal em sequência:

**1. Atualizar o banco de dados interno:**
```bash
docker exec -it superset_dashboard superset db upgrade
```

**2. Criar o usuário Administrador:**
```bash
docker exec -it superset_dashboard superset fab create-admin \
              --username admin \
              --firstname Admin \
              --lastname Superset \
              --email admin@superset.com \
              --password admin
```

**3. Inicializar as permissões e dependências:**
```bash
docker exec -it superset_dashboard superset init
```

### Passo 3: Acessar a Interface
Acesse `http://localhost:8088` no seu navegador e faça o login com as credenciais criadas (`admin` / `admin`).

Para conectar o banco de dados PostgreSQL e começar a criar seus gráficos, vá em **Settings > Database Connections**, adicione o PostgreSQL e utilize a seguinte *SQLAlchemy URI*:

```text
postgresql://admin:password123@db-postgres:5432/pokemon_dw
```

---

## 📁 Estrutura de Diretórios

```text
.
├── docker-compose.yml      # Orquestração de todos os serviços (DB, Superset, ETL)
├── etl-go/                 # Aplicação principal de extração e carga de dados em Go
│   └── cmd/main.go         # Script de conexão e lógica do banco de dados
└── superset/               # Configurações customizadas do Apache Superset
    └── Dockerfile          # Injeção do driver PostgreSQL no ambiente virtual
```

---

## 👥 Equipe
Projeto desenvolvido por **João Victor Benetti Filipim**, **Mauro Pezzetta Roncata**, **Tiago Andrei de Alemida Mendonça**.

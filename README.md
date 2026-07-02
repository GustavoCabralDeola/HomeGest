# GestHome (Versão Golang)

Criado para ser simples e direto ao ponto, o **GestHome** é o meu sistema de gestão financeira doméstica. O objetivo aqui foi montar uma ferramenta que realmente ajude a entender para onde o dinheiro está indo, sem complicação e com um visual profissional. 

Esta versão do backend foi completamente traduzida de C# / .NET para **Golang**, mantendo as mesmas regras de negócio e a mesma arquitetura.

---

## O que o sistema faz?

### Gestão de Pessoas
Dá para cadastrar todo mundo de casa. A lista é prática: você consegue editar nome ou idade ali mesmo na tabela (edição inline) ou apagar alguém se precisar. Ah, se você apagar uma pessoa, o sistema limpa todas as transações dela automaticamente para não deixar lixo no banco (Cascade Delete).

### Controle de Transações
Aqui é o coração do app. Você registra o que entrou (**Receita**) e o que saiu (**Despesa**).
Coloquei algumas regras de "pé no chão":
- Menor de idade não pode registrar receita (regra da casa!).
- O sistema te avisa se você tentar colocar uma despesa em uma categoria que é só de receitas.
- O visual ajuda: o que é entrada fica verdinho, o que é saída fica em destaque vermelho.

### Resumos que ajudam
Tem duas páginas de totais que são o "pulo do gato":
1. **Totais por Pessoa**: Para saber quem está gastando mais ou quem trouxe mais receita no mês.
2. **Totais por Categoria**: Ótimo para ver se o gasto com "Lazer" está passando da conta ou quanto foi para "Alimentação".

---

## Arquitetura do Sistema (Golang)

O projeto foi reescrito em Go seguindo o padrão de **Clean Architecture**, o mesmo que existia no projeto original em C#. O código está organizado nas seguintes camadas (pacotes `internal`):

- **Domain (`internal/domain`)**: O coração do software. Aqui ficam as entidades (`Categoria`, `Pessoa`, `Transacao`), os Enums (`CategoriaFinalidade`, `TipoTransacao`) e as interfaces dos repositórios.
- **Application (`internal/application`)**: Onde moram as regras de negócio. Contém os DTOs para tráfego de dados e os Services (`CategoriaService`, `PessoaService`, `TransacaoService`) que executam as validações.
- **Infrastructure (`internal/infra`)**: Acesso a dados. A configuração do banco e a implementação dos repositórios utilizando **GORM** para acesso ao SQL Server.
- **API (`internal/api`)**: A camada de comunicação externa. Onde o roteador HTTP (`go-chi`) e os Handlers recebem as requisições, chamam a camada de Application e devolvem o JSON pro frontend.

---

## O que usei para construir

### No Backend (Golang)
- **Go 1.23**: Linguagem principal, escolhida pela sua performance, concorrência e simplicidade.
- **go-chi**: Um roteador HTTP leve, idiomático e muito eficiente.
- **GORM**: Um ORM fantástico para mapear as structs do Go para o banco de dados e facilitar as queries complexas.
- **SQL Server**: O banco de dados original foi mantido.

### No Frontend
- **React + TypeScript**: Para uma interface rápida e sem bugs de tipo.
- **Vite**: Porque ninguém merece esperar build demorado.
- **CSS Puro (Vanilla)**: Fiz questão de não usar bibliotecas de componentes. Cada botão e card foi estilizado do zero.

---

## Como rodar na sua máquina

### 1. Preparando o Terreno
Você vai precisar do **Go** (versão 1.23 ou superior) instalado na sua máquina, além do **Node.js** para rodar o frontend.

### 2. Conectando ao banco de dados
O projeto utiliza SQL Server. Na raiz do backend em Go, crie um arquivo `.env` (você pode copiar do `.env.example`) e configure suas credenciais:

```env
DB_HOST=localhost
DB_PORT=1433
DB_USER=sa
DB_PASSWORD=SuaSenhaAqui
DB_NAME=GestHome
SERVER_PORT=8080
```

### 3. Rodando o Backend (API)
Abra o terminal na raiz do projeto (onde está o arquivo `go.mod`) e baixe as dependências:
```bash
go mod tidy
```

Depois, inicie a aplicação:
```bash
go run cmd/server/main.go
```
*A API subirá, por padrão, na porta `8080` e aplicará as migrations no banco automaticamente.*

### 4. Rodando o Frontend (Web)
(Assumindo que você tenha a pasta do frontend `gesthome.web` separada)
Entre na pasta do frontend:
```bash
npm install
```
Não se esqueça de ajustar a URL da API no frontend (caso use arquivos `.env` lá) apontando para `http://localhost:8080`.

Inicie o frontend:
```bash
npm run dev
```

---
Feito com muito café e foco por **Gustavo Deola** (agora em Golang!).

# Atividade Docker com GO
Aplicação de anotações REST que possui quatro endpoints, sendo possível realizar um CRUD básico (Cadastrar, editar, remover e visualizar), tendo como propriedades "Text" e o "Id" construído na linguagem GO.  

## 📋 Pré-requisitos
Para que a aplicação funcione, é necessário, estar instalado o [Docker](https://www.docker.com/get-started) e o [Docker Compose](https://docs.docker.com/compose/install/) para que o compose file seja roda e o ambiente seja corretamente criado.

## 🔧 Instalação
No terminal, clone o projeto:
```
git clone https://github.com/matheussASM/pre-atividade-02.git
```
Agora é necessario realizar o build das imagens docker, então primeiro entre na pasta:
```
cd pre-atividade-02/api/notes-api/
```
Dentro da pasta realize o comando:
```
docker image build -t pre-atividade-api .
```
Após o build ser concluido, entre na pasta db através de:
```
cd ../../db
```
Agora realize o build dessa nova imagem:
```
docker image build -t pre-atividade-db .
```
Agora retornamos a pasta principal para que possamos executar o compose file:
```
cd ..
```
Para finalizar execute:
```
docker-compose up -d
```
Pronto sua API já está pronta para uso.

## 📦 Consumindo a API
Caso queira consumir o CRUD da API, pode ser utilizada alguma ferramenta como [postman](https://www.postman.com/) ou [insomnia](https://insomnia.rest/).

Para buscar todos as notas, deve ultilizar o metodo GET:
```
http://localhost:5000/api/v1/notes/
```
Para buscar uma anotação com um ID específico:
```
http://localhost:5000/api/v1/notes/{id}
```
Para criar uma anotação, com o metodo POST:
```
http://localhost:5000/api/v1/notes/
```
Os dados da nota devem ser enviados no corpo da requisição:
```
{
	"text":"Exemplo"
}
```
Para apagar uma nota, com o metodo DELETE:
```
http://localhost:5000/api/v1/notes/{id}
```
### :computer: Consumindo a API via terminal
Outra alternativa de consumo é via terminal
Para buscar todas as notas:
```
curl http://localhost:5000/api/v1/notes/
```
Para criar uma nova nota:
```
curl -X POST http://localhost:5000/api/v1/notes/ -H 'Content-Type: application/json' -d '{"text":"Exemplo"}'
```

## 🛠️ Construído com

GO - É uma linguagem de programação criada pela Google e lançada em código livre em novembro de 2009. É uma linguagem compilada e focada em produtividade e programação concorrente

PostgreSQL - É um SGBDs (Sistema Gerenciador de Bancos de Dados) de código aberto

Docker - Docker é um conjunto de produtos de plataforma como serviço que usam virtualização de nível de sistema operacional para entregar software em pacotes chamados contêineres. 

Docker compose - É o orquestrador de containers da Docker.

## ✒️ Autores

Desenvolvimento, Documentação - Matheus Moreira

Linkedin - https://www.linkedin.com/in/matheus-antonio-9831b9157/

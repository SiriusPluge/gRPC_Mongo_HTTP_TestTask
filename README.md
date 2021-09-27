# A CRUD - simple project
## Performs CRUD (Create, Read, Update and Delete) operations:
1. HTTP-server
2. gRPC-server
3. Client
2. Using MongoDB  


## To use it, you must enter the following commands:
1. Cloning the repository
`cd $HOME/src/github.com/`  
`git clone github.com/SiriusPluge/gRPC_Mongo_HTTP_TestTask`
2. Docker build:
'sudo docker build -t my-server -f Dockerfile .'
'docker-compose up -d'

## Congratulations, you have launched the project!!!

## Running the app
Make sure MongoDB server is up and running on port `27017`.
Make sure gRPC server is up and running on port `50051`.
Make sure HTTP server is up and running on port `4112`.

## to conduct testing gRPC method, you must:
1.Open two terminals and run:
- Create book: ' go run client/main.go create -a "PASTE author's name" -n "PASTE book title" -t "PASTE Tag" ';
- Get/Read book: ' go run client/main.go read -i "PASTE book id" ';
- Update book: ' go run client/main.go update -i "PASTE book id" -a "New author" -n "New title" -t New Tag" ';
- Delete book: ' go run client/main.go delete -i "PASTE book id" '.
2. The result is in the console!

## To conduct testing HTTP method, you must:
1. Open the post man collection, which is located in the root folder of the project and send requests to the HTTP server
1.1. Открыть коллекцию post man, которая находится в корневой папке проекта и направить запросы на сервер HTTP;

## it is possible to send the following requests:
- Create book/POST: "localhost:4112/api/book";
- Delete book/DELETE: "localhost:4112/api/book/delete";
- Get book/GET: "localhost:4112/api/book/get";
- Update book/PUT: "localhost:4112/api/book/put".

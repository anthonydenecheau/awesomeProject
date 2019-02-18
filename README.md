[DONE]
* je crée un rest-api connecté à une base postgreSql
* je respecte une organisation des packages (go clean architecture)
* je déploie une image docker (via travis-ci)
* je démarre un container docker

[IN PROGRESS]
* j'ajoute un environnement de configuration
* j'ajoute les test de mes packages

[TODO]
* je finalise la construction ws-dog
* je crée la documentation (swagger, test-driven)
* je migre le projet vers un environnement GCP
    - j'ajoute des informations de logging, tracing

[Librairies]
* Middleware : labstack/echo/
* DAO : go-pg/pg
* Configuration : Viper, Cobra

[References]
* Google Cloud:
    -   https://github.com/abronan/todo-grpc/blob/master/main.go
* Architecture
    - https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
    - https://github.com/hirotakan/go-cleanarchitecture-sample
    - https://github.com/bxcodec/go-clean-arch

# README #

The News API using GO Language, MySql and redis

Fitur :
* RBAC (Role Based Access Control)
* Catching news list using redis
* Auto resize image and auto create thumbnail image
* Storage image to AWS S3
* Security
* Api documentation using swagger
* Unit Testing and Rest api testing

There is 2 (two) role, Editor and Writer. Editor can access all api. Writer only can create new news and edit own news. Guest can read topics and read news.

Guest :
* GET /news
* GET /news/{id}
* GET /topics
* GET /topics/{id}

Writer :
* POST /news
* PUT /news/{id} (but only the news he wrote himself)

Editor :
* POST /news
* PUT /news/{id}
* PUT /news/{id}/publish
* DELETE /news/{id}
* POST /topics
* PUT /topics/{id}
* DELETE /topics/{id}

for demo documentation, see https://api-news.rijalasepnugroho.com/documentation/

### What is this repository for? ###

* API for News
* 1.0
* [Learn Markdown]

### How do I get set up? ###

This API using Golang. Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. 

to install go, https://golang.org/doc/install

#### Dependencies : ####
* go get github.com/spf13/viper
* go get github.com/gorilla/handlers
* go get github.com/gorilla/mux
* go get golang.org/x/crypto/bcrypt
* go get github.com/dgrijalva/jwt-go
* go get github.com/go-sql-driver/mysql
* go get github.com/aws/aws-sdk-go
* go get github.com/gosimple/slug
* go get gopkg.in/redis.v5
* go get golang.org/x/net/html
* go get github.com/nfnt/resize

File Configuration on ./config/config.json

You can run application with going the directory api-news and command : go run main.go

For api documentation, you can open on the browser: http://localhost:8080/documentation/

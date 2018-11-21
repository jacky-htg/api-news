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
* Queue Messaging using redis, every new publish news will be push to queue messaging
* Websocket to broadcast new publish news

### What is this repository for? ###

* API for News
* 1.0
* [Learn Markdown]

### RBAC ###

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

### Web Socket ###

Websocket is used to broadcast new publish news. WebSocket address is : ws://localhost:8080/newssocket. Here is sample script on client when consum a websocket. Create a file client.html on your document root and open it using your favorite browser : http://localhost/client.html

```
<script>
var ws = new WebSocket("ws://localhost:8080/newssocket");

ws.onopen = function() {
  ws.send("hello server!")
}

ws.onmessage = function(event) {
  var m = event.data;
  if (m === "pong") console.debug("Received message", m);
  else notifyMe(m);
}

ws.onerror = function(event) {
  console.debug(event)
}

function notifyMe(message) {
  // Let's check if the browser supports notifications
  if (!("Notification" in window)) {
    alert("This browser does not support system notifications");
  }

  // Let's check whether notification permissions have already been granted
  else if (Notification.permission === "granted") {
    // If it's okay let's create a notification
    var notification = new Notification(message);
  }

  // Otherwise, we need to ask the user for permission
  else if (Notification.permission !== 'denied') {
    Notification.requestPermission(function (permission) {
      // If the user accepts, let's create a notification
      if (permission === "granted") {
        var notification = new Notification(message);
      }
    });
  }

  // Finally, if the user has denied notifications and you 
  // want to be respectful there is no need to bother them any more.
}
</script>
```

### Demo ###

for demo documentation, see https://api-news.rijalasepnugroho.com/documentation/

Api Key : n3WsAp1D3v

Editor : email (editor@gmail.com) password (1234)

Writer : email (writer@gmail.com) password (1234)


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
* go get github.com/gorilla/websocket
* go get github.com/nfnt/resize
* go get github.com/adjust/rmq

File Configuration on ./config/config.json

You can run application with going the directory api-news and command : go run main.go

You can test the api using command : go test -v

For api documentation, you can open on the browser: http://localhost:8080/documentation/

FROM golang:1.9.7-alpine
RUN apk add --update --no-cache git
RUN go get github.com/spf13/viper 
RUN go get github.com/gorilla/handlers 
RUN go get github.com/gorilla/mux 
RUN go get golang.org/x/crypto/bcrypt 
RUN go get github.com/dgrijalva/jwt-go 
RUN go get github.com/go-sql-driver/mysql 
RUN go get github.com/aws/aws-sdk-go 
RUN go get github.com/gosimple/slug 
RUN go get gopkg.in/redis.v5 
RUN go get github.com/gorilla/websocket 
RUN go get github.com/nfnt/resize 
RUN go get github.com/adjust/rmq 

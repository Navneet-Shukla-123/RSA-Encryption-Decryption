FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD [ "go","run","all" ]
FROM golang:1.18-alpine
RUN apk add build-base
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

EXPOSE 3000

CMD [ "/app/server" ]
FROM golang:1-alpine3.19 as builder

WORKDIR /go-app

COPY go.mod go.sum ./

RUN go mod download && \ 
    go get -u github.com/cosmtrek/air && \
    go install github.com/cosmtrek/air

COPY . ./

CMD [ "air", "-c", ".air.toml" ]
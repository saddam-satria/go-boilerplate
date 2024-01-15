FROM golang:1-alpine3.19 as builder

WORKDIR /app/go-boilerplate 

COPY go.mod go.sum ./

RUN go mod download && \ 
    go get -u github.com/cosmtrek/air && \
    go install github.com/cosmtrek/air

COPY cmd ./cmd

COPY pkg ./pkg

COPY .air.toml ./

COPY main.go ./

COPY migrate.go ./

CMD [ "air", "-c", ".air.toml" ]
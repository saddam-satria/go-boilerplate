FROM golang:1-alpine3.19 as builder

WORKDIR /app/go-boilerplate 

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd

COPY pkg ./pkg

COPY main.go ./

COPY migrate.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o dist/main .


FROM alpine:latest

WORKDIR /app/go-boilerplate

COPY --from=builder /app/go-boilerplate/dist ./dist

CMD ["./dist/main"]


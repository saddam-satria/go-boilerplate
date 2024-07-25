FROM golang:1-alpine3.19 as builder

WORKDIR /go-app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

COPY main.go migrate.go seed.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o dist/main .

FROM alpine:latest

WORKDIR /go-app

COPY --from=builder /app/posq-be/dist ./dist

CMD ["./dist/main"]
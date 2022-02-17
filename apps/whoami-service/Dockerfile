FROM golang:1.17-alpine as builder

WORKDIR /app
COPY . .

RUN go build -o ./dist/api

FROM golang:1.17-alpine as runner

WORKDIR /app
COPY --from=builder /app/dist ./

EXPOSE 5000

ENTRYPOINT ["/app/api"]

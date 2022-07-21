FROM golang:1.14.15
COPY main.go /app/golang-app/main.go
CMD ["go", "run", "/app/golang-app/main.go"]
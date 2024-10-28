FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Deploy the application binary into a lean image
FROM scratch

COPY --from=build /app/main /bin/main
EXPOSE 8080

CMD ["/bin/main"]
# 1 stage - build
FROM golang:alpine AS build

WORKDIR /src

COPY . .

RUN go build -o main cmd/main.go
RUN go build -o migrate internal/migrate/migrate.go

FROM alpine:latest AS run

WORKDIR /src

COPY --from=build /src/main .
COPY --from=build /src/migrate .

EXPOSE 8080
RUN chmod +x main
RUN chmod +x migrate
CMD ["sh", "-c", "./migrate && ./main"]
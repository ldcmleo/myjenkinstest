FROM golang:1.24 AS build
WORKDIR /app/src
COPY . .
ENV CGO_ENABLED=0
RUN go build -o appapi ./main.go

FROM alpine:latest AS runtime
COPY --from=build /app/src/appapi ./
EXPOSE 8081/tcp
ENTRYPOINT ["./appapi"]
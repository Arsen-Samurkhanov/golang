FROM golang:1.22-alpine AS build
WORKDIR /app
COPY  . .
RUN go build -o app


FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/* .
EXPOSE 8080
CMD ["./app"]
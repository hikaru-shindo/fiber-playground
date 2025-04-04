FROM golang:1.24.2-alpine AS build

WORKDIR /build

RUN echo "app:x:10001:10001:App User:/:/sbin/nologin" > /etc/minimal-passwd

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s" -o api -v ./cmd/api

FROM scratch AS prod

ENV SERVER_PORT=8080

COPY --from=build /etc/minimal-passwd /etc/passwd
COPY --from=build /build/api /api

USER app
EXPOSE ${SERVER_PORT}

ENTRYPOINT ["/api"]

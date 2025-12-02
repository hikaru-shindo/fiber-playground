
FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.25.5-alpine AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /build

RUN apk add --no-cache gcc musl-dev
RUN echo "app:x:10001:10001:App User:/:/sbin/nologin" > /etc/minimal-passwd

RUN mkdir ./database
RUN chown 10001:10001 ./database

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags='-s -extldflags "-static"' -o api -v ./cmd/api

FROM scratch AS prod

ENV SERVER_PORT=8080

COPY --from=build /etc/minimal-passwd /etc/passwd
COPY --from=build /build/database /database
COPY --from=build /build/api /api

VOLUME /database
USER app
EXPOSE ${SERVER_PORT}

ENTRYPOINT ["/api"]

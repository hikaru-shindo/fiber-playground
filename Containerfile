FROM docker.io/library/golang:1.23-alpine as builder

RUN mkdir /build && mkdir /out
WORKDIR /build

COPY . .

RUN go build -o /out/server cmd/main.go

FROM scratch

COPY --from=builder --chmod=755 /out/server /server

EXPOSE 3000
ENTRYPOINT [ "/server" ]

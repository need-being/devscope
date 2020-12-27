FROM golang:alpine as builder
COPY . /src
WORKDIR /src
RUN go build ./cmd/devscope

FROM alpine
COPY --from=builder /src/devscope /bin/devscope
ENTRYPOINT [ "/bin/devscope" ]

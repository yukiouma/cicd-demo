FROM golang:alpine as builder

WORKDIR /home

COPY ./ .

ENV CMD ./cmd/server

RUN CGO_ENABLED=0 go build -o server ./main.go


FROM alpine:3.18

WORKDIR /home

COPY --from=builder /home/server ./server

CMD ./server
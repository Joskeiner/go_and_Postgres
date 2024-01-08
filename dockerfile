FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod  go.sum ./

RUN go mod download

COPY .  .

#Build

RUN go build -o /main

RUN echo "appgo:x:65534:65534:appgo:/:" > /etc_passwd

FROM scratch

COPY --from=builder /etc_passwd /etc/passwd

COPY --chown=appgo:appgo --from=builder /main /main

USER appgo

EXPOSE 8080

CMD ["main"]

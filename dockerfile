FROM golang:1.20

WORKDIR /app

COPY .  .
#COPY go.mod .

#COPY go.sum .

#COPY *.go  .

RUN go mod download

#Build

RUN go build -o main .

RUN chmod +x main

RUN mv main /usr/local/bin/
#RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8080



CMD ["main"]

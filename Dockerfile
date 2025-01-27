FROM golang:1.22

WORKDIR /learn-go
COPY go.mod go.sum /learn-go/
RUN go mod download

ADD . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/learn-go ./cmd/learn-go

EXPOSE 8888

# Start the application
ENTRYPOINT ["learn-go"]

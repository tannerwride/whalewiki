FROM golang:1.18
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /whalewiki
EXPOSE 4000
CMD ["/whalewiki"]



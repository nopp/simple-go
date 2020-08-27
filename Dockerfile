FROM golang:1.14
COPY . .
RUN unset GOPATH \
    && go build -o main .

ENTRYPOINT ["./main"]

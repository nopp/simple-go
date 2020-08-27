FROM golang:1.14
COPY . .
RUN go get -v -t . \
    && go build -o main .

ENTRYPOINT ["./main"]

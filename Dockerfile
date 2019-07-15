FROM golang:alpine
RUN mkdir /app
ADD . /go/src/app
WORKDIR /go/src/app

RUN apk add git mercurial

RUN go get -d -v ./...

#RUN apk del git mercurial

# Install the package
RUN go install -v ./...

RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser



EXPOSE 8000
CMD ["./main"]


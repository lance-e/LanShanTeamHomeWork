FROM golang:1.21 

ENV GOPROXY=https://goproxy.io,direct
ENV GOOS=darwin

WORKDIR /go/src/homework10

#ADD ./go.mod .
#ADD  ./go.sum .
COPY . .
RUN go mod download \
    && go build -o helloTest main.go

#FROM scratch

#WORKDIR /root/
#COPY ./helloTest .

CMD [ "helloTest" ]

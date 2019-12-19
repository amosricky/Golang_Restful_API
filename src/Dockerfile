FROM golang:latest

ENV GOPROXY https://proxy.golang.org,direct
WORKDIR $GOPATH/src/Golang_Restful_API
ADD .. $GOPATH/src/Golang_Restful_API
RUN cd $GOPATH/src/Golang_Restful_API
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./Golang_Restful_API"]

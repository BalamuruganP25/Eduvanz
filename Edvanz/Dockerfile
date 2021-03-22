FROM golang:1.15
RUN apt-get update
RUN apt-get install pkg-config gcc g++ -y 
WORKDIR /go/src/app
RUN echo $GOPATH
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
#WORKDIR $GOPATH/src
#RUN echo $GOPATH
CMD ["go", "run", "./src/main.go"]
#CMD ["app"]
EXPOSE 8084
FROM golang:1.19.8-alpine
RUN mkdir /services
ADD . /services/
WORKDIR /services/cmd
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]
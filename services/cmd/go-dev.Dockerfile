FROM golang:1.16-alpine
RUN mkdir /services
ADD . /services/
WORKDIR /services/cmd
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]
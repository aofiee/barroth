FROM golang:1.22.0-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]
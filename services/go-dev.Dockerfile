FROM golang:1.19.5-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]
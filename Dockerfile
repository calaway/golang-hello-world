FROM golang:1.11
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN go get "github.com/go-martini/martini"
RUN go build -o /usr/src/app/example .
EXPOSE 9000
CMD /usr/src/app/example

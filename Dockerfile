FROM golang:1.17-alpine
RUN apk add build-base

WORKDIR /chabokan

COPY go.mod ./
COPY go.sum ./

COPY . .

CMD ["/bin/sh" ,"-c" ,"go mod download"]

RUN go build -o ./chabokan

EXPOSE 443
EXPOSE 444


#CMD tail -f /dev/null

ENTRYPOINT [ "./chabokan" ]
FROM golang:1.14-alpine3.11

WORKDIR /app

CMD ["air"] 

EXPOSE 80

RUN apk add --no-cache git && \
  go get -u github.com/cosmtrek/air

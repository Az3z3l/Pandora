FROM golang:1.16.6

ENV APP_NAME pandora

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

COPY . .

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE 80
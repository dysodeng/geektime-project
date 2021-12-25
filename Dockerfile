# build stage
FROM golang:1.16

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
 && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /app

ADD ./go.mod /app
ADD ./go.sum /app

RUN export GOPROXY=https://goproxy.cn && go mod download

ADD . /app

RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest

COPY --from=0 /app/app /app/app
COPY --from=0 /app/var /app/var
COPY --from=0 /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=0 /etc/localtime /etc/localtime
COPY --from=0 /etc/timezone /etc/timezone

WORKDIR /app

RUN chmod -R a+w /app/var

EXPOSE 9000
EXPOSE 8080

CMD /app/app

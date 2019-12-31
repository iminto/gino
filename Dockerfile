FROM golang
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY . .
RUN go build -mod=vendor -o gino
EXPOSE 8090

ENTRYPOINT ["./gino"]

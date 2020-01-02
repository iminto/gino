FROM golang
ENV TZ=Asia/Shanghai
ENV GOPROXY https://goproxy.cn,direct
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /app
COPY . /app
RUN go build -mod=vendor -o gino
# windows not support
#EXPOSE 8090
#COPY run.sh run.sh
#RUN chmod a+x run.sh
#CMD ["/run.sh"]
ENTRYPOINT ["./gino"]

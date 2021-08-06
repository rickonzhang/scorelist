#FROM golang
#
#WORKDIR /app
#COPY . .
##设置环境变量，开启go module和设置下载代理
#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
##会在当前目录生成一个go.mod文件用于包管理
##RUN go mod init
##增加缺失的包，移除没用的包
#RUN go mod tidy
#RUN go build
#EXPOSE 9000:9000
#CMD ["sh", "/app/run.sh"]
FROM scratch
# 最小镜像，可以有效减少 image 文件大小

WORKDIR /app

# 将编译之后的可执行文件，数据库文件、配置文件 拷贝到 image 中
COPY ./scorelist .
COPY ./conf ./conf

EXPOSE 9000

# 启动服务
CMD ["sh", "/app/run.sh"]
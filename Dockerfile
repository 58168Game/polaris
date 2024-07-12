FROM instrumentisto/flutter:3 AS flutter
WORKDIR /app
COPY ./ui/pubspec.yaml ./ui/pubspec.lock ./
RUN flutter pub get
COPY ./ui/ ./
RUN flutter build web 


# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.20 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY  go.mod .
COPY  go.sum .
RUN go mod download

COPY . .

COPY --from=flutter /app/build/web ./ui/build/web/
# 指定OS等，并go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o polaris ./cmd/ 

FROM debian:12

WORKDIR /app
RUN apt-get update && apt-get -y install ca-certificates

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/polaris .

EXPOSE 8080

ENTRYPOINT ["./polaris"]
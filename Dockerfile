#
#使用golang1.8.1作为基础镜像,
# 这里使用公司仓库镜像,若不存在会自动拉取
FROM dev-registry.zhonganonline.com:5000/exchange/golang:1.8.1

# 设置author
MAINTAINER dylenfu

# 统一工作目录
WORKDIR /alidata1/admin/nextex

# 从环境中拷贝目录到容器,保证数据一致性
ADD . /alidata1/admin/nextex

# 统一rest端口
EXPOSE 9090

# 执行命令传递参数
ENTRYPOINT ["./server"]

CMD []
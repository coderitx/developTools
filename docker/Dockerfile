FROM centos:latest as build

WORKDIR /workdir/data/

COPY ./xxx /workdir/data


FROM alpine:latest
WORKDIR=/workdir/data/
COPY --from=build /xxx/xxx /workdir/data/
EXPOSE port

CMD ["/bin/sh","-c","ls"]

# COPY 的路径是从context_path 开始计算的
# docker build -t ${image_name}:${image_tag} -f Dockerfile ${context_path}

FROM BASE_IMAGE
LABEL maintainer="<cwzdzg@foxmail.com>"

WORKDIR /opt/miniblog

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/miniblog

COPY miniblog /opt/miniblog/bin/

ENTRYPOINT ["/opt/miniblog/bin/miniblog"]
CMD ["-c", "/etc/miniblog/miniblog.yaml"]
FROM alpine:3.5

ENV kubectlDownload https://storage.googleapis.com/kubernetes-release/release/v1.8.0/bin/linux/amd64/kubectl

RUN curl -L -o /usr/local/bin/kubectl \
        ${kubectlDownload}; \
    chmod +x /usr/local/bin/kubectl

COPY assets/  /opt/resource/
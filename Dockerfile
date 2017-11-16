FROM alpine:3.5
MAINTAINER Sharor
ENV kubectlVersion v1.8.0
RUN apk add --no-cache curl && curl -L -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/$(kubectlVersion)/bin/linux/amd64/kubectl && \
  apk del curl && chmod +x usr/local/bin/kubectl
COPY .  /opt/resource/

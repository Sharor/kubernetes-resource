FROM gliderlabs/alpine:3.6
LABEL Maintainer=Sharor,groenborg,naesheim
RUN apk-install python
COPY check/check.py in/in.py out/out.py /opt/resource/
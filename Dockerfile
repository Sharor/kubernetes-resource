FROM gliderlabs/alpine:3.6
LABEL Maintainer=Sharor,groenborg,naesheim
COPY requirements.txt .
RUN apk-install python3
RUN apk --update add --virtual build-dependencies py-pip \
  && pip install -r requirements.txt \
  && apk del build-dependencies

COPY check/check.py in/in.py out/out.py /opt/resource/  
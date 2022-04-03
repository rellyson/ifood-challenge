## production base image (targeted as base)
FROM python:3.10.4-alpine3.14 as base

WORKDIR /opt/ifood/slack-notifier-worker

RUN apk add --no-cache build-base

## download requirements
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["python", "./main.py"]

#############################################

## development image (targeted as develop)
FROM python:3.10.4-alpine3.14 as develop

WORKDIR /opt/ifood/slack-notifier-worker

RUN apk add --no-cache autoconf git build-base \
    automake libtool gettext gettext-dev make g++ texinfo curl

RUN pip install --upgrade pip

# Setup fswatch
RUN cd /root && wget https://github.com/emcrisostomo/fswatch/releases/download/1.16.0/fswatch-1.16.0.tar.gz && \
    tar -xf fswatch-1.16.0.tar.gz && rm fswatch-1.16.0.tar.gz
RUN cd /root/fswatch-1.16.0 && ./configure && make && make install
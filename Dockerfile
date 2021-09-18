FROM ubuntu:latest

LABEL maintainer "user"

ENV USER user
ENV HOME /home/${USER}
ENV SHELL /bin/bash

RUN useradd -m ${USER}
RUN gpasswd -a ${USER} sudo
RUN echo "${USER}:work" | chpasswd
SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get install -y
RUN apt-get install -y git sudo curl gcc sqlite3 make

RUN curl -O https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin
RUN go version

WORKDIR /home/${USER}

COPY . .

RUN go get -v github.com/rubenv/sql-migrate/...
RUN make create-db

CMD ["make", "run"]

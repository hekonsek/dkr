FROM ubuntu AS docker
RUN apt update
RUN apt install -y curl
RUN apt install -y apt-transport-https ca-certificates curl software-properties-common
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
RUN apt update
RUN apt install -y docker-ce

FROM ubuntu AS go
RUN apt update
RUN apt install -y wget
RUN wget https://golang.org/dl/go1.14.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.7.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/root/go/bin/:${PATH}"

FROM docker as docker-go
COPY --from=go /usr/local/go /usr/local/go
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/root/go/bin/:${PATH}"
FROM ubuntu

RUN apt update
RUN apt install -y wget unzip

RUN wget https://releases.hashicorp.com/packer/1.6.0/packer_1.6.0_linux_amd64.zip
RUN unzip packer_1.6.0_linux_amd64.zip
RUN cp packer /usr/bin/packer

ENTRYPOINT ["/usr/bin/packer"]
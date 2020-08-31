FROM ubuntu

RUN apt update
RUN apt install -y curl gnupg software-properties-common
RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add -
RUN apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
RUN apt-get update
RUN apt-get install -y terraform

ENTRYPOINT ["/usr/bin/terraform"]
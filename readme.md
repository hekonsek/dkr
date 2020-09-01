# DKR - dockerize your shell commands

DKR (pronounced *dockerizer*) is a simple toolkit to help you dockerize your shell commands. Under the hood 
DKR handles all the heavy lifting necessary to run commands in a container environment:
- passing environment variables
- bridging host's network
- mounting host file system
- mounting current working directory 
- creating symlink/alias like proxy to commands 

# Usage

In order to install command from DKR catalog use the following command:

```
sudo dkr cmd install packer
```

For example the command above installs HashiCorp Packer:

```
$ packer version
Packer v1.6.1
```

## Installation

The easiest way to install DKR is via DockerHub distributed image:

```
docker create --name dkr hekonsek/dkr
sudo docker cp dkr:/bin/dkr /usr/bin/
sudo docker cp dkr:/bin/dkr-proxy /usr/bin/
```

 ## License
 
 This project is distributed under [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0.html).
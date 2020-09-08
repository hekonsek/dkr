# DKR - dockerize your shell commands

DKR (pronounced *dockerizer*) is a simple toolkit to help you dockerize your shell commands. Under the hood 
DKR handles all the heavy lifting necessary to run commands in a container environment:
- passing environment variables
- bridging host's network
- mounting host file system
- mounting current working directory
- mounting host user home and setting up $HOME variable 
- creating Bash aliases to commands 

# Usage

In order to install command from DKR catalog use the following command:

```
$ sudo dkr cmd install packer
```

For example the command above installs HashiCorp Packer. You can check if installation was successful
by executing the installed command:

```
$ packer version
Packer v1.6.1
```

## Installation

The easiest way to install DKR is via DockerHub distributed image:

```
docker create --name dkr hekonsek/dkr
sudo docker cp dkr:/bin/dkr /usr/bin/
```

 ## License
 
 This project is distributed under [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0.html).
# DKR - dockerize your shell commands

DKR (pronounced *dockerizer*) is a simple toolkit to help you dockerize your shell commands. Under the hood 
DKR handles all the heavy lifting necessary to run commands in a container environment:
- passing environment variables
- bridging host's network
- mounting host file system
- mounting current working directory
- mounting host user home and setting up $HOME variable 
- creating Bash "proxy" scripts for installed commands 
- ensuring that installed commands are included in $PATH
- telling Docker to use current user instead of root

## Usage

In order to install command from DKR catalog use the following command:

```
$ dkr cmd install terraform
Command terraform installed.
```

For example the command above installs HashiCorp Terraform. You can check if installation was successful
by executing the installed command:

```
$ terraform version
Terraform v0.13.2
```

## Installation

The easiest way to install DKR is via DockerHub distributed image:

```
docker create --name dkr hekonsek/dkr
sudo docker cp dkr:/bin/dkr /usr/bin/
```

## License
 
This project is distributed under [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0.html).
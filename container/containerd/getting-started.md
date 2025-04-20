# Getting started with containerd

Reference: [getting started with containerd](https://github.com/containerd/containerd/blob/main/docs/getting-started.md)

The platform I use at the moment of documenting this note is WSL2 (Ubuntu 24.04.1 LTS):

```console
$ cat /etc/os-release
PRETTY_NAME="Ubuntu 24.04.1 LTS"
NAME="Ubuntu"
VERSION_ID="24.04"
VERSION="24.04.1 LTS (Noble Numbat)"
VERSION_CODENAME=noble
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=noble
LOGO=ubuntu-logo
```

To get into the world of `containerd`, I follow the guide mentioned in the above reference:

- Installing `containerd` [v2.0.3](https://github.com/containerd/containerd/releases/tag/v2.0.3)
- Installing `runc` [v1.3.0-rc.1](https://github.com/opencontainers/runc/releases/tag/v1.3.0-rc.1)
- Installing CNI plugins [v1.6.2](https://github.com/containernetworking/plugins/releases/tag/v1.6.2)

After having these tools installed on the system, I start up the containerd service in the background:

```console
$ sudo nohup containerd &
```

And now I can see the containerd server is running:

```console
$ sudo ctr version
[sudo] password for mklan:
Client:
  Version:  v2.0.3
  Revision: 06b99ca80cdbfbc6cc8bd567021738c9af2b36ce
  Go version: go1.23.6

Server:
  Version:  v2.0.3
  Revision: 06b99ca80cdbfbc6cc8bd567021738c9af2b36ce
  UUID: 823c67db-887e-4728-a446-1ce02aed54d0
```

Before interacting with the `containerd` server, I pull the `golang:alpine` image from Docker Hub using `docker`:

```console
$ docker pull golang:alpine
alpine: Pulling from library/golang
31e352740f53: Pull complete
7f9bcf943fa5: Pull complete
ee52342d2eff: Pull complete
5107867bbeaf: Pull complete
Digest: sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c
Status: Downloaded newer image for golang:alpine
docker.io/library/golang:alpine
```

This time, let's try to pull the same image using the `ctr` command line tool.

```console
$ sudo ctr images ls
REF TYPE DIGEST SIZE PLATFORMS LABELS
$ sudo ctr image pull docker.io/library/golang:alpine
docker.io/library/golang:alpine:                                                  resolved       |++++++++++++++++++++++++++++++++++++++|
index-sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c:    done           |++++++++++++++++++++++++++++++++++++++|
manifest-sha256:e7cc33118f807c67d9f2dfc811cc2cc8b79b3687d0b4ac891dd59bb2a5e4a8d3: done           |++++++++++++++++++++++++++++++++++++++|
layer-sha256:5107867bbeaf1ff53c00102bd5046a59c1daf6eb36caf190abafa4cc43aea99b:    done           |++++++++++++++++++++++++++++++++++++++|
layer-sha256:7f9bcf943fa5571df036dca6da19434d38edf546ef8bb04ddbc803634cc9a3b8:    done           |++++++++++++++++++++++++++++++++++++++|
layer-sha256:ee52342d2eff6551935616ac36a72d1c8a116bcf454e22ce50061be1c14bc5cf:    done           |++++++++++++++++++++++++++++++++++++++|
layer-sha256:31e352740f534f9ad170f75378a84fe453d6156e40700b882d737a8f4a6988a3:    done           |++++++++++++++++++++++++++++++++++++++|
config-sha256:9e57a8e8195932a90847f4c081ca57cce48c32cccc98c3dc17ef99f7cb823855:   done           |++++++++++++++++++++++++++++++++++++++|
elapsed: 26.3s                                                                    total:  99.5 M (3.8 MiB/s)

unpacking linux/amd64 sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c...
done: 2.560423s
$ sudo ctr images ls
REF                             TYPE                                                      DIGEST                                                                  SIZE     PLATFORMS                                                                                LABELS
docker.io/library/golang:alpine application/vnd.docker.distribution.manifest.list.v2+json sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c 99.8 MiB linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64/v8,linux/ppc64le,linux/s390x -
```

And now we are able to perform normal routines of manipulating containers as we do via `docker`:

```console
$ sudo ctr container ls
CONTAINER    IMAGE    RUNTIME
$ sudo ctr run -d docker.io/library/golang:alpine golang
$ sudo ctr container ls
CONTAINER    IMAGE                              RUNTIME
golang       docker.io/library/golang:alpine    io.containerd.runc.v2
$ sudo ctr task ls
TASK      PID      STATUS
golang    26604    RUNNING
$ sudo ctr task exec -t --exec-id 1234 golang /bin/sh
/go # go version
go version go1.20.5 linux/amd64
/go # exit
$ sudo ctr container rm golang
ERRO[0000] failed to delete container "golang"           error="cannot delete a non stopped container: {running 0 0001-01-01 00:00:00 +0000 UTC}"
ctr: cannot delete a non stopped container: {running 0 0001-01-01 00:00:00 +0000 UTC}
$ sudo ctr task kill -s SIGKILL golang
$ sudo ctr task ls
TASK      PID      STATUS
golang    26604    STOPPED
$ sudo ctr container rm golang
$ sudo ctr container ls
CONTAINER    IMAGE    RUNTIME
```

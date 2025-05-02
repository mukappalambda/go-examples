# pubsub

Instructions:

```console
$ docker run -dt --name redis -p 6379:6379 redis:8.0-rc1-alpine
5c1690d308e2039186839dced5c543cdb52c6a53661a070996bf0893ce6f640f
$ docker ps --filter name=redis
CONTAINER ID   IMAGE                  COMMAND                  CREATED         STATUS        PORTS                    NAMES
5c1690d308e2   redis:8.0-rc1-alpine   "docker-entrypoint.s…"   2 seconds ago   Up 1 second   0.0.0.0:6379->6379/tcp   redis
$ go build
$ ./pubsub
[Time]: 2025-05-02 16:24:17.309807555 +0800 CST m=+0.004429045 [Payload]: go
[Time]: 2025-05-02 16:24:17.411347087 +0800 CST m=+0.105968602 [Payload]: redis
[Time]: 2025-05-02 16:24:17.513303427 +0800 CST m=+0.207924935 [Payload]: go-redis
^C
$ docker rm -f redis
redis
$
```

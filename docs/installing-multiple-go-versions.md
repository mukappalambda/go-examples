# Installing multiple Go versions

Ref: [Managing Go installations](https://go.dev/doc/manage-install)

Check the current Go version:

```
$ go version
go version go1.19.1 linux/amd64
```

Check the environment variable `GOPATH` is set, and `GOPATH` is included in your `PATH`:

```
$ go env GOPATH # or alternatively, run echo $GOPATH 
```

If not, add the following two lines into the `~/.bashrc` file and source that file:

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Now, if you want to download a specific Go version, say `go1.19.3`, run:

```
$ go install golang.org/dl/go1.19.3@latest
$ go1.19.3 download
```

After executing the first command, you should see `go1.19.3` under the `GOPATH` folder by checking with:

```
$ find $HOME/go/bin -name go1.19.3
```

What are the available Go versions? See the [download page](https://go.dev/dl/).

Finally, you should see the Go version you installed:

```
$ go1.19.3 version
go version go1.19.3 linux/amd64
```


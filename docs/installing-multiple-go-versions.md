# Installing multiple Go versions

As of 2026, I tend to manage multiple Go versions using [mise](https://github.com/jdx/mise):

```bash
# have mise installed on the machine (see the mise's Quickstart section in README.md)
curl https://mise.run | sh
# search the go tool in the mise registry
mise search go | grep 'go lang'
# list go versions available for install
mise ls-remote go
# install the go tool locally
mise use go
# install the go tool globally
mise use go --global
# install the go tool of a given version
mise use go@1.26.6
# check installation
mise ls go
mise where go
go version
mise exec go -- go version
```

---

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

# Find project

A simple Go app to jump-to-a-project inside your `$GOPATH`. I use it as an alias.

```bash
# Go project jump
function gs() {
   cd $(find-project $1)
}
```

So then if you have this folder... `$GOPATH/src/github.com/project-a/package-b/subpackage-c`, you can execute...

```bash
$ gs package-b && pwd
$GOPATH/src/github.com/project-a/package-b

$ gs project-a && pwd
$GOPATH/src/github.com/project-a

$ gs subpackage-c
$GOPATH/src/github.com/project-a/package-b/subpackage-c
```

### Install

Simply issue a `go get` command, as follows.

```bash
go get -u github.com/patrickdappollonio/find-project
```

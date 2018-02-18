# Find project [![Build Status](https://travis-ci.org/patrickdappollonio/find-project.svg?branch=master)](https://travis-ci.org/patrickdappollonio/find-project)

A simple Go app to jump-to-a-project inside your `$GOPATH`. I use it as a command / alias, so it can be connected to the Current Working Directory (CWD). Otherwise it'll just return the path to the folder with the given name inside `$GOPATH`. I use mine with the `gs` alias (as in "go switch"):

```bash
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

### How it works

`find-project` traverses your `$GOPATH` looking for a folder with a given name in a depth-first way, so for short paths (or even better, _paths-not-too-deep_) it'll use less resources as well as run faster. It'll start by listing all directories inside `$GOPATH/src` and checking at the same time if the folder you're looking for exists there. If it doesn't, it'll go one level deep *on each folder* inside `$GOPATH/src`, so most likely, it'll go inside `$GOPATH/src/github.com` and `$GOPATH/src/golang.org` and `$GOPATH/src/bitbucket.org` and check for the given folder there. If it still doesn't exist, it'll go one level deep inside each one of those folders. Rinse and repeat until we find yours.

As an alternative, you can also have a pure-bash solution:

```bash
function gs() {
  cd $(find $GOPATH/src -type d -name $1 | sed 1q)
}
```

The downsides to this function are:

1. It may take a while to run, whereas `find-project` should run much quickly.
2. It'll prefer longer paths, while `find-project` won't.
3. It'll include dot-folders and `vendor/` folders, whereas `find-project` will silently omit them.

While point (3) is really good, point (2) is IMHO the most important. Say you have two paths:

```
$GOPATH/src/github.com/patrickdappollonio
$GOPATH/src/github.com/abc-corporation/a-simple-name/internal/users/patrickdappollonio
```
Then the issue here is depth versus text sorting. `find` works by recursively listing content, and since `abc-corporation` in a sorted list will be first, then `patrickdappollonio`, then `abc-corporation` will be scanned first. This is almost never the intention I found myself at while switching folders. I always wanted the folder close to `$GOPATH`, so this tool works.

For point (3) consider a copy of a given folder inside `$GOPATH/src/github.com/user/project/.git` which may be scanned, or even a vendored dependency which will be checked in first than the original source. You probably get my point.

### Install

Simply issue a `go get` command, as follows.

```bash
go get -u github.com/patrickdappollonio/find-project
```

Then, modify your bash file (`.bashrc`, `.bash_profile` or whatever you use, it may even work for oh-my-zsh) and copy-paste the first bash script in this README. You can change the function name `gs()` to any other name. That'll be the name of the command you will call from your bash terminal.

# Find project

A simple Go app to jump-to-a-project inside your `$GOPATH`. I use it as a command / alias.

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

`find-project` traverses your `$GOPATH` and finds the shortest way to the folder, so for short paths (or even better, _paths-not-too-deep_) it'll use less resources as well as run faster. The pure-bash alternative to this is:

```bash
function gs() {
    cd $(find $GOPATH/src -type d -name $1 | sed 1q)
}
```

The downsides to this function are:

1. It may take a while to run, whereas `find-project` should run much quickly.
2. It'll prefer longer paths, while `find-project` won't.
3. It'll include dot-folders and `vendor/` folders, whereas `find-project` will silently omit them.

### Install

Simply issue a `go get` command, as follows.

```bash
go get -u github.com/patrickdappollonio/find-project
```

Then, modify your bash file (`.bashrc`, `.bash_profile` or whatever you use, it may even work for oh-my-zsh) and copy-paste the first bash script in this README. You can change the function name `gs()` to any other name. That'll be the name of the command you will call from your bash terminal.

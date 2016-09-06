# Find project

A simple Go app to jump-to-a-project inside your `$GOPATH`. I use it as an alias.

```bash
# Go project jump
function gs() {
   cd $(find-project $1)
}
```
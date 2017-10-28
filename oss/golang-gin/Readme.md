
# Golang Gin

Hello World using the Gin framework.

## Install dependencies

This will install the required dependencies and uses the [Go Dependency Management Tool](https://github.com/golang/dep),
which you may need to install.

```
$ dep ensure
```

## Deploy

```
$ up
```

## Notes

Typically in .upignore you'd list specific files or directories to omit, however in cases such as Golang where you deploy a single binary, it may be simpler to blacklist all files, and whitelist the `./server` binary.

```
*
!server
```

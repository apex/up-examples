
# Golang Gin

Hello World using the Gin framework.

## Install dependencies

This will install the required dependencies and uses the [Go Dependency Management Tool](https://github.com/golang/dep),
which you may need to install.

```
$ deps ensure
```

## Deploy

```
$ up
```

## Notes

In this case .upignore specifies files to omit, however in cases such as Golang where
typically only a binary is uploaded, you could whitelist everything and blacklist only `server`,
the binary itself.

```
*
!server
```

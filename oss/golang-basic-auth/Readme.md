
# Golang Basic Auth

Basic authentication example.

## Deploy

```
$ up
```

## Notes

Use "tobi" for the username and "ferret" for the password.

This example currently fails because AWS is strange and decides to remap header fields. If enough of us bother them about this, hopefully they will fix this issue.

```
x-amzn-remapped-www-authenticate:Basic realm="Ferret Land"
```

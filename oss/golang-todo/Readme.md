
# Golang Todo

Golang traditional server-side todo list.

## Deploy

```
$ up
```

## Notes

Usage:

```
$ curl -d 'command=ls -la' `up url`
```

Or paste this in your terminal:

```
lambda() {
  curl -d command="$*" `up url`
}
```

And run commands:

```
$ lambda ls -la /
```

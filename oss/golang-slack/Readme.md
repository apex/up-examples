
# Golang Slack

Slack command for timing a website response.

## Deploy

Deploy the app:

```
$ up
```

Copy the URL and register it on Slack as a command.

```
$ up url -c
Copied to the clipboard!
```

## Notes

This example is used to create a command such as `/time https://apex.sh` and respond with a breakdown of timing information for the request, such as DNS, connection, TLS handshake, wait and download times.

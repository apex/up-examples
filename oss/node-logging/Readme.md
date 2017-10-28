
# Node Logging

Logging example with plain-text and JSON logs.

## Deploy

```
$ up
```

## Usage

Tail the logs in one tab:

```
$ up logs -f
```

Hit the plain-text route a few times and you should see the logs.

```
$ curl `up url`/text
```

Hit the json route to see how the JSON logs are translated to the internal
format, allowing you to specify fields etc.

```
$ curl `up url`/json
```

The following end-point outputs a stack trace:

```
$ curl `up url`/stack
```

Or alternatively crash and output an error in the uncaughtException handler.

```
$ curl `up url`/crash
```

## Notes

Up supports indented logs, useful for cases such
as stack traces, so they're considered a single
log message and not one per line.

JSON logs need at least these fields:

- `level` (string) – severity level
- `message` (string) – log message

Optionally:

- `fields` (object) – key/value pairs for contextual information

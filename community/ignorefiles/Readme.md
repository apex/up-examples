
# Ignorefiles

Example of ignore file usage. Precedence is:

- .gitignore
- .npmignore
- .upignore

## Deploy

```
$ up
```

Run with `-v` to view which files are added to the zip:

```
$ up -v
```

## Notes

The following files must be deployed, as they are used for the reverse proxy:

- up.json
- main

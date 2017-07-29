
# Node Next

Next.js application.

## Setup

```
$ yarn
```

## Deploy

```
$ up
```

## Notes

The `.upignore` file contains `!.next` to include the `./.next` build directory. This is necessary because `.gitignore`'d files and dotfiles (or dirs) are ignored by default.

Next.js requires more `.lambda.memory` to improve serving speed, as AWS scales CPU alongside the memory setting. 128mb is not recommended as Next will take ~400ms or more to render.

Next.js is also quite slow to initialize, so a cold function may take ~3s.

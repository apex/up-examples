
# Environment Static

Environment variable availability during hooks.

## Notes

This example illustrates how environment variables are available at build time,
so you may use `UP_STAGE`, `NODE_ENV` which is implied, or any custom env vars
mapped to stages via `up env` to alter your builds on these values.

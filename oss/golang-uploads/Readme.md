
# Golang Multipart Form

Simple multi-part form processing.

## Deploy

```
$ up
```

## Notes

Lambda & API Gateway impose restrictions, so uploads of > 10MiB will error, view [API Gateway Limits](http://docs.aws.amazon.com/apigateway/latest/developerguide/limits.html) for details.

For large files you should upload to S3 directly from the client using signed requests, and process in Lambda if necessary.

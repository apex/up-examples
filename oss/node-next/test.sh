up

curl -s `up url`
# contains: 1.png

curl -s -D- `up url`/static/1.png
# contains: image/png

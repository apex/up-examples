up

curl -s `up url`
# contains: <title>Hello</title>

curl -s -D- `up url`
# contains: Content-Type: text/html

curl -s -D- `up url`/style.css
# contains: Content-Type: text/css

up

curl -s `up url`
# contains: Static site example

curl -s -D- `up url`/blog
# contains: Location: https://blog.apex.sh

curl -s -D- `up url`/blog/foo
# not contains: Location: https://blog.apex.sh

curl -s -D- `up url`/whatever/something/here
# contains: Static site example

curl -s -D- `up url`/store/starbucks/products/green-tea-latte
# contains: Location: /shop/starbucks/green-tea-latte

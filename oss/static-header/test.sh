up

curl -s `up url`
# contains: Static site example

curl -s -D- `up url`
# contains: X-Something: I am applied to everything

curl -s -D- `up url`/Readme.md
# contains: X-Something: I am applied to everything

curl -s -D- `up url`/Readme.md
# not contains: X-Something-Else

curl -s -D- `up url`/style.css
# contains: X-Something-Else

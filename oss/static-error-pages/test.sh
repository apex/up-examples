up

curl -s `up url`
# contains: Index page

curl -s `up url`/subpage/
# contains: Sub page

curl -s `up url`/missing
# contains: Not Found

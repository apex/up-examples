up

curl -s `up url`
# contains: Static site example

curl -s `up url`/style.css
# contains: Helvetica

curl -s `up url`/../super-secret.txt
# contains: 404 page not found

up

curl -s `up url`
# contains: Enter your name please

curl -s -d first=Tobi -d last=Ferret -D- `up url`
# contains: set-cookie: first=Tobi

curl -s -d first=Tobi -d last=Ferret -D- `up url`
# contains: set-cookie: last=Ferret

up

curl -s `up url`
# contains: <p>Fill out this amazing form.</p>

curl -D- -d 'name=Tobi&email=tobi@apex.sh' `up url`/submit
# contains: Set-Cookie: name=Tobi

curl -D- -d 'name=Tobi&email=tobi@apex.sh' `up url`/submit
# contains: Location: / 

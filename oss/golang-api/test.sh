up

curl -s `up url`
# equals: no pets

curl -d Tobi `up url`
# equals: welcome to the family Tobi!

curl -d Loki `up url`
# equals: welcome to the family Loki!

curl -d Jane `up url`
# equals: welcome to the family Jane!

curl -X DELETE `up url`/Tobi
# equals: removed Tobi!

curl -s `up url`
# contains: Jane

curl -s `up url`
# contains: Loki

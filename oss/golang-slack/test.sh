up

curl `up url`
# equals: url is required

curl -d 'command=/time&user_name=tj' `up url`
# equals: url is required

curl -d 'command=/time&user_name=tj&text=https://apex.sh' `up url`
# contains: Status 200

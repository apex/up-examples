up env set MSG Hello
up env set NAME Tobi
up env set -s staging NAME Loki
up env set -s staging NODE_ENV production
up env set -s production NAME Jane

up deploy production
up deploy staging
up

curl -s `up url`
# equals: Hello Tobi from development (NODE_ENV=development)

curl -s `up url staging`
# equals: Hello Loki from staging (NODE_ENV=production)

curl -s `up url production`
# equals: Hello Jane from production (NODE_ENV=production)

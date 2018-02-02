up env set MSG Hello
up env set NAME Tobi
up env set -s staging NAME Loki
up env set -s staging NODE_ENV production
up env set -s production NAME Jane

up
up production

curl -s `up url`
# equals: Hello Loki from staging (NODE_ENV=production)

curl -s `up url production`
# equals: Hello Jane from production (NODE_ENV=production)

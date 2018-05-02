up env set -s staging NAME Loki
up env set -s production NAME Jane

up
up production

curl -s `up url -s staging`
# contains: Hello Loki from staging

curl -s `up url -s production`
# contains: Hello Jane from production

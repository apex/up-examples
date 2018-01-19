up env set NAME Tobi
up env set -s staging NAME Loki
up env set -s production NAME Jane

up
up staging
up production

curl -s `up url`
# contains: Hello Tobi from development

curl -s `up url staging`
# contains: Hello Loki from staging

curl -s `up url production`
# contains: Hello Jane from production

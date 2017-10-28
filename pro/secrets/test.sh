up env set MSG Hello
up env set NAME Tobi
up env set -s staging NAME Loki
up env set -s production NAME Jane

up deploy production
up deploy staging
up

curl -s `up url`
# equals: Hello Tobi from development

curl -s `up url staging`
# equals: Hello Loki from staging

curl -s `up url production`
# equals: Hello Jane from production

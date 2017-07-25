
set -e

for plan in $(find . -type d -depth 1 | grep -v .git); do
  echo "==> testing $plan"
  for dir in $(find $plan -type d -depth 1); do
    echo "  --> testing $dir"
    up -C $dir config > /dev/null
  done
done

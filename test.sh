
set -e

for plan in $(find . -type d -depth 1 | grep -v .git); do
  echo "==> $plan"
  for dir in $(find $plan -type d -depth 1); do
    echo "  --> $dir"
    up -C $dir config 2>&1 > /dev/null
  done
done

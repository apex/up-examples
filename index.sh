
cd $1
for dir in $(find . -type d -depth 1 | grep -v .git); do
  desc=$(grep -E '^# ' $dir/Readme.md -A 2 | tail -n 1)
  name=$(echo $dir | sed 's|./||' | tr '-' ' ')
  echo "- [$name]($dir) – $desc"
done

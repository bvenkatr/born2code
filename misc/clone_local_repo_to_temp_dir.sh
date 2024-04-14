tempDir=$(mktemp -d)
git clone --branch $1 $PWD $tempDir
echo "branch $1 is cloned to $tempDir"


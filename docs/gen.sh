#
#	for file in files:
# 		if file.svg isn't in the repo || file.mmd hasn't been changed:
# 			mermaid.generate(file)
#

for file in "$@"; do
    if !(test -f lock-manager/gen/file}.svg) || !(git diff --exit-code --quiet lock-manager/src/file}.mmd)
    then
        npx -p @mermaid-js/mermaid-cli mmdc -i src/${file}.mmd -o lock-manager/gen/${file}.svg --cssFile lock-manager/src/sequence/style.mmd.css --configFile lock-manager/src/sequence/config.json
    fi
done

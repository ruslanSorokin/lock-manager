#
#	for file in files:
# 		if file.svg isn't in the repo || file.mmd hasn't been changed:
# 			mermaid.generate(file)
#

for file in "$@"; do
    if !(test -f gen/${file}.svg) || !(git diff --exit-code --quiet src/${file}.mmd)
    then
        npx -p @mermaid-js/mermaid-cli mmdc -i src/${file}.mmd -o gen/${file}.svg --cssFile src/sequence/style.mmd.css --configFile src/sequence/config.json
    fi
done

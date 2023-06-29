install-tools:
	npm install


mmd_src := lock-manager/sequence/lock lock-manager/sequence/unlock

.SILENT: generate-uml
generate-uml:
	./gen.sh $(mmd_src)

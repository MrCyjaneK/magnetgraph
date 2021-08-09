install:
	cp share/magnetgraph.service /etc/systemd/system/magnetgraph.service
	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/${BINNAME}
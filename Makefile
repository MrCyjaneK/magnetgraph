install:
	@if [ "${BINNAME}" = "magnetgraph" ]; then\
		cp share/magnetgraph.service /etc/systemd/system/magnetgraph.service;\
	fi
	@if [ "${BINNAME}" = "magnetgraph-desktop" ]; then\
		cp share/magnetgraph.png /usr/share/pixmaps/magnetgraph.png
		cp share/magnetgraph.desktop /usr/share/applications
	fi

	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/${BINNAME}
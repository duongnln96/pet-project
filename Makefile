proto-go-gen:
	echo "++++ buf generate ++++"
	rm -rf ./gen/go/* && buf generate && go mod tidy
	echo "++++ complete ++++"
.PHONY: proto-go-gen

clean:
	go clean

wire:
	cd internal/proxy && wire && cd - && \
	cd internal/product && wire && cd -
.PHONY: wire
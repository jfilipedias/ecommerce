build/catalog:
	go build -o ./catalog-api/bin/catalog-api ./catalog-api/cmd/catalog-api

run/catalog: build/catalog
	./catalog-api/bin/catalog-api

test/catalog:
	go test -v ./catalog-api/...

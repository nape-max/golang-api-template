.SILENT generate:
	echo "Remove previous generated code..."
	rm -rf ./internal/generated/schema

	echo "Create directory after delete..."
	mkdir -p ./internal/generated/schema

	echo "Run \"oapi-codegen\" to generate server from OpenAPI Schema..."
	oapi-codegen --config=.oapi-codegen.yml ./api/schema.yaml

	echo "Run generator to create API stubs for new endpoints..."
	go run ./cmd/generator

	echo "Done."


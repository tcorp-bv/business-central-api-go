
# Run the unit tests
test:
	go test ./...

# Run the integration tests (require environment variables BC_CLIENT_ID, BC_USERNAME, BC_PASSWORD and BC_TENANT_ID)
# Also requires the tenant to have a 'sandbox' business central environment with at least 1 company
test-integration:
	go test ./... --tags=integration


# Regenerate the openapi client
generate:
	rm -rf openapi
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3:unstable generate -i https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/contracts/bcoas1.0.yaml -l go -o /local/swagger
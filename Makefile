
# Run the unit tests
test:
	go test ./...

# Run the integration tests (require environment variables BC_CLIENT_ID, BC_USERNAME, BC_PASSWORD and BC_TENANT_ID)
# Also requires the tenant to have a 'sandbox' business central environment with at least 1 company
test-integration:
	go test ./... --tags=integration
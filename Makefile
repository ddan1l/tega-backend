diff:
	atlas migrate diff --env gorm --dev-url "docker://postgres/15/dev?search_path=public"
migrate:
	atlas migrate apply --env gorm

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
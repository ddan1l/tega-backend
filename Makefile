PARENT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/..)

BACKEND_DIR := $(PARENT_DIR)/tega-backend
COMMON_DIR := $(PARENT_DIR)/tega-common
FRONTEND_DIR := $(PARENT_DIR)/tega-frontend

# .PHONY
.PHONY: diff migrate visualize test bash generate-swagger generate-types restart-swag docs

### Atlas ###
diff:
	atlas migrate diff --env gorm --dev-url "docker://postgres/15/dev?search_path=public"

migrate:
	atlas migrate apply --env gorm

visualize:
	atlas schema inspect --env gorm --url env://src -w --dev-url "docker://postgres/15/dev?search_path=public"

### Tests ###
test:
	docker-compose exec backend go test -v -coverprofile=coverage.out ./...
	docker-compose exec backend go tool cover -html=coverage.out -o coverage.html

### Utils ###
bash:
	docker-compose exec backend bash

### Docs ###
generate-swagger:
	@echo "🔄 Generate Swagger docs..."
	swag init -g handlers/**/*handler.go --output $(BACKEND_DIR)/docs
	
	@# Check is other services exists
	@if [ -d "$(COMMON_DIR)/swagger" ] && [ -d "$(FRONTEND_DIR)" ]; then \
		echo "📁 Copy docs..."; \
		cp -r $(BACKEND_DIR)/docs $(COMMON_DIR)/swagger; \
		cp $(BACKEND_DIR)/docs/swagger.json $(FRONTEND_DIR)/swagger.json; \
		rm -rf $(BACKEND_DIR)/docs; \
		echo "📁 Generate ts files..."; \
		cd $(FRONTEND_DIR) && npm run generate-types-file; \
	else \
		echo "⚠️ Skipping copy: COMMON_DIR/swagger or FRONTEND_DIR does not exist"; \
	fi

restart-swagger-service:
	@if docker-compose ps -q swagger >/dev/null 2>&1; then \
		echo "🔄 Restarting swagger service..."; \
		docker-compose restart swagger; \
	else \
		echo "⚠️ Warning: swagger service does not exist - skipping restart"; \
	fi

docs: generate-swagger restart-swagger-service
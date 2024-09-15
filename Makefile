build:
	@go build -o out/main cmd/api/main.go


# Run the application
run:
	@go run cmd/api/main.go


docker-run:
	@if command -v docker > /dev/null; then \
	    docker compose -f docker/dev.docker-compose.yaml --env-file .env up -d; \
	else \
		echo "Docker is not installed on your machine. Exiting..."; \
		exit 1; \
    fi

# Down Database
docker-down:
	@if command -v docker > /dev/null; then \
	    docker compose -f docker/dev-docker-compose.yaml down; \
	else \
		echo "Docker is not installed on your machine. Exiting..."; \
		exit 1; \
	fi

# Live Reload
# ---
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

copy-env:
	cp .env.sample .env

# Setup the development environment
setup:
	@echo "--- Copying env ---"
	@make copy-env

	@echo "--- Setting up docker ---"
	@make docker-run

	@echo "\n"
	@echo "Setup complete!"
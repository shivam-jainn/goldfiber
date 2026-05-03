.PHONY: help dev lint test dev-infra dev-infra-down dev-infra-reset prod sqlc migrate

# Default target
.DEFAULT_GOAL := help

# Colors
GREEN  := \033[0;32m
CYAN   := \033[0;36m
YELLOW := \033[1;33m
RESET  := \033[0m

# =========================
# Help
# =========================

help: ## Show this help message
	@echo ""
	@echo "   $(YELLOW)██████╗  ██████╗ ██╗     ██████╗ $(RESET)███████╗██╗██████╗ ███████╗██████╗"
	@echo "  $(YELLOW)██╔════╝ ██╔═══██╗██║     ██╔══██╗$(RESET)██╔════╝██║██╔══██╗██╔════╝██╔══██╗"
	@echo "  $(YELLOW)██║  ███╗██║   ██║██║     ██║  ██║$(RESET)█████╗  ██║██████╔╝█████╗  ██████╔╝"
	@echo "  $(YELLOW)██║   ██║██║   ██║██║     ██║  ██║$(RESET)██╔══╝  ██║██╔══██╗██╔══╝  ██╔══██╗"
	@echo "  $(YELLOW)╚██████╔╝╚██████╔╝███████╗██████╔╝$(RESET)██║     ██║██████╔╝███████╗██║  ██║"
	@echo "   $(YELLOW)╚═════╝ ╚═════╝ ╚══════╝╚═════╝ $(RESET) ╚═╝     ╚═╝╚═════╝╚══════╝╚═╝  ╚═╝"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"; printf "\n$(YELLOW)Available Commands:\n$(RESET)"} \
	/^[a-zA-Z0-9_-]+:.*##/ { \
		printf "  $(GREEN)%-20s$(RESET) %s\n", $$1, $$2 \
	}' $(MAKEFILE_LIST)
	@echo ""

# =========================
# Development
# =========================

dev: ## Run development server with Air (hot reload)
	@echo "$(GREEN)Starting GoldFiber dev server with Air...$(RESET)"
	@air

lint: ## Run Go linters
	@echo "$(GREEN)Running linters...$(RESET)"
	@golangci-lint run ./...

test: ## Run unit tests
	@echo "$(GREEN)Running tests...$(RESET)"
	@go test ./...

# =========================
# Infrastructure (Dev)
# =========================

dev-infra: ## Start development infrastructure (Docker Compose)
	@echo "$(GREEN)Starting dev infra...$(RESET)"
	@docker compose -f infra/compose-dev.yml up -d

dev-infra-down: ## Stop development infrastructure
	@echo "$(GREEN)Stopping dev infra...$(RESET)"
	@docker compose -f infra/compose-dev.yml down

dev-infra-reset: ## Reset development infrastructure (wipe volumes)
	@echo "$(GREEN)Resetting dev infra...$(RESET)"
	@docker compose -f infra/compose-dev.yml down -v
	@docker compose -f infra/compose-dev.yml up -d

# =========================
# Infrastructure (Prod)
# =========================

prod: ## Start production infrastructure
	@echo "$(GREEN)Starting prod infra...$(RESET)"
	@docker compose -f infra/compose-prod.yml up -d

# =========================
# Database
# =========================

sqlc: ## Generate Go code from SQL (sqlc)
	@echo "$(GREEN)Generating SQLC code...$(RESET)"
	@sqlc generate

migrate: ## Run database migrations (goose)
	@echo "$(GREEN)Running migrations...$(RESET)"
	@goose up
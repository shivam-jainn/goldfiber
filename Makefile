.PHONY: help dev

# Default target

.DEFAULT_GOAL := help

# Colors (optional)

GREEN  := \033[0;32m
CYAN   := \033[0;36m
YELLOW := \033[1;33m
RESET  := \033[0m


help:
	@echo ""
	@echo "   $(YELLOW)██████╗  ██████╗ ██╗     ██████╗ $(RESET)███████╗██╗██████╗ ███████╗██████╗"
	@echo "  $(YELLOW)██╔════╝ ██╔═══██╗██║     ██╔══██╗$(RESET)██╔════╝██║██╔══██╗██╔════╝██╔══██╗"
	@echo "  $(YELLOW)██║  ███╗██║   ██║██║     ██║  ██║$(RESET)█████╗  ██║██████╔╝█████╗  ██████╔╝"
	@echo "  $(YELLOW)██║   ██║██║   ██║██║     ██║  ██║$(RESET)██╔══╝  ██║██╔══██╗██╔══╝  ██╔══██╗"
	@echo "  $(YELLOW)╚██████╔╝╚██████╔╝███████╗██████╔╝$(RESET)██║     ██║██████╔╝███████╗██║  ██║"
	@echo "   $(YELLOW)╚═════╝ ╚═════╝ ╚══════╝╚═════╝ $(RESET) ╚═╝     ╚═╝╚═════╝╚══════╝╚═╝  ╚═╝"
	@echo ""
	@echo "$(YELLOW)Available Commands:$(RESET)"
	@echo "  $(GREEN)make dev       Run development server with Air (hot reload)"
	@echo "  $(GREEN)make test      Run unit tests"
	@echo "  $(GREEN)make lint      Run Go linters"
	@echo ""

 dev:
	@echo "$(GREEN)Starting GoldFiber dev server with Air...$(RESET)"
	@air

lint:
	@echo "$(GREEN)Running linters...$(RESET)"
	@golangci-lint run ./...

 test:
	@echo "$(GREEN)Running tests...$(RESET)"
	@go test ./...

dev-infra:
	@echo "$(GREEN)Starting GoldFiber development infrastructure with Docker Compose...$(RESET)"
	@docker compose -f infra/compose-dev.yml up -d

dev-infra-down:
	@echo "$(GREEN)Stopping GoldFiber development infrastructure...$(RESET)"
	@docker compose -f infra/compose-dev.yml down

dev-infra-reset:
	@echo "$(GREEN)Resetting GoldFiber development infrastructure...$(RESET)"
	@docker compose -f infra/compose-dev.yml down -v
	@docker compose -f infra/compose-dev.yml up -d

prod:
	@echo "$(GREEN)Starting GoldFiber production infrastructure with Docker Compose...$(RESET)"
	@docker compose -f infra/compose-prod.yml up -d

.PHONY: ui build run dev

# Build Vue into ../dist (picked up by Go automatically)
ui:
	cd ui && npm run build

# Run Go server only
run:
	go run main.go

# Build Vue then compile Go binary
build: ui
	GOTOOLCHAIN=local go build -o server main.go

# Dev: run Go + Vite dev server concurrently
dev:
	@trap 'kill 0' SIGINT; \
	GOTOOLCHAIN=local go run main.go & \
	cd ui && npm run dev & \
	wait

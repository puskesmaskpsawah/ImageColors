OUTPUT_DIR=./bin/
CMD_DIR=./cmd/imagecolors/

# Development build
build-dev:
	go build -o $(OUTPUT_DIR) $(CMD_DIR)

# Production build with optimizations
build-prod:
	go build -ldflags="-s -w" -trimpath -o $(OUTPUT_DIR) $(CMD_DIR)

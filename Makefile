.PHONY: console

console:
	@echo "Building the console of library binary"
	go build -o bin/console_library main.go
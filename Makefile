GREEN := $(shell tput setaf 2)
RESET := $(shell tput sgr0)
build:
	go build
rebuildbuild: clean build 
clean:
	rm -rf myRotate ../test *.tar.gz
rebuild: clean build
test: build
	@echo "$(GREEN)Running tests...$(RESET)"
	@echo "$(GREEN)test1 flag$(RESET)"
	./myRotate -a test go.mod Makefile myRotate.go
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test2 $(RESET)"
	./myRotate go.mod Makefile myRotate.go
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)All tests completed!$(RESET)"
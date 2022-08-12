# color config
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m
COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

.PHONY: test
test:
	cd src; go test ./... -v | $(COLORIZE_PASS) | $(COLORIZE_FAIL)
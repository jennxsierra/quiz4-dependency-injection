## run-original: run the non-DI version
.PHONY: run-original
run-original:
	@echo "Running original (no DI) on :4000..."
	@go run ./cmd/original

## run-refactored: run the DI version
.PHONY: run-refactored
run-refactored:
	@echo "Running refactored (DI) on :4000..."
	@go run ./cmd/refactored
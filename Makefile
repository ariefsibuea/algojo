# Default variables
PLATFORM    ?= leetcode
LANG   		?= go
FILE        ?= challenge.go
TEMPLATE     = problems/template.go
PROBLEM     ?= SampleProblem

.PHONY: problems-create problems-exec-leetcode-go problems-exec-etc-go

# Usage: make problems-create PLATFORM=leetcode LANG=go FILE=506_subarray_sum_equals_k.go
problems-create:
	@mkdir -p problems/$(PLATFORM)/$(LANG)
	@if [ -f $(TEMPLATE) ]; then \
		cp $(TEMPLATE) problems/$(PLATFORM)/$(LANG)/$(FILE); \
		echo "✅ Created: problems/$(PLATFORM)/$(LANG)/$(FILE)"; \
	else \
		echo "❌ Error: Template $(TEMPLATE) not found."; \
	fi

# Usage: make problems-exec-leetcode-go PROBLEM=SampleProblem
problems-exec-leetcode-go:
	@echo "⏳ Executing: $(PROBLEM)\n"
	@go run problems/leetcode/go/*.go --solution $(PROBLEM)

# Usage: make cc-exec-etc-go PROBLEM=SampleProblem
problems-exec-etc-go:
	@echo "⏳ Executing: $(PROBLEM)\n"
	@go run problems/etc/go/*.go --solution $(PROBLEM)

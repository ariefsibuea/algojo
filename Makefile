# Default variables
PLATFORM    ?= leetcode
PROG_LANG   ?= go
FILE        ?= challenge.go
TEMPLATE     = problems/template.go
PROBLEM     ?= SampleProblem

.PHONY: coding-challenge cc-exec-leetcode-go cc-exec-etc-go

# Usage: make coding-challenge PLATFORM=leetcode PROG_LANG=go FILE=506_subarray_sum_equals_k.go
coding-challenge:
	@mkdir -p problems/$(PLATFORM)/$(PROG_LANG)
	@if [ -f $(TEMPLATE) ]; then \
		cp $(TEMPLATE) problems/$(PLATFORM)/$(PROG_LANG)/$(FILE); \
		echo "✅ Created: problems/$(PLATFORM)/$(PROG_LANG)/$(FILE)"; \
	else \
		echo "❌ Error: Template $(TEMPLATE) not found."; \
	fi

# Usage: make cc-exec-leetcode-go PROBLEM=SampleProblem
cc-exec-leetcode-go:
	@echo "⏳ Executing: $(PROBLEM)\n"
	@go run problems/leetcode/go/*.go --solution $(PROBLEM)

# Usage: make cc-exec-etc-go PROBLEM=SampleProblem
cc-exec-etc-go:
	@echo "⏳ Executing: $(PROBLEM)\n"
	@go run problems/etc/go/*.go --solution $(PROBLEM)

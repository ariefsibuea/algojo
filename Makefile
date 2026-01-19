# Default variables
PLATFORM    ?= leetcode
PROG_LANG   ?= go
FILE        ?= challenge.go
TEMPLATE     = problems/template.go

.PHONY: coding-challenge

# Usage: make coding-challenge PLATFORM=leetcode PROG_LANG=go FILE=506_subarray_sum_equals_k.go
coding-challenge:
	@mkdir -p problems/$(PLATFORM)/$(PROG_LANG)
	@if [ -f $(TEMPLATE) ]; then \
		cp $(TEMPLATE) problems/$(PLATFORM)/$(PROG_LANG)/$(FILE); \
		echo "✅ Created: problems/$(PLATFORM)/$(PROG_LANG)/$(FILE)"; \
	else \
		echo "❌ Error: Template $(TEMPLATE) not found."; \
	fi

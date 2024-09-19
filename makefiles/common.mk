apply-env:
	@if [ -z "$(IMAGE_TAG)" ]; then \
		echo "ERROR: IMAGE_TAG is not set."; \
		exit 1; \
	fi
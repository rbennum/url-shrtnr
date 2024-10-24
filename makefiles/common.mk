apply-env:
	@if [ -z "$(IMAGE_TAG)" ]; then \
		echo "ERROR: IMAGE_TAG is not set."; \
		exit 1; \
	fi

check-docker-host:
	@if [ "$(CHECK)" = "exists" ]; then \
		if [ -z "$(DOCKER_HOST)" ]; then \
			echo "DOCKER_HOST is not set, please set it first"; \
			exit 1; \
		else \
			echo "DOCKER_HOST exists"; \
		fi \
	elif [ "$(CHECK)" = "not-exists" ]; then \
		if [ -n "$(DOCKER_HOST)" ]; then \
			echo "Unset DOCKER_HOST first"; \
			exit 1; \
		else \
			echo "DOCKER_HOST is not set, continue"; \
		fi \
	else \
		echo "Invalid CHECK flag: use 'exists' or 'not-exists'"; \
		exit 1; \
	fi

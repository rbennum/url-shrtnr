web-build-local:
	@npm --prefix ./frontend run build-local

web-start-local:
	@npm --prefix ./frontend run start-local

web-build-dev:
	@npm --prefix ./frontend run build-dev

web-build:
	@npm --prefix ./frontend run build-prod
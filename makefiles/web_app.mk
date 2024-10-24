web-build-local:
	@npm --prefix ./frontend run build-local

web-start-local:
	@npm --prefix ./frontend run start-local

web-build-dev:
	@npm --prefix ./frontend run build-dev

web-build:
	@npm --prefix ./frontend run build-prod

web-docker-dev: check-docker-host web-build-dev
	@docker image build \
	-f ./frontend/Dockerfile.web \
	-t shrtnr_web:$(IMAGE_TAG) \
	./frontend

web-docker: web-build
	@docker build -f ./docker/Dockerfile.web \
	-t rbennum2329/shrtnr_web:$(IMAGE_TAG) .
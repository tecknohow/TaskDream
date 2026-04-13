.PHONY: all build build-backend build-frontend run test lint docker-build clean

all: build

build-backend:
	CGO_ENABLED=1 go build -o taskdream .

build-frontend:
	cd frontend && pnpm install && pnpm build

build: build-frontend build-backend

run:
	go run . web

dev-backend:
	go run . web

dev-frontend:
	cd frontend && pnpm dev

test:
	go test ./... -v -cover

lint:
	golangci-lint run ./...
	cd frontend && pnpm lint

docker-build:
	docker build -t taskdream:latest .

docker-run:
	docker-compose up -d

clean:
	rm -f taskdream
	rm -rf frontend/dist
	rm -rf frontend/node_modules

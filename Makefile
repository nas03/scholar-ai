APP_NAME=web

.PHONY: backend frontend

backend:
	go run backend/cmd/$(APP_NAME)

frontend:
	bun --cwd frontend dev
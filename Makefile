
.PHONY: tailwind
tailwind:
	npx tailwindcss -i ./internal/view/css/input.css -o ./internal/view/css/output.css --watch

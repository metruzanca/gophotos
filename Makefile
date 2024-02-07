run:
	@pnpm tailwindcss -o static/styles.css --minify
	@templ generate
	@go run src/main.go

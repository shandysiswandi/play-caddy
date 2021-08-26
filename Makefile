build: clear
	@docker build -f Dockerfile --tag playcaddy .

run: clear
	@docker run --restart=always -d --name playcaddy -p 8080:8080 playcaddy:latest
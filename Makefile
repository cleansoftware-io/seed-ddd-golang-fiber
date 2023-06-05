build:
	docker build -t fiber-seed:latest .
run:
	docker run -p 3000:3000 -e PORT=3000 fiber-seed:latest


build-and-run: build run

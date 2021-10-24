# Running docker build using Dockerfile as the default file
build:
	docker build . -t docker-and-gorm
# If you wan to run docker with a different docker file, then you specify -f
build-file:
	docker build . -t docker-and-gorm -f Dockerfile-dev
run:
	go run main.go

up:
# name means the name you want to give the container, -p means publish, map the port 8080 on my system to 8080 on docker
# the last part is the image you want to run which is the image we built from our dockerfile
	#docker run -it --name=docker-gorm -p 8080:8080 --env PORT=8080 docker-and-gorm

	docker-compose up --build
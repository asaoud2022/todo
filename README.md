todo application
Todo application using Go, Gorm, go-fiber, redis, and PostgreSQL

Make sure you have git and docker installed.

Run:

git clone https://github.com/asaoud2022/todo/

docker-compose up
Debug:
docker-compose --file docker-compose-debug.yml up
DB runs at port 5438 external, backend is on 5001 external. Feel free to change it to whatever you like in the docker-compose.yml and backend/Dockerfile

** Repository Structure **
This repository contains Maintainer and its supporting packages and files.

Visual Code Workspace File - Open this with VS Code for development
Location: ./workspace.code-workspace
Main - Loads Config, Start Service Connections, Setup and starts Server
Location: ./main.go
Middleware - Custom Server Middlewares, Access Logger, Authentication, force HTTPS, force trailing slash, hsts, and suppress www
Location: ./app/middleware/*

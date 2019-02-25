#!/bin/bash
go build -o userService user/cmd/user/main.go
pkill userService
./userService


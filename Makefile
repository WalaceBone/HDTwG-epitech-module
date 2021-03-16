OUTPUT=main


all: clean build run

run:
	./build/main

clean:
	rm -rf build

fclean: clean all

build: .prefix
	go build -i -o build/$(OUTPUT) main.go

.prefix:
ifeq ($(OS), Windows_NT)
	if not exist build mkdir build
else
	mkdir build
endif
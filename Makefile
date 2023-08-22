spelling:
	pyspelling

output = ./builder
ifeq ($(GOOS),windows)
  output = './builder.exe'
endif

all: build
build:
	go build -o $(output) ./src/

generate:
	$(output) generate ./public/

watch:
	$(output) watch ./public/

serve: build
	$(output) serve ./public/

clean:
	rm -rf ./public/

all: build
build:
	go build -o $(output) ./src/

generate:
	$(output) generate ./docs/

watch:
	$(output) watch ./docs/

serve: build
	$(output) serve ./docs/

clean:
	rm -rf ./docs/

spelling:
	pyspelling

output = ./builder
ifeq ($(GOOS),windows)
  output = './builder.exe'
endif

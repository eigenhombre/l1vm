.PHONY: all install

all: l1vm

l1vm: main.go
	go build .

clean:
	rm -f l1vm

test:
	go test -v

install:
	go install .

docker:
	docker build -t l1vm .

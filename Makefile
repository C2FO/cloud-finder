
test:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

docker-image:
	docker build -t cloud-finder .

clean:
	rm -rf coverage.txt

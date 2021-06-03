# we will put our integration testing in this path
INTEGRATION_TEST_PATH?=./integration-tests

# this command will start a docker components that we set in docker-compose.yml
docker.start.components:
  docker-compose up -d;

# shutting down docker components
docker.stop:
  docker-compose down;

# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
  go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -run=$(INTEGRATION_TEST_SUITE_PATH)

# this command will trigger integration test with verbose mode
test.integration.debug:
  go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)

test:
go test -v  ./...

build:
GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o dmblogs .
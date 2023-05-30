mock-expected-keepers:
	mockgen -source=x/checkers/types/expected_keepers.go \
		-package testutil \
		-destination=x/checkers/testutil/expected_keepers_mocks.go
build-checkers-main:
	go build -o /Users/zhouwenzhe/go/bin/checkersd cmd/checkersd/main.go

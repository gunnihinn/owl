src := main.go
bin := owl

date := $(shell date '+%Y-%m-%d_%H:%M:%S')
epoch := $(shell date '+%s')
hash := $(shell git rev-parse --short HEAD)
LDFLAGS :=

$(bin): $(src)
	go build $(LDFLAGS) -o $(bin) $(src)

release: LDFLAGS = -ldflags "-X main.BuildDate=$(date) -X main.BuildEpoch=$(epoch) -X main.BuildHash=$(hash)"

.PHONY: release
release: $(bin) repo

.PHONY: clean
clean:
	rm -f $(bin)

.PHONY: repo
repo:
	git diff --quiet HEAD

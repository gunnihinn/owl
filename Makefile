src=main.go
bin=owl

$(bin): $(src)
	go build -o $(bin) $(src)

.PHONY: clean
clean:
	rm -f $(bin)

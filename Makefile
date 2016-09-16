EXEC = terraform-provider-function

all:
	go build -o $(EXEC)

build_linux:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o $(EXEC)

clean:
		rm -f $(EXEC)

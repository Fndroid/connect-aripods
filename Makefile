NAME=connect-airpods
BINDIR=bin
GOBUILD=go build -trimpath -ldflags "-s -w" 

all: amd64

amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)

clean:
	rm $(BINDIR)/*

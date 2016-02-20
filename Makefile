setup:
	go get github.com/parnurzeal/gorequest
	go get github.com/hashicorp/go-multierror
	go get github.com/stretchr/testify

test:
	go test -v github.com/ashwanthkumar/go-gocd

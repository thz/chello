
chello: deps
	env GOPATH=`pwd`:`pwd`/vendor go build -v
	
deps: 
	env GOPATH=`pwd`/vendor go get -d -v

.PHONY: deps

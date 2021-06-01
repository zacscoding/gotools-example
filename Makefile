.PHONY: doc generate zapcheck

doc:
	godoc -http :6060

generate:
	go generate ./...

# copy from https://github.com/google/exposure-notifications-server/blob/main/Makefile#L57
zapcheck:
	@command -v zapw > /dev/null 2>&1 || (cd $${TMPDIR} && go get github.com/sethvargo/zapw/cmd/zapw)
	@zapw ./...
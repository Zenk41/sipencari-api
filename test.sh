go test -v ./businesses/... --race -coverprofile=./coverage/profile.out -covermode=atomic
go tool cover -html=./coverage/profile.out -o ./coverage/coverage.html
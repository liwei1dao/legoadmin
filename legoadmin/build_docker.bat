set GOOS=linux
set CGO_ENABLED=0

go build -o ./bin/go_admin/a11/api.a11 ./services/api/main.go
go build -o ./bin/go_admin/a11/timer.a11 ./services/timer/main.go

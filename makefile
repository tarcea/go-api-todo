dev:
		CompileDaemon -command="./go-api-todo"
test:
		go test -v --cover ./...
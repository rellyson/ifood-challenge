PID_FILE = /tmp/server.pid

prepare:
	go mod download

build:
	go build -o bin/server main.go

kill:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 

run-bin:
	bin/server & echo $$! > $(PID_FILE)

run:
	go run main.go & echo $$! > $(PID_FILE)

watch: run
	fswatch -or --event=Updated  ./ ./main.go | \
	xargs -n1 make reload

reload: kill watch

# .PHONY is used for reserving tasks words
.PHONY: build kill reload watch prepare run run-bin

# supress echo commands on cli
.SILENT: build kill run run-bin watch reload prepare
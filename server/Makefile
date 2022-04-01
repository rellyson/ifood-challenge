PID_FILE = /tmp/server.pid

prepare:
	go mod download

build:
	go build -o bin/server main.go && clear

kill:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 

run:   
	bin/server & echo $$! > $(PID_FILE)

watch: build run
	fswatch -or --event=Updated  ./ ./main.go | \
	xargs -n1 make reload

reload: kill build run

# .PHONY is used for reserving tasks words
.PHONY: build kill reload watch prepare

# supress echo commands on cli
.SILENT: build kill run watch reload prepare
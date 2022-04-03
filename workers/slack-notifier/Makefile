prepare:
	pip install --no-cache-dir -r requirements.txt --quiet --exists-action i

run:
	python ./main.py

kill:
	- pkill -9 -f main.py 

watch: run
	fswatch -or --event=Updated  ./ | \
	xargs -n1 make reload

reload: kill watch

# .PHONY is used for reserving tasks words
.PHONY: watch run reload kill prepare

# supress echo commands on cli
.SILENT: reload
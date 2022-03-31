run:
	python main.py

watch: run
	fswatch -or --event=Updated  ./ | \
	xargs -n1 make reload

reload: run watch

# .PHONY is used for reserving tasks words
.PHONY: watch run reload

# supress echo commands on cli
.SILENT: run watch reload
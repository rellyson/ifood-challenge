prepare:
	pip install --no-cache-dir -r requirements.txt --quiet --exists-action i

run:
	python -u main.py

watch: run
	fswatch -or --event=Updated  ./ ./main.py | \
	xargs -n1 make reload

reload: watch

# .PHONY is used for reserving tasks words
.PHONY: watch run reload kill prepare

# supress echo commands on cli
.SILENT: run watch kill reload
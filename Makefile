fmt:
	gofumpt -w -l . && goimports -w -l . && dprint fmt ./README.md

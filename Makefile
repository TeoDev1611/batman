fmt:
	gofumpt -w -l . && goimports -w -l . && dprint fmt

fmt-win:
	powershell -Command 'iwr https://deno.land/install.ps1 -useb | iex'
	powershell -Command 'iwr https://dprint.dev/install.ps1 -useb | iex'

fmt-sh:
	curl -fsSL https://dprint.dev/install.sh | sh
	curl -fsSL https://deno.land/install.sh | sh

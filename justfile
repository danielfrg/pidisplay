default:
  just --list

[working-directory: 'controller']
build:
  go build -o tmp/main cmd/main.go

[working-directory: 'controller/frontend']
build-web:
  bun run build

[working-directory: 'controller']
server:
  #!/usr/bin/env bash
  if command -v air > /dev/null; then
      air
      echo "Watching..."
  else
      read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice
      if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then
          go install github.com/air-verse/air@latest
          air
          echo "Watching..."
      else
          echo "You chose not to install air. Exiting..."
          exit 1
      fi
  fi

[working-directory: 'controller/frontend']
web:
  bun run dev


[working-directory: 'controller']
go-fmt:

[working-directory: 'ansible']
ansible-display:
  uv run ansible-playbook display.yml

start-display:
  uv run --no-default-groups fastapi dev --host 0.0.0.0 src/pidisplay/display/app.py


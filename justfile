default:
  just --list

[working-directory: 'ansible']
ansible-display:
  uv run ansible-playbook display.yml

start-display:
  uv run --no-default-groups fastapi dev --host 0.0.0.0 src/pidisplay/display/app.py


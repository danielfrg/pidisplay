default:
  just --list

[working-directory: 'ansible']
ansible-display:
  uv run ansible-playbook display.yml


1. Prepare the environment  
sh prepare.sh
2. Start to tuning  
atune-adm tuning --project go_gc --detail go_gc_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project go_gc

1. Prepare the environment  
sh prepare.sh [test_path] [diskname]
2. Start to tuning  
atune-adm tuning --project iozone --detail tuning_iozone_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project iozone
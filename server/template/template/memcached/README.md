1. Prepare the environment  
sh prepare.sh
2. Start to tuning  
atune-adm tuning --project memcached_memaslap --detail memcached_memaslap_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project memcached_memaslap
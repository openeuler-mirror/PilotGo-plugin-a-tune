1. Prepare the environment  
sh prepare.sh
2. Start to tuning  
atune-adm tuning --project nginx --detail nginx_client.yaml  
atune-adm tuning --project nginx_http_long --detail nginx_http_long_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project nginx  
atune-adm tuning --restore --project nginx_http_long

1. Prepare the environment  
sh prepare.sh
2. Start to tuning  
atune-adm tuning --project tensorflow_train --detail tensorflow_train_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project tensorflow_train

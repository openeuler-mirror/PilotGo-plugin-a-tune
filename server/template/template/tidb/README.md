if you are the first one to use
please use this function
1.Prepare the initenvironment
sh config.sh IPserver servertargefile
#IPserver:your server ip
#servertargetfile:the file you wanted to store 
if you don't have the tidb
please run following point 
```
ssh root@IPserver "source ~/.bash_profile; cd servertargefile; sh install-tidb.sh clustername tidbversion ip "
```
PS
clustername: any name you wanted to call 
tidbversion: here are verions of tidb
 https://docs.pingcap.com/zh/tidb/stable/release-notes
 you can choose one to use
 example v5.1.0
ip: your server ip 
if you don't have sysbench
please run following point 
```
ssh root@IPserver "source ~/.bash_profile; cd servertargefile; sh install-sysbench.sh"
```
if you have already prepared above all
please run following instruction
1. Prepare the environment  
sh prepare.sh
2. Start to tuning  
atune-adm tuning --project tidb --detail tidb_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project tidb

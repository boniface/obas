### Create Network For Cluster 
`
docker network create -d bridge roachnet
`

### Start the Node 

`
docker run -d \
--name=roach1 \
--hostname=roach1 \
--net=roachnet \
-p 26257:26257 -p 8080:8080  \
-v /[home/local/folder/data]:/cockroach/cockroach-data"  \
cockroachdb/cockroach:v19.1.1 start --insecure
`

### Access the  Node 

`docker exec -it roach1 ./cockroach sql --insecure`

You can now run normal SQL commands 

`CREATE DATABASE obas;`


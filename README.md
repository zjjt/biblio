# biblio

#golang #go-micro #grpc #maygodblessthefoolofheart

To initialize a service in your monorepo install micro runtime on your localmachine and run

micro new --namespace=your.namespace --type=service --alias=user github.com/path/to/your/service

A monorepo project used as practice to understand how to use go-micro in a production app exemple from the micro-in-cn github repo.
Whereas theirs is supposed to be deployed as is, using go-micro directly, mine will concentrate on using docker the container's engine and kubernetes as a deployment platform.(i dont know if that also can be done with their way of using go-micro since i didnt see any Dockerfile in their repo)

I'll include comments in here telling which part of their tutorial i swapped with my preferences
as for a database i'll use POSTGRESQL instead of Mysql and the GORM orm for interacting with the DB

PART 1 USER SERVICE
Reading the part 1 it seems that they load all configuration from yml files which they parse in the basic folder and run the Init function.Since i will be using docker-compose for configuration, i dont need the conf folder.Instead of using directly the os.Getenv method directly,i'll parse the environment variables directly in the basic folder as they did.I renamed this basic folder into config , it holds the createConnection function to the POSTGRESDB and the etcd configuration.etcd is used here as a service registery and it is better suited for production instead of mdns.

They used the model folder to represent the database service,i noticed that they used private variables in order to initialize it, they used the same method of initializing with the handler service.My approach here is different ,both handler and repository will be initialized in the main function and passed as fields in their respective structs.I find this approach simpler. 



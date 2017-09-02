# pegasus

golang hybrid store compatible with omniql 


1. use nats for communication (initially); after it should laverage hybrids for this
2. use consul for store the resources (wrap consul in kv interface), this will allow use another kvs in the future
3. for the first version, please don't  add the versioning functionality of omniql
4. hear queries on the nats channel `query.*`, where * is the resource id
5. return query result on the channel `component.componentId`, where componentID is the resource making the query
6. hear command on the nats channel `command.*` where * is the resource id
7. return command result on the channel `component.componentId`, where componentID is the resource making the query
8. store the resources in the kv storage in the binary format
9. this implementation only accepts stacks of type state. see 3

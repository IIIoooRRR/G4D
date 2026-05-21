### Dispatch
[file](Go4D/connect/gispatch.go) \
Despite the name, it is responsible for receiving events. It passes the event to the queue in the connect format.RawEvent
### Gateway
[file](Go4D/connect/gateway.go)\
A simple method to connect to discord 
### Heartbeat 
[file](Go4D/connect/heartbeat.go)\
Once in a certain interval, received from discord, sends a message to discord that he is alive

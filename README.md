# What is RPC?
RPC is an inter-process communication that exchanges information between various
distributed systems. A computer called Alice can call functions (procedures) in another
computer called Bob in protocol format and can get the computed result back. Without
implementing the functionality locally, we can request things from a network that lies in
another place or geographical region.

### The entire process can be broken down into the following steps:
- Clients prepare function name and arguments to send
- Clients send them to an RPC server by dialing the connection
- The server receives the function name and arguments
- The server executes the remote process
- The message will be sent back to the client
- The client collects the data from the request and uses it appropriately

Did you see the magic? The client is running as an independent program from the server.
Here, both the programs can be on different machines, and computing can still be shared.
This is the core concept of distributed systems. The tasks are divided and given to various
RPC servers. Finally, the client collects the results and uses them to take further decisions.

Custom RPC code is only useful when the client and server are both written in Go. So, in
order to have the RPC server consumed by multiple services, we need to define the JSON-
RPC over HTTP. Then, any other programming language can send a JSON string and get
JSON as the result back.
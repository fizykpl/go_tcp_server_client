# go_tcp_server_client


1 run server on terminal :
```
src>go run server.go
```

expected 
```
Starting listening ...
Listening on 127.0.0.1:2012
Listening on 127.0.0.1:2011
```


2 run client  on terminal :
```
src>go run client.go 2011
```

expected 
```
Client try connect on 2011 port.
Text to send:
>>HERE PUT YOUR MESSAGE (ex ^A^B^C^) <<
```

after send message (clic enter)
```
Client connected to 127.0.0.1:2011
Message from server: A
Client connected to 127.0.0.1:2011
Message from server: B
Client connected to 127.0.0.1:2011
Message from server: C
Client connected to 127.0.0.1:2011
Message from server:
Text to send:
```


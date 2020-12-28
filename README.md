# devscope
The tool `devscope` allows the server to control the client, while the client can see what the server is doing.
It is useful for remote assistance, and explore non-interactive develop environments such as containers by cloud providers.

## Build and Install
Installation by building requires go `>= 1.15`. Simply, run
```shell
go install -i github.com/need-being/devscope/cmd/devscope
```

## Example Run

1. Server starts listening:

```
~$ devscope listen :5072
   • Listening at :5072
```

2. Client connects to the server, and transfer control to the server.

```
~$ devscope connect localhost:5072
   • Connected to localhost:5072
   • Running /bin/sh -i
   • Control transferred to Server
$
```

3. Server can explore on behalf of the client:

```
~$ devscope listen :5072
   • Listening at :5072
   • Accepted from 127.0.0.1:44694
$ echo hello world
hello world
$
```

# lokidb bencher
Benchmark tool for the lokidb-server

---

### Table of Contents
- [lokidb bencher](#lokidb-bencher)
    - [Table of Contents](#table-of-contents)
    - [Installation](#installation)
      - [Docker](#docker)
    - [Usage](#usage)


### Installation
#### Docker
```shell
docker pull yoyocode/lokidb-bencher
```

### Usage
```shell
$ docker run -it yoyocode/lokidb-bencher -addr <grpc_host>:<grpc_port> --help
Usage of /tmp/go-build1556986286/b001/exe/main:
  -addr string
        Server address (default "localhost:50051")
  -c string
        Command to benchmark options: [get, set, del] (default "get")
  -it int
        Number of iterations (default 10000)
```

```shell
$ docker run -it yoyocode/lokidb-bencher --addr 1.2.3.4:50051
Ran 10000 iterations in 380.736179ms
op/s 26265

$ docker run -it yoyocode/lokidb-bencher --addr 1.2.3.4:50051 -c set
Ran 10000 iterations in 376.878398ms
op/s 26534
```
## Architecture Layers

![ddd layers diagram image](./ddd-arch-diag.jpeg)

## Building application 
Use the `build.sh` executable to bring docker instances for the application and its dependencies with docker-compose.
```
./build.sh
```

## Layers Implementation Description

- Domain layer is having Entity and Repository, Service is combined with Entity for better representation. (this may vary depending on the language used.)


**reference**

[Using Domain-Driven Design(DDD)in Golang](https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5 "Using Domain-Driven Design(DDD)in Golang")

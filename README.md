# This repository is intent to have some code with gRPC and Protocol Buffers


Build image
```bash
docker build . -t image_name:tag
```

Run Container 
```bash
docker run -itd --rm --name go_container -v .:/home image_name:tag
```

Enter on container
```bash 
docker exec -it go_container bash
```

Inside the container run 
```bash
go mod tidy 
go run cmd/grpcServer/main.go
```

On another terminar enter on the container and run evans on read eval print loop mode
```bash
evans -r repl   
```

Generating gRPC service by protoc code generator
```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```





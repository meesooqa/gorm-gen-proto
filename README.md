# WIP: gorm-gen-proto

Copy `var/config.yml.example` to `var/config.yml`. Set `%path_maps%`, `%path_tmpl%`, `%proto_root%`.

These files should be created:
- `%path_maps%/import.json`,
- `%path_maps%/type.json`,
- `%path_tmpl%/proto3.tmpl`.

Install `protoc`:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

Add dependencies, e.g. `%proto_root%/google/api`:
```bash
git clone https://github.com/googleapis/googleapis.git
```
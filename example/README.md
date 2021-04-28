# Dependency Tracing Example

Files in this directory is created by:

```bash
go mod init
go get github.com/prometheus/client_golang@v1.10.0
go mod graph > go_mod_graph.output
```
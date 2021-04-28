# Go Dependency Trace

## Use Case

Knowing that a go module transitively depends on another module, find all the transitive dependencies that lead to that module.

For example, when you are alerted of a vulnerable dependency is used in your code, find out which dependency libraries actually cause that dependency to be included.

## Example

```bash
go run main.go -f example/go_mod_graph.output -l 4 -m golang.org/x/crypto
```

result:

```bash

2021/04/28 23:44:05 file parsed
2021/04/28 23:44:05 tree analyzed
2021/04/28 23:44:05 found 73 dependency paths to golang.org/x/crypto (7 versions)
2021/04/28 23:44:05 found 3 distinct top level paths to golang.org/x/crypto

golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9
golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2
golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
golang.org/x/crypto@v0.0.0-20191011191535-87dc89f01550
golang.org/x/crypto@v0.0.0-20181029021203-45a5f77698d3
golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793
golang.org/x/crypto@v0.0.0-20190510104115-cbcb75029529
============================================

path 0 from root to golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       golang.org/x/net@v0.0.0-20200625001655-4c5254603344
         golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9

path 1 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/prometheus/client_golang@v1.7.1
         github.com/prometheus/common@v0.10.0
           golang.org/x/net@v0.0.0-20190613194153-d28f0bde5980
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 2 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/prometheus/client_golang@v1.3.0
           github.com/prometheus/common@v0.7.0
             golang.org/x/net@v0.0.0-20190613194153-d28f0bde5980
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 3 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.etcd.io/etcd@v0.0.0-20191023171146-3cf2f69b5738
           golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 4 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.opencensus.io@v0.22.2
           golang.org/x/net@v0.0.0-20190620200207-3b0461eec859
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 5 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         golang.org/x/tools@v0.0.0-20200103221440-774c71fcf114
           golang.org/x/net@v0.0.0-20190620200207-3b0461eec859
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 6 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             golang.org/x/tools@v0.0.0-20191029190741-b9c20aec41a5
               golang.org/x/net@v0.0.0-20190620200207-3b0461eec859
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 7 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/atomic@v1.5.0
             golang.org/x/tools@v0.0.0-20191029041327-9cc4af7d6b2c
               golang.org/x/net@v0.0.0-20190620200207-3b0461eec859
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 8 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             go.uber.org/atomic@v1.5.0
               golang.org/x/tools@v0.0.0-20191029041327-9cc4af7d6b2c
                 golang.org/x/net@v0.0.0-20190620200207-3b0461eec859
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 9 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.etcd.io/etcd@v0.0.0-20191023171146-3cf2f69b5738
           golang.org/x/net@v0.0.0-20190813141303-74dc4d7220e7
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 10 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats.go@v1.9.1
           github.com/nats-io/nkeys@v0.1.0
             golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
               golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 11 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/nats.go@v1.9.1
             github.com/nats-io/nkeys@v0.1.0
               golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
                 golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 12 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats.go@v1.9.1
           github.com/nats-io/jwt@v0.3.0
             github.com/nats-io/nkeys@v0.1.0
               golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
                 golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 13 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/nats.go@v1.9.1
             github.com/nats-io/jwt@v0.3.0
               github.com/nats-io/nkeys@v0.1.0
                 golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
                   golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                     golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 14 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
             golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 15 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/nkeys@v0.1.3
             golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
               golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 16 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/jwt@v0.3.2
             github.com/nats-io/nkeys@v0.1.3
               golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
                 golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 17 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         golang.org/x/tools@v0.0.0-20200103221440-774c71fcf114
           golang.org/x/mod@v0.1.1-0.20191105210325-c90efee705ee
             golang.org/x/crypto@v0.0.0-20191011191535-87dc89f01550
               golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 18 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       golang.org/x/net@v0.0.0-20200625001655-4c5254603344
         golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9
           golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 19 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             honnef.co/go/tools@v0.0.1-2019.2.3
               golang.org/x/mod@v0.0.0-20190513183733-4bf6d317e70e
                 golang.org/x/crypto@v0.0.0-20190510104115-cbcb75029529
                   golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
                     golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 20 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin/zipkin-go@v0.2.2
           golang.org/x/net@v0.0.0-20190311183353-d8887717615a
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 21 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           github.com/openzipkin/zipkin-go@v0.2.1
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 22 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.opencensus.io@v0.22.2
           google.golang.org/grpc@v1.20.1
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 23 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           google.golang.org/grpc@v1.21.0
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 24 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           github.com/lightstep/lightstep-tracer-common/golang/gogo@v0.0.0-20190605223551-bc2310a04743
             google.golang.org/grpc@v1.21.0
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 25 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.etcd.io/etcd@v0.0.0-20191023171146-3cf2f69b5738
           google.golang.org/grpc@v1.23.1
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 26 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.opencensus.io@v0.22.2
           google.golang.org/grpc@v1.20.1
             golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 27 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           google.golang.org/grpc@v1.21.0
             golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 28 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           github.com/lightstep/lightstep-tracer-common/golang/gogo@v0.0.0-20190605223551-bc2310a04743
             google.golang.org/grpc@v1.21.0
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 29 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.opencensus.io@v0.22.2
           google.golang.org/grpc@v1.20.1
             golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 30 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           google.golang.org/grpc@v1.21.0
             golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 31 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           github.com/lightstep/lightstep-tracer-common/golang/gogo@v0.0.0-20190605223551-bc2310a04743
             google.golang.org/grpc@v1.21.0
               golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
                 golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                   golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                     golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 32 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.etcd.io/etcd@v0.0.0-20191023171146-3cf2f69b5738
           google.golang.org/grpc@v1.23.1
             golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 33 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           google.golang.org/grpc@v1.22.1
             golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 34 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin/zipkin-go@v0.2.2
           google.golang.org/grpc@v1.20.0
             golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 35 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           github.com/openzipkin/zipkin-go@v0.2.1
             google.golang.org/grpc@v1.20.0
               golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
                 golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                   golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                     golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 36 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         google.golang.org/grpc@v1.26.0
           github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473
             google.golang.org/grpc@v1.23.0
               golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
                 golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                   golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                     golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 37 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin/zipkin-go@v0.2.2
           google.golang.org/grpc@v1.20.0
             golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 38 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           github.com/openzipkin/zipkin-go@v0.2.1
             google.golang.org/grpc@v1.20.0
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 39 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           golang.org/x/lint@v0.0.0-20190930215403-16217165b5de
             golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 40 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/atomic@v1.5.0
             golang.org/x/lint@v0.0.0-20190930215403-16217165b5de
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 41 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             go.uber.org/atomic@v1.5.0
               golang.org/x/lint@v0.0.0-20190930215403-16217165b5de
                 golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                   golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                     golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 42 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             golang.org/x/lint@v0.0.0-20190930215403-16217165b5de
               golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 43 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           go.opencensus.io@v0.20.2
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 44 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.etcd.io/etcd@v0.0.0-20191023171146-3cf2f69b5738
           google.golang.org/grpc@v1.23.1
             golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 45 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           google.golang.org/grpc@v1.22.1
             golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 46 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         google.golang.org/grpc@v1.26.0
           github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473
             google.golang.org/grpc@v1.23.0
               golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 47 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         google.golang.org/grpc@v1.26.0
           golang.org/x/net@v0.0.0-20190311183353-d8887717615a
             golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 48 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           google.golang.org/grpc@v1.22.1
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 49 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           go.opencensus.io@v0.20.2
             google.golang.org/api@v0.3.1
               golang.org/x/tools@v0.0.0-20190312170243-e65039ee4138
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 50 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/smartystreets/goconvey@v1.6.4
           golang.org/x/tools@v0.0.0-20190328211700-ab21143f2384
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 51 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin/zipkin-go@v0.2.2
           google.golang.org/grpc@v1.20.0
             golang.org/x/net@v0.0.0-20190311183353-d8887717615a
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 52 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/openzipkin-contrib/zipkin-go-opentracing@v0.4.5
           github.com/openzipkin/zipkin-go@v0.2.1
             google.golang.org/grpc@v1.20.0
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 53 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         google.golang.org/grpc@v1.26.0
           github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473
             google.golang.org/grpc@v1.23.0
               golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 54 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             honnef.co/go/tools@v0.0.1-2019.2.3
               golang.org/x/tools@v0.0.0-20190621195816-6e04913cbbac
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 55 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           go.opencensus.io@v0.20.2
             google.golang.org/api@v0.3.1
               go.opencensus.io@v0.20.1
                 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
                   golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 56 from root to golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           github.com/lightstep/lightstep-tracer-common/golang/gogo@v0.0.0-20190605223551-bc2310a04743
             golang.org/x/net@v0.0.0-20190603091049-60506f45cf65
               golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2

path 57 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats.go@v1.9.1
           github.com/nats-io/nkeys@v0.1.0
             golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 58 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/nats.go@v1.9.1
             github.com/nats-io/nkeys@v0.1.0
               golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 59 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats.go@v1.9.1
           github.com/nats-io/jwt@v0.3.0
             github.com/nats-io/nkeys@v0.1.0
               golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 60 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/nats.go@v1.9.1
             github.com/nats-io/jwt@v0.3.0
               github.com/nats-io/nkeys@v0.1.0
                 golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 61 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 62 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/nkeys@v0.1.3
             golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 63 from root to golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/nats-io/nats-server/v2@v2.1.2
           github.com/nats-io/jwt@v0.3.2
             github.com/nats-io/nkeys@v0.1.3
               golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4

path 64 from root to golang.org/x/crypto@v0.0.0-20191011191535-87dc89f01550 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         golang.org/x/tools@v0.0.0-20200103221440-774c71fcf114
           golang.org/x/mod@v0.1.1-0.20191105210325-c90efee705ee
             golang.org/x/crypto@v0.0.0-20191011191535-87dc89f01550

path 65 from root to golang.org/x/crypto@v0.0.0-20181029021203-45a5f77698d3 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/hashicorp/consul/api@v1.3.0
           github.com/hashicorp/serf@v0.8.2
             github.com/hashicorp/memberlist@v0.1.3
               golang.org/x/crypto@v0.0.0-20181029021203-45a5f77698d3

path 66 from root to golang.org/x/crypto@v0.0.0-20181029021203-45a5f77698d3 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/hashicorp/consul/api@v1.3.0
           github.com/hashicorp/serf@v0.8.2
             github.com/hashicorp/mdns@v1.0.0
               golang.org/x/crypto@v0.0.0-20181029021203-45a5f77698d3

path 67 from root to golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/prometheus/client_golang@v1.7.1
         github.com/prometheus/common@v0.10.0
           github.com/prometheus/client_golang@v1.0.0
             github.com/prometheus/common@v0.4.1
               github.com/sirupsen/logrus@v1.2.0
                 golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793

path 68 from root to golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.etcd.io/etcd@v0.0.0-20191023171146-3cf2f69b5738
           github.com/prometheus/client_golang@v1.0.0
             github.com/prometheus/common@v0.4.1
               github.com/sirupsen/logrus@v1.2.0
                 golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793

path 69 from root to golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/prometheus/client_golang@v1.3.0
           github.com/prometheus/common@v0.7.0
             github.com/prometheus/client_golang@v1.0.0
               github.com/prometheus/common@v0.4.1
                 github.com/sirupsen/logrus@v1.2.0
                   golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793

path 70 from root to golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           go.opencensus.io@v0.20.2
             github.com/prometheus/client_golang@v0.9.3-0.20190127221311-3c4408c8b829
               github.com/prometheus/common@v0.2.0
                 github.com/sirupsen/logrus@v1.2.0
                   golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793

path 71 from root to golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         github.com/lightstep/lightstep-tracer-go@v0.18.1
           go.opencensus.io@v0.20.2
             google.golang.org/api@v0.3.1
               go.opencensus.io@v0.20.1
                 github.com/prometheus/client_golang@v0.9.3-0.20190127221311-3c4408c8b829
                   github.com/prometheus/common@v0.2.0
                     github.com/sirupsen/logrus@v1.2.0
                       golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793

path 72 from root to golang.org/x/crypto@v0.0.0-20190510104115-cbcb75029529 :
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0
         go.uber.org/zap@v1.13.0
           go.uber.org/multierr@v1.3.0
             honnef.co/go/tools@v0.0.1-2019.2.3
               golang.org/x/mod@v0.0.0-20190513183733-4bf6d317e70e
                 golang.org/x/crypto@v0.0.0-20190510104115-cbcb75029529
============================================

distinct top path 0 occurs 69 time(s)
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/go-kit/kit@v0.10.0

distinct top path 1 occurs 2 time(s)
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       golang.org/x/net@v0.0.0-20200625001655-4c5254603344

distinct top path 2 occurs 2 time(s)
 github.com/gmm1900/deptrace/example@
   github.com/prometheus/client_golang@v1.10.0
     github.com/prometheus/common@v0.18.0
       github.com/prometheus/client_golang@v1.7.1

```
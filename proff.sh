#!/usr/bin/env bash
curl http://127.0.0.1:6060/debug/pprof/heap?seconds=20 -o mem_out.txt
curl http://127.0.0.1:6060/debug/pprof/profile?seconds=20 -o cpu_out.txt

go tool pprof -svg -alloc_objects ./prof mem_out.txt > mem_ao.svg
go tool pprof -svg ./prof cpu_out.txt > cpu.svg
go tool pprof -svg -alloc_space ./prof mem_out.txt > mem_sp.svg

#ab -c 20 -n 100 http://127.0.0.1:6060/
#apt install graphviz

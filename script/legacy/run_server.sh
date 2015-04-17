#!/bin/sh
/usr/bin/mysqld_safe &
sleep 10s
goop exec go run server.go

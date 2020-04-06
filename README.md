# tls-cert-expire
检查TLS证书过期日期

## 使用方法
```bash
➜  tls-cert-expire git:(master) ✗ ./tls-cert-expire 
Usage: ./tls-cert-expire -d domain [-p port] [-t timeout] [-s]
```

```bash
➜  tls-cert-expire git:(master) ✗ ./tls-cert-expire -d www.io-notes.com
2020-06-24 08:19:05 +0800 CST
```

```bash
➜  tls-cert-expire git:(master) ✗ ./tls-cert-expire -d www.io-notes.com -s
1592957945
```

```bash
➜  tls-cert-expire git:(master) ✗ ./tls-cert-expire -h
Usage of ./tls-cert-expire:
  -d domain
        domain: 域名 IP
  -p port
        port : 端口 (default 443)
  -s    seconds since 1970-01-01 00:00:00 UTC
  -t timeout
        timeout : 连接超时 (default 5)
```
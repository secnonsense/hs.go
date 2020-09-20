# hs.go
Simple HTTP/HTTPS web server written in go
| Command| Explanation|
|:--|:--|
|hs -p 8081|Specify the port number for the server (8080 is the default)|
|hs -t |Run the web server in HTTPS mode with the integrated self-signed certificate and key|
|hs -c|Delete the self-signed cert and key from the current director|
|hs -f filename|Serve up a file over HTTP (or HTTPS with -t)|


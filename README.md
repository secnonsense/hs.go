# hs.go
Simple HTTP/HTTPS web server written in go

hs -p 8081  < specify the port number for the server (8080 is the default)

hs -t < run the web server in HTTPS mode with the integrated self-signed certificate and key

hs -c < delete the self-signed cert and key from the current directory

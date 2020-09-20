# hs.go
Simple HTTP/HTTPS web server written in go


<table>
<tr>
<th>hs -p 8081
<th>Specify the port number for the server (8080 is the default)
</tr>
<th>hs -t   
<th>Run the web server in HTTPS mode with the integrated self-signed certificate and key
</tr>
<th>hs -c
<th>Delete the self-signed cert and key from the current director
</tr>
<th>hs -f filename
<th>Serve up a file over HTTP (or HTTPS with -t)
</tr>
</table>

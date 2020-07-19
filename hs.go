package main


import (
    "github.com/integrii/flaggy"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "bytes"
    "io/ioutil"
    
)

var file = ""

func createCertnKey(rebuild int) {
    cert:="-----BEGIN CERTIFICATE-----\nMIICLjCCAbUCCQCqjs7a3Kbg+zAKBggqhkjOPQQDAjCBgDELMAkGA1UEBhMCdXMxCzAJBgNVBAgMAmhpMQ8wDQYDVQQHDAZtYXRyaXgxDTALBgNVBAoMBHRlc3QxCzAJBgNVBAsMAmVzMRYwFAYDVQQDDA1zb21ldGhpbmcubGNsMR8wHQYJKoZIhvcNAQkBFhBtZUBzb21ldGhpbmcubGNsMB4XDTIwMDcxNzEzNDU0MloXDTMwMDcxNTEzNDU0MlowgYAxCzAJBgNVBAYTAnVzMQswCQYDVQQIDAJoaTEPMA0GA1UEBwwGbWF0cml4MQ0wCwYDVQQKDAR0ZXN0MQswCQYDVQQLDAJlczEWMBQGA1UEAwwNc29tZXRoaW5nLmxjbDEfMB0GCSqGSIb3DQEJARYQbWVAc29tZXRoaW5nLmxjbDB2MBAGByqGSM49AgEGBSuBBAAiA2IABHjrzQKbirpKOWQfnwp0vc7A0awf82qr2Xb/JAtz7xUJN23WWSgEP5IAWxitxner0KKTlpx/ku54oKqeL9q+hKgbYwg3qMktPDmWkXZIDit8G6lE51H4gVFhOE0SBsYRWjAKBggqhkjOPQQDAgNnADBkAjBuWpInfs8g2vA/nHW/4Cwmv2aAxG36hZ/9OQqgr4VByAClzEgj19uLvD42D1EDXHYCMAsDhaj2BD7yDBrw5rOuQVwuvX8F7W4PaOmwU7VTId1/LV25QdfsdTrj55y2xblbnQ==\n-----END CERTIFICATE-----"

    key:="-----BEGIN EC PARAMETERS-----\nBgUrgQQAIg==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDCwwKl19hWLU7DQPg5iBgs/oMICgu8qxQvKwgjUaLFMFt/iyScE90UV\nPTCNhIgMc+egBwYFK4EEACKhZANiAAR4680Cm4q6SjlkH58KdL3OwNGsH/Nqq9l2\n/yQLc+8VCTdt1lkoBD+SAFsYrcZ3q9Cik5acf5LueKCqni/avoSoG2MIN6jJLTw5\nlpF2SA4rfBupROdR+IFRYThNEgbGEVo=\n-----END EC PRIVATE KEY-----"
    
    if rebuild == 1 || rebuild == 3 {
        writeFile(cert,"server.crt")
    }
    if rebuild == 2 || rebuild == 3 {
        writeFile(key,"server.key")  
    }
}

func writeFile(infile string,outfile string){
        
    s, err := os.Create(outfile)
        if err != nil {
            fmt.Println(err)
            return
    }
        b, err := s.WriteString(infile)
        if err != nil {
            fmt.Println(err)
            s.Close()
            return
    }
        fmt.Println("\n============ File Saved =============")
        fmt.Println(b, "bytes were successfully written to",outfile,"\n")
        err = s.Close()
        if err != nil {
            fmt.Println(err)
            return
    }
}

func HTTPServer(w http.ResponseWriter, r *http.Request) {

    if len(file) > 0 {
        data, _ := ioutil.ReadFile(file)
        http.ServeContent(w, r, file , time.Now(), bytes.NewReader(data))
    } else {
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
        fmt.Fprintf(w, "Header: %q = %q\n", k, v)
        fmt.Printf("Header: %q = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Address = %q\n", r.RemoteAddr)
    fmt.Printf("Host: %q\n", r.Host)
    fmt.Printf("Remote Address: %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form %q = %q\n", k, v)
        fmt.Printf("Form %q = %q\n", k, v)
	}


    fmt.Fprintf(w, "Successful Connection")
    fmt.Printf("Successful Connection\n\n")
    }
}


func main() {
    
    var p ="8080"
    var m ="HTTP"
    var t = false
    var c = false
    var rebuild = 0
    flaggy.String(&p,"p","port" ,"Input the Port")
    flaggy.String(&file,"f","file" ,"Input file to serve")
    flaggy.Bool(&t,"t","tls" ,"Start TLS server")
    flaggy.Bool(&c,"c","cleanup" ,"Cleanup Cert and Key")
    flaggy.Parse()
    
    if c {
         os.Remove("server.crt")
         os.Remove("server.key")

         fmt.Println("Files Deleted")
         os.Exit(1)
    }
    
    http.HandleFunc("/", HTTPServer)

    pp:=":"+p
    if t {
        if _, err := os.Stat("server.crt"); err == nil {
            print("server.crt exists\n")
        } else if os.IsNotExist(err) {
            print("server.crt doesn't exist\n")
            rebuild=1
        }
        if _, err := os.Stat("server.key"); err == nil {
            print("server.key exists\n")
        } else if os.IsNotExist(err) {
            print("server.key doesn't exist\n")
            rebuild=rebuild+2
        }
        
        if rebuild>0 {
            createCertnKey(rebuild)
        }
        m="HTTPS"
        fmt.Printf("\nStarting %s server at port%s\n\n", m,pp)
        if err := http.ListenAndServeTLS(pp, "server.crt", "server.key", nil); err != nil {
            log.Fatal(err)
        }
    } else {
        fmt.Printf("\nStarting %s server at port%s\n\n", m,pp)
        if err := http.ListenAndServe(pp, nil); err != nil {
            log.Fatal(err)
        }
    }
    
}

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


func uploadFile(w http.ResponseWriter,r *http.Request){
    if r.Method != "POST"{
        http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
		fmt.Println("1")
        return 
    }

    r.ParseMultipartForm(10<<20)

    file, handler, err := r.FormFile("myfile")
    if err != nil{
        fmt.Println("Error in reading the file")
        fmt.Println(err)
		fmt.Println("2")
        return 
    }

    defer file.Close() 

    dst, err := os.Create(handler.Filename)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("3")
        return 
    }

    defer dst.Close()

    if _, err:= io.Copy(dst, file); err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("4")
        return
    }

    fmt.Fprintf(w, "Successfully uploaded\n")
}

func getFileMD5Hash(w http.ResponseWriter, r * http.Request){
    if r.Method != "GET"{
        http.Error(w, "Invalid Error", http.StatusInternalServerError)
        return
    }

    filename:= r.URL.Query().Get("filename")
    if filename ==""{
        http.Error(w, "Filename is Required", http.StatusBadRequest)
        return 
    }

    file, err := os.Open(filename)
    if err != nil{
        http.Error(w, "File not Found", http.StatusNotFound)
        return 
    }
    
    defer file.Close()

    hash := md5.New()

    if _,err := io.Copy(hash,file); err != nil{
        http.Error(w, "Error calculating hash", http.StatusInternalServerError)
        return 
    }

    hashInbytes := hash.Sum(nil)[:16]
    md5string:= hex.EncodeToString(hashInbytes)
    fmt.Fprintf(w, "MD5hash: %s\n",md5string)

}


func main(){
    http.HandleFunc("/upload",uploadFile)
    http.HandleFunc("/checkmd5",getFileMD5Hash)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err!= nil{
        log.Fatalf("Failed to start server: %v\n",err)
    }
}


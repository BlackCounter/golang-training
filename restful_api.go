package main
import (
	"fmt"
//	"log"
	"net/http"
)
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome to HomePage")


}
func handleRequests(){
	http.HandleFunc("/",homePage)
	http.ListenAndServe(":10000",nil)
}
func main() {
    handleRequests()
}
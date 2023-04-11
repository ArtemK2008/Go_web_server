package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/factorial", HandleFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	return ans
}

func HandleFactorial(writer http.ResponseWriter, request *http.Request) {
	num := request.FormValue("num")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(writer, err.Error(), 404)
		return
	}
	io.WriteString(writer, strconv.Itoa(factorial(n)))
}

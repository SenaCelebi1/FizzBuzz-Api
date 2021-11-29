package middleware

import (
	"container/list"
	"encoding/json"
	models "fizzbuzzapi/Models"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func FizzbuzzFunction(w http.ResponseWriter, r *http.Request) {

	var fizzbuzz models.FizzBuzz
	_ = json.NewDecoder(r.Body).Decode(&fizzbuzz)

	params := mux.Vars(r)
	countStr := params["count"]
	fmt.Println("count:")
	fmt.Println(countStr)
	count, _ := strconv.Atoi(countStr)

	list := list.New()

	for i := 1; i <= count; i++ {

		if (i)%15 == 0 {
			list.PushBack("FizzBuzz")

		} else if (i)%5 == 0 {
			list.PushBack("Buzz")

		} else if (i)%3 == 0 {
			list.PushBack("Fizz")

		} else {
			var x string = strconv.Itoa(i)
			list.PushBack(x)

		}
	}

	for element := list.Front(); element != nil; element = element.Next() {

		fizzbuzz.FizzBuzz += element.Value.(string)
		fizzbuzz.FizzBuzz += " "

	}
	json.NewEncoder(w).Encode(fizzbuzz)

}

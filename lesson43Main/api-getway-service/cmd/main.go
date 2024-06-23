package main

import "bill_service/api"

func main() {
	panic(api.Routes().ListenAndServe())
}

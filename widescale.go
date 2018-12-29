package main

func main() {
	err := CreateIndex("wikidump")
	if err != nil {
		panic(err)
	}
	//b, err := json.MarshalIndent(InvertedIndex, "", "  ")
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Print(string(b))
	/*	var keys []string
		for key, val := range InvertedIndex {
			keys = append(keys, key)
			fmt.Println("Key : " + key)
			for v, z := range val {
				fmt.Println("Value :" + string(z) + string(v))
			}
		}*/
	//router := mux.NewRouter()
	//log.Fatal(http.ListenAndServe(":8000", router))
}

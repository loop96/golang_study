package persist

import "fmt"

func GetItemSaver() chan interface{} {
	//Create ItemSaver channel to handle data
	saverChannel := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-saverChannel
			fmt.Printf("ItemSaver got item ,count : %d item :\n %v \n\n", itemCount, item)
			itemCount++
			save(item)
		}
	}()
	return saverChannel
}

func save(item interface{}) {
	// put data in Elasticsearch

}

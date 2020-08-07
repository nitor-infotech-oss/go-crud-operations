package db
//Fake backend created for returning data for different verbs
var database = make(map[string]interface{})

func FindAll() []interface{} {
	items := make([]interface{}, 0, len(database))
	for _, v := range database {
		items = append(items, v)
	}

	return items
}

func FindBy(key string) (interface{}, bool) {
	patient, ok := database[key]

	return patient, ok
}

func Save(key string, item interface{}) {
	database[key] = item
}

func Remove(key string) {
	delete(database, key)
}

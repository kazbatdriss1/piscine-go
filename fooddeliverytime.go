package piscine

type food struct {
	preptime int
}

var menu = map[string]food{
	"burger":  {15},
	"chips":   {10},
	"nuggets": {12},
}

func FoodDeliveryTime(order string) int {
	item, exists := menu[order]
	if !exists {
		return 404
	}
	return item.preptime
}

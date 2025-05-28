package moudle

func TranslateFoodCategory(category string) string {
	translations := map[string]string{
		"MainDish":  "主菜",
		"Dessert":   "甜点",
		"Fruit":     "水果",
		"Beverage":  "饮料",
		"Soup":      "汤",
		"Meat":      "肉类",
		"Specialty": "特色菜",
		"Vegetable": "蔬菜",
		"Snack":     "小吃",
		"Sauce":     "酱料",
	}

	if translated, ok := translations[category]; ok {
		return translated
	}
	return category // 未找到翻译时返回原字符串
}

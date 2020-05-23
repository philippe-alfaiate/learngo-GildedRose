package main

import "strings"

//Item format of item
type Item struct {
	name            string
	sellIn, quality int
}

//UpdateQuality to include Conjured items
func UpdateQuality(items []*Item) {
	for _, item := range items {
		step := -1
		coef := 1
		if item.sellIn <= 0 {
			coef = 2
		}
		if strings.HasPrefix(item.name, "Backstage passes") {
			if item.sellIn <= 0 {
				item.quality = 0
				coef = 0
			} else {
				coef = 1
				step = 1
				if item.sellIn <= 10 {
					step++
				}
				if item.sellIn <= 5 {
					step++
				}
			}

		} else if strings.HasPrefix(item.name, "Aged Brie") {
			step = 1
		} else if strings.HasPrefix(item.name, "Sulfuras") {
			continue
		} else if strings.HasPrefix(item.name, "Conjured") {
			coef *= 2
		}
		item.quality += step * coef
		if item.quality > 50 {
			item.quality = 50
		} else if item.quality < 0 {
			item.quality = 0
		}
		item.sellIn += -1
	}
}

//OldUpdateQuality to update the quality of each item over one day
func OldUpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				if items[i].name != "Sulfuras, Hand of Ragnaros" {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}

}

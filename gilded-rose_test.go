package main

import (
	"fmt"
	"testing"
)

func Test_OldAgainstNewUpdate(t *testing.T) {
	var itemOld = []*Item{
		&Item{"+5 Dexterity Vest", 10, 20},
		&Item{"Aged Brie", 2, 0},
		&Item{"Aged Brie", 4, 0},
		&Item{"Elixir of the Mongoose", 5, 7},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
	}
	var itemNew = []*Item{
		&Item{"+5 Dexterity Vest", 10, 20},
		&Item{"Aged Brie", 2, 0},
		&Item{"Aged Brie", 4, 0},
		&Item{"Elixir of the Mongoose", 5, 7},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
	}

	day := 155
	for i := 0; i < day; i++ {
		OldUpdateQuality(itemOld)
		UpdateQuality(itemNew)

		res := ""

		for i := 0; i < len(itemOld); i++ {
			res = ""
			tmp := ">"
			err := 0
			if itemOld[i].name != itemNew[i].name {
				tmp = "!"
				err++
			}
			res += fmt.Sprintf("Name [%s] %s [%s]", itemOld[i].name, tmp, itemNew[i].name)
			tmp = ">"
			if itemOld[i].sellIn != itemNew[i].sellIn {
				tmp = "!"
				err++
			}
			res += fmt.Sprintf(", sellIn [%d] %s [%d], ", itemOld[i].sellIn, tmp, itemNew[i].sellIn)
			tmp = ">"
			if itemOld[i].quality != itemNew[i].quality {
				tmp = "!"
				err++
			}

			res += fmt.Sprintf(", quality [%d] %s [%d] %d/3", itemOld[i].quality, tmp, itemNew[i].quality, 3-err)

			pass := "pass"
			if err != 0 {
				pass = "fail"
				t.Errorf("%s %s", res, pass)
			} else {
				//println(res, pass)
			}

		}
	}

}

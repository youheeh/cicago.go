package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type entry struct {
	val int
	key int
}

type entries []entry

func (s entries) Len() int           { return len(s) }
func (s entries) Less(i, j int) bool { return s[i].val < s[j].val }
func (s entries) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	fmt.Println("CICAGO")
	fmt.Println("[ルール]")
	fmt.Println("[12回ダイスを投げて、一番ダイスの目が多いプレイヤーの勝ち]")
	fmt.Println(`プレイヤーは何人ですか？`)

	var playercount int
	fmt.Scan(&playercount)
	if playercount < 5 {
		fmt.Println(playercount, `人で対戦します`)

	} else {
		//４人以上の場合不可と表示する
		fmt.Println(`４人以上では対戦できません`)
		return
	}

	var mapPlayer = make(map[int]int)

	//プレイヤー分作成
	//var randomArray [4]int
	var randomArray = make([]int, playercount)

	for i := 1; i <= 5; i++ {

		for i := 0; i < playercount; i++ {
			//
			fmt.Println("プレイヤー", (i + 1), "ダイスを投げてください")
			randomArray[i] = rand.Intn(6) + 1
			fmt.Println(randomArray[i], `が出ました`)

			// add point to mapPlayer
			mapPlayer[i+1] += randomArray[i]

			//scanner := bufio.NewScanner(os.Stdin)
			//scanner.Scan()
		}

	}

	var es entries
	for k, v := range mapPlayer {
		es = append(es, entry{val: v, key: k})
	}

	sort.Sort(sort.Reverse(es))

	for _, v := range es {
		fmt.Println("player number: ", v.key, " has point: ", v.val)
		//fmt.Println("player number: %d has point: %d\n", v.key, v.val)

		//fmt.Printf("%v\t%v\n", 'プレイヤー' + v.key, v.val)
	}

}

func player() {

	var playercount string
	fmt.Scan(&playercount)
	fmt.Println(playercount, `で対戦します`)
}

// mapNumber := make(map[int]int, 0)

// 	for i := 0; i < 5; i++ {
// 		mapNumber[i] = i

// 	}

// 	for key, value:= range mapNumber {
// 	   fmt.Println("key tatoeba: ", key, " value tatoeba: ", value)
// 	}

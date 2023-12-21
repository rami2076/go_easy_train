package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type branchRecord struct {
	cd         int
	name       string
	officeCd   int
	officeName string
}

type branch struct {
	cd      int
	name    string
	offices []office
}

type office struct {
	officeCd   int
	officeName string
}

func main2() {

	tom := branchRecord{
		cd:         1,
		name:       "name",
		officeCd:   1,
		officeName: "officeName",
	}

	nick := branchRecord{
		cd:         1,
		name:       "name",
		officeCd:   2,
		officeName: "officeName2",
	}

	branches := []branchRecord{tom, nick}

	branchesByCd := map[int][]branchRecord{}

	for _, v := range branches {
		branchesByCd[v.cd] = append(branchesByCd[v.cd], v)
	}

	fmt.Println(branchesByCd)

	root := []branch{}

	for _, values := range branchesByCd {

		o := []office{}

		for _, value := range values {
			o = append(o, office{
				officeCd:   value.officeCd,
				officeName: value.officeName,
			})
		}

		b := branch{
			cd:      values[0].cd,
			name:    values[0].name,
			offices: o,
		}

		root = append(root, b)
	}

	a, e := json.Marshal(root)

	fmt.Println(e)
	fmt.Println(root)

	fmt.Printf("%s\n", a)
	os.Stdout.Write(a)

}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BranchRecord struct {
	cd         int
	name       string
	officeCd   int
	officeName string
}

type Branch struct {
	cd      int
	name    string
	offices []Office
}

type Office struct {
	officeCd   int
	officeName string
}

func main2() {

	record1 := BranchRecord{
		cd:         1,
		name:       "name",
		officeCd:   1,
		officeName: "officeName",
	}

	record2 := BranchRecord{
		cd:         1,
		name:       "name",
		officeCd:   2,
		officeName: "officeName2",
	}

	branchRecords := []BranchRecord{record1, record2}

	branchRecordsByCd := map[int][]BranchRecord{}

	for _, v := range branchRecords {
		branchRecordsByCd[v.cd] = append(branchRecordsByCd[v.cd], v)
	}

	fmt.Println(branchRecordsByCd)

	var branches []Branch

	for _, values := range branchRecordsByCd {

		var o []Office

		for _, value := range values {
			o = append(o, Office{
				officeCd:   value.officeCd,
				officeName: value.officeName,
			})
		}

		b := Branch{
			cd:      values[0].cd,
			name:    values[0].name,
			offices: o,
		}

		branches = append(branches, b)
	}

	a, e := json.Marshal(branches)

	fmt.Println(e)
	fmt.Println(branches)

	fmt.Printf("%s\n", a)
	os.Stdout.Write(a)

}

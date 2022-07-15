package extract

import (
	"bufio"
	"fmt"
	"jubilant-potato/internal/types"
	"os"
	"strings"
)

// Record struct in go
// something like
// type Info struct{
//     Uid int `json:"uid"`
//     Data SuperData `json:"data"`
// }

func ReadStructFromGoFile(fs *os.File) (structsArray []types.GoStruct, err error) {
	scanner := bufio.NewScanner(fs)

	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		if hasStruct(strs) {
			// gs represent a struct
			gs, err := extractStruct(scanner, strs)
			if err != nil {
				return structsArray, err
			}
			structsArray = append(structsArray, gs)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scan file error:", err)
		return structsArray, err
	}

	return structsArray, err
}

// Usually, a go struct starts with
// type XXX struct {
//     ...
// }
// so check the start sign {
func hasStruct(strs []string) bool {
	for _, v := range strs {
		if v == "{" {
			return true
		}
	}

	return false
}

func checkStructEnd(strs []string) bool {
	for _, v := range strs {
		if v == "}" {
			return true
		}
	}

	return false
}

// Try to extract a struct
func extractStruct(sc *bufio.Scanner, firstLine []string) (goStruct types.GoStruct, err error) {
	// get the name
	goStruct.Name = firstLine[1]
	goStruct.Members = make(map[string]string)
	for sc.Scan() {
		strs := strings.Fields(sc.Text())

		if checkStructEnd(strs) {
			break
		}

		memberAttributes := strs[1]
		nameStartIndex := strings.Index(strs[2], "\"") + 1
		nameEndIndex := strings.LastIndex(strs[2], "\"")
		memberName := strs[2][nameStartIndex:nameEndIndex]

		goStruct.Members[memberName] = memberAttributes
	}

	if err := sc.Err(); err != nil {
		fmt.Println("scan file error:", err)
		return goStruct, err
	}

	return goStruct, err
}

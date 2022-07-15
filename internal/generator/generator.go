package generator

import (
	"bufio"
	"fmt"
	"jubilant-potato/internal/types"
	"os"
	"strings"
)

func GenerateTargetFile(fs *os.File, structsArray []types.GoStruct) (err error) {
	w := bufio.NewWriter(fs)

	_, err = w.WriteString("use serde::{Serialize, Deserialize};\n\n")
	if err != nil {
		return err
	}

	for _, sample := range structsArray {
		err := writeStruct(w, sample)
		if err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

func writeStruct(w *bufio.Writer, sample types.GoStruct) (err error) {
	_, err = w.WriteString("#[derive(Serialize,Deserialize)]\n")
	if hasWriteStringError(err) {
		return err
	}
	_, err = w.WriteString(fmt.Sprintf("pub struct %s{\n", sample.Name))
	if hasWriteStringError(err) {
		return err
	}

	for name, attr := range sample.Members {

		// Example:
		// []int in go
		// Vec<i64> in rust
		if strings.Contains(attr, "[]") {
			if targetAttr, ok := types.TransDict[attr[2:]]; ok {
				attr = strings.Replace(attr, attr[2:], targetAttr, 1)
			}
			attr = fmt.Sprintf("Vec<%s>", attr[2:])
		}

		// trans attribute
		if targetAttr, ok := types.TransDict[attr]; ok {
			attr = targetAttr
		}

		_, err = w.WriteString(fmt.Sprintf("\tpub %s:%s,\n", name, attr))
		if hasWriteStringError(err) {
			return err
		}
	}

	_, err = w.WriteString("}\n\n")
	if hasWriteStringError(err) {
		return err
	}

	return nil
}

func hasWriteStringError(e error) bool {
	return e != nil
}

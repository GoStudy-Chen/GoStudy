package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var array []string

func main() {

	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyODg=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyODk=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTE=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTI=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTM=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTQ=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTU=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTY=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTc=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTg=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEyOTk=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEzMDA=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEzMDE=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTEzMDI=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDM=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDI=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDQ=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDU=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDY=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDc=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDg=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MDk=")
	array = append(array, "MzU2MjA0MTMzOF8xMDAwMTE0MTA=")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MDg4")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTIx")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTIx")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTIz")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTI0")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTI1")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTI2")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTI3")
	array = append(array, "MzU2MjA0MTMzOF8yMjQ3NDk1MTI4")

	url := "https://res.wx.qq.com/voice/getvoice?mediaid="
	for idx, fname := range array {
		if idx >= 23 {
			fileurl := url + fname
			saveName := strconv.Itoa(idx+1) + ".mp3"
			fmt.Println(saveName)
			res, err := http.Get(fileurl)
			if err != nil {
				panic(err)
			}
			f, err := os.Create(saveName)
			if err != nil {
				panic(err)
			}
			io.Copy(f, res.Body)
			fmt.Println(saveName + " is download!")
		}

	}

}

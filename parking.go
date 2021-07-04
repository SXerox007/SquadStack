package main

import (
	constant "SquadStack/base/constants"
	"SquadStack/base/utils"
	"SquadStack/models/common"
	"SquadStack/parkinglot"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// upload parking lot file
func UploadParkingLotFile() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		runType := r.FormValue("type")

		log.Println("Type:", runType)
		if runType == constant.DEFAULT {
			file, err := os.Open("sample.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				if err = file.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			scanner := bufio.NewScanner(file)
			var data []string
			for scanner.Scan() {
				result := parkinglot.PrakingQuery(scanner.Text())
				fmt.Println(result)
				data = append(data, result)
			}
			utils.RespondWithJSON(w, http.StatusOK, common.CommonResponse{
				true,
				"Success",
				http.StatusOK,
				data,
			})

		} else {
			r.ParseMultipartForm(32 << 20)
			file, header, err := r.FormFile("file")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			name := strings.Split(header.Filename, ".")
			//
			if name[1] != "txt" {
				utils.RespondWithJSON(w, http.StatusOK, common.CommonResponse{
					false,
					"Invalid File Type",
					http.StatusBadRequest,
					nil,
				})
				return
			}
			b, err := ioutil.ReadAll(file)
			dataDetails := strings.Split(string(b), "\n")
			var data []string
			for _, item := range dataDetails {
				result := parkinglot.PrakingQuery(item)
				fmt.Println(result)
				data = append(data, result)
			}

			txtContent := strings.Join(data, "\n")
			// read file and send output file
			w.Header().Set("Content-Disposition", "attachment; filename="+header.Filename+"_output.txt")
			w.Header().Set("txt", r.Header.Get("txt"))
			w.Write([]byte(txtContent))
		}
		return
	}
}

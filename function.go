package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func reqDosen(nama, idPt string, f func(Dosen)) error {
	url := BASE_URL + "search_dosen"
	data := fmt.Sprintf(`{
		"nama": "%s",
		"nip": "",
		"pt": "%s",
		"prodi": ""
	}`, nama, idPt)

	var jsonStr = []byte(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resData := struct {
		Dosen []Dosen `json:"dosen"`
	}{
		Dosen: []Dosen{},
	}

	err = json.Unmarshal(body, &resData)
	if err != nil {
		return err
	}

	for _, v := range resData.Dosen {
		f(v)
	}

	return nil
}

func reqDosenDetail(id string, f func(DosenDetail)) error {
	url := BASE_URL + "detail_dosen/" + id

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var resData DosenDetail

	err = json.Unmarshal(body, &resData)
	if err != nil {
		return err
	}

	f(resData)

	return nil
}

func streamJSONFile(fname string, f func(PerguruanTinggi)) error {

	jsonStream, err := os.Open(fname)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(jsonStream)

	_, err = dec.Token()
	if err != nil {
		return err
	}

	for dec.More() {
		var m PerguruanTinggi

		err := dec.Decode(&m)
		if err != nil {
			return err
		}

		f(m)
	}

	_, err = dec.Token()
	if err != nil {
		return err
	}

	return nil
}

package api

//TODO логи для ошибок

import (
	"io"
	"net/http"
	"net/url"

	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
	"github.com/tidwall/gjson"
)

var (
	agifyApi       = "https://api.agify.io/"
	genderizeApi   = "https://api.genderize.io/"
	nationalizeApi = "https://api.nationalize.io/"
)

func FillTheMessage(fio *models.FIO) error {
	err := fillAge(fio)
	if err != nil {
		return err
	}
	fillGender(fio)
	if err != nil {
		return err
	}
	fillNationality(fio)
	if err != nil {
		return err
	}
	return nil
}

func fillAge(fio *models.FIO) error {
	body, err := sendRequest(agifyApi, fio.Name)
	if err != nil {
		return err
	}
	age := gjson.Get(string(body), "age")
	fio.Age = int(age.Num)
	return nil
}

func fillGender(fio *models.FIO) error {
	body, err := sendRequest(genderizeApi, fio.Name)
	if err != nil {
		return err
	}
	gender := gjson.Get(string(body), "gender")
	fio.Gender = gender.String()
	return nil
}

func fillNationality(fio *models.FIO) error {
	body, err := sendRequest(nationalizeApi, fio.Name)
	if err != nil {
		return err
	}
	nationality := gjson.Get(string(body), "country.0.country_id")
	fio.Nationality = nationality.String()
	return nil
}

func sendRequest(requestURL, name string) ([]byte, error) {
	baseURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("name", name)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

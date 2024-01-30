package api

//TODO логи для ошибок

import (
	"io"
	"net/http"
	"net/url"

	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
	"github.com/Scr3amz/EffectiveMobile/logger"
	"github.com/tidwall/gjson"
)

var (
	agifyApi       = "https://api.agify.io/"
	genderizeApi   = "https://api.genderize.io/"
	nationalizeApi = "https://api.nationalize.io/"
)

func FillTheMessage(fio *models.FIO, logger logger.Logger) error {
	err := fillAge(fio, logger)
	if err != nil {
		logger.WarningLog.Println("failed to get age from external api")
		return err
	}
	fillGender(fio, logger)
	if err != nil {
		logger.WarningLog.Println("failed to get gender from external api")
		return err
	}
	fillNationality(fio, logger)
	if err != nil {
		logger.WarningLog.Println("failed to get nationality from external api")
		return err
	}
	return nil
}

func fillAge(fio *models.FIO, logger logger.Logger) error {
	body, err := sendRequest(agifyApi, fio.Name, logger)
	if err != nil {
		return err
	}
	age := gjson.Get(string(body), "age")
	fio.Age = int(age.Num)
	return nil
}

func fillGender(fio *models.FIO, logger logger.Logger) error {
	body, err := sendRequest(genderizeApi, fio.Name, logger)
	if err != nil {
		return err
	}
	gender := gjson.Get(string(body), "gender")
	fio.Gender = gender.String()
	return nil
}

func fillNationality(fio *models.FIO, logger logger.Logger) error {
	body, err := sendRequest(nationalizeApi, fio.Name, logger)
	if err != nil {
		return err
	}
	nationality := gjson.Get(string(body), "country.0.country_id")
	fio.Nationality = nationality.String()
	return nil
}

func sendRequest(requestURL, name string, logger logger.Logger) ([]byte, error) {
	baseURL, err := url.Parse(requestURL)
	if err != nil {
		logger.WarningLog.Println("failed to parse url")
		return nil, err
	}
	params := url.Values{}
	params.Add("name", name)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		logger.WarningLog.Println("failed to send GET request")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.WarningLog.Println("failed to read response body")
		return nil, err
	}
	return body, nil
}

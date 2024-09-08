package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"server/infrastructure"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thedevsaddam/govalidator"
)

type validatorInterface interface {
	ValidateStruct(dataStruct any) (validErrorSlice url.Values, err error)
}

func NewValidation() validatorInterface {
	return &validStruct{}
}

func (v *validStruct) ValidateStruct(dataStruct any) (validErrorSlice url.Values, err error) {
	rv := reflect.ValueOf(dataStruct)
	rt := rv.Type()
	var validRulesSlice []validStruct
	var validMessagesSlice []validStruct
	for i := 0; i < rt.NumField(); i++ {
		var lemper []string
		if value, ok := rt.Field(i).Tag.Lookup("valid"); ok {
			validRulesSlice = append(validRulesSlice, validStruct{
				Key:   rt.Field(i).Tag.Get("json"),
				Value: strings.Split(value, ";"),
			})
			for _, caca := range strings.Split(value, ";") {
				lala := strings.Split(caca, ":")
				lema := lala[0]
				lele := fmt.Sprintf("%s:%s", lema, callMessage(caca))
				lemper = append(lemper, lele)
			}
			validMessagesSlice = append(validMessagesSlice, validStruct{
				Key:   rt.Field(i).Tag.Get("json"),
				Value: lemper,
			})
		}
		if value, ok := rt.Field(i).Tag.Lookup("valid_message"); ok {
			ab := value
			bc := strings.Split(ab, ";")
			for le, mama := range validMessagesSlice {
				if mama.Key == rt.Field(i).Tag.Get("json") {
					var merah []string
					for _, la := range mama.Value {
						ne := strings.Split(la, ":")
						for _, cd := range bc {
							de := strings.Split(cd, ":")
							if ne[0] == de[0] {
								ne[1] = de[1]
							}
						}
						lela := strings.Join(ne, ":")
						merah = append(merah, lela)
					}
					validMessagesSlice[le].Value = merah
				}
			}
		}
	}
	validRulesMap := convertStructToMap(validRulesSlice)
	validMessagesMap := convertStructToMap(validMessagesSlice)
	var validDataMap map[string]interface{}
	data, _ := json.Marshal(dataStruct)
	json.Unmarshal(data, &validDataMap)
	validErrorSlice = generateErrorSlice(validRulesMap, validMessagesMap, validDataMap)
	if len(validErrorSlice) > 0 {
		return validErrorSlice, errors.New("Error on validation")
	}
	return
}

type validStruct struct {
	Key   string
	Value []string
}

func convertStructToMap(validSlice []validStruct) (validMap map[string][]string) {
	validMap = make(map[string][]string)
	for _, val := range validSlice {
		validMap[val.Key] = append(validMap[val.Key], val.Value...)
	}
	return
}

func generateErrorSlice(rules map[string][]string, messages map[string][]string, data map[string]interface{}) (validErrorSlice url.Values) {
	opts := govalidator.Options{
		Data:     &data,
		Rules:    rules,
		Messages: messages,
	}
	v := govalidator.New(opts)
	validErrorSlice = v.ValidateStruct()
	return
}

func callMessage(key string) string {
	clearKeyTemp := strings.Split(key, ":")
	clearKey := clearKeyTemp[0]

	mitos := map[string]string{
		"required": infrastructure.Localize("M_VAL_REQUIRED"),
		"email":    infrastructure.Localize("M_VAL_EMAIL"),
		"min": infrastructure.Localize(&i18n.LocalizeConfig{
			MessageID:    "M_VAL_MIN",
			TemplateData: map[string]string{"min_char": ifKey(key)},
		}),
		"date": infrastructure.Localize(&i18n.LocalizeConfig{
			MessageID:    "M_VAL_DATE",
			TemplateData: map[string]string{"date_format": ifKey(key)},
		}),
		"in": infrastructure.Localize(&i18n.LocalizeConfig{
			MessageID:    "M_VAL_IN",
			TemplateData: map[string]string{"in_char": ifKey(key)},
		}),
		"digits": infrastructure.Localize(&i18n.LocalizeConfig{
			MessageID:    "M_VAL_DIGIT",
			TemplateData: map[string]string{"digits": ifKey(key)},
		}),
		"max": infrastructure.Localize(&i18n.LocalizeConfig{
			MessageID:    "M_VAL_MAX",
			TemplateData: map[string]string{"max_char": ifKey(key)},
		}),
	}
	return mitos[clearKey]
}

func ifKey(data string) (param string) {
	parts := strings.Split(data, ":")
	if len(parts) > 1 {
		param = parts[1]
	}
	return
}

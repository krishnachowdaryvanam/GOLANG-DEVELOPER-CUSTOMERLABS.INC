package worker

import (
	"encoding/json"
	"golang_test_v_1/model"
	"reflect"
	"testing"
)

func TestConvertRequest(t *testing.T) {
	// Input data
	inputData := model.OriginalEvent{
		Ev:     "contact_form_submitted",
		Et:     "form_submit",
		ID:     "cl_app_id_001",
		UID:    "cl_app_id_001-uid-001",
		MID:    "cl_app_id_001-uid-001",
		T:      "Vegefoods - Free Bootstrap 4 Template by Colorlib",
		P:      "http://shielded-eyrie-45679.herokuapp.com/contact-us",
		L:      "en-US",
		SC:     "1920 x 1080",
		Atrk1:  "form_varient",
		Atrv1:  "red_top",
		Atrt1:  "string",
		Atrk2:  "ref",
		Atrv2:  "XPOWJRICW993LKJD",
		Atrt2:  "string",
		UAtrk1: "name",
		UAtrv1: "iron man",
		UAtrt1: "string",
		UAtrk2: "email",
		UAtrv2: "ironman@avengers.com",
		UAtrt2: "string",
		UAtrk3: "age",
		UAtrv3: "32",
		UAtrt3: "integer",
	}

	// Expected output
	expectedData := model.TransformedEvent{
		Event:           "contact_form_submitted",
		EventType:       "form_submit",
		AppID:           "cl_app_id_001",
		UserID:          "cl_app_id_001-uid-001",
		MessageID:       "cl_app_id_001-uid-001",
		PageTitle:       "Vegefoods - Free Bootstrap 4 Template by Colorlib",
		PageURL:         "http://shielded-eyrie-45679.herokuapp.com/contact-us",
		BrowserLanguage: "en-US",
		ScreenSize:      "1920 x 1080",
		Attributes: map[string]model.AttributeValue{
			"form_varient": {
				Value: "red_top",
				Type:  "string",
			},
			"ref": {
				Value: "XPOWJRICW993LKJD",
				Type:  "string",
			},
		},
		Traits: map[string]model.TraitValue{
			"name": {
				Value: "iron man",
				Type:  "string",
			},
			"email": {
				Value: "ironman@avengers.com",
				Type:  "string",
			},
			"age": {
				Value: "32",
				Type:  "integer",
			},
		},
	}

	convertedData := ConvertRequest(inputData)

	var convertedEvent model.TransformedEvent
	err := json.Unmarshal(convertedData, &convertedEvent)
	if err != nil {
		t.Errorf("Error unmarshaling convertedData: %v", err)
	}

	// Compare converted data with expected data
	if !reflect.DeepEqual(convertedEvent, expectedData) {
		t.Errorf("Converted data does not match expected data:\nExpected: %+v\nGot: %+v", expectedData, convertedEvent)
	}
}

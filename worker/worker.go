package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang_test_v_1/model"
	"net/http"
)

var WorkerPool chan model.OriginalEvent

const MaxWorkers = 10

func InitializeWorkerPool() {
	WorkerPool = make(chan model.OriginalEvent)
	for i := 0; i < MaxWorkers; i++ {
		go worker()
	}
}

func worker() {
	for req := range WorkerPool {
		convertedReq := ConvertRequest(req)
		sendToEndpoint(convertedReq)
	}
}

func ConvertRequest(req model.OriginalEvent) []byte {
	convertedReq := model.TransformedEvent{
		Event:           req.Ev,
		EventType:       req.Et,
		AppID:           req.ID,
		UserID:          req.UID,
		MessageID:       req.MID,
		PageTitle:       req.T,
		PageURL:         req.P,
		BrowserLanguage: req.L,
		ScreenSize:      req.SC,
	}

	// Populate attributes
	attributes := make(map[string]model.AttributeValue)
	if req.Atrk1 != "" {
		attributes[req.Atrk1] = model.AttributeValue{
			Value: req.Atrv1,
			Type:  req.Atrt1,
		}
	}
	if req.Atrk2 != "" {
		attributes[req.Atrk2] = model.AttributeValue{
			Value: req.Atrv2,
			Type:  req.Atrt2,
		}
	}
	convertedReq.Attributes = attributes

	// Populate traits
	traits := make(map[string]model.TraitValue)
	if req.UAtrk1 != "" {
		traits[req.UAtrk1] = model.TraitValue{
			Value: req.UAtrv1,
			Type:  req.UAtrt1,
		}
	}
	if req.UAtrk2 != "" {
		traits[req.UAtrk2] = model.TraitValue{
			Value: req.UAtrv2,
			Type:  req.UAtrt2,
		}
	}
	if req.UAtrk3 != "" {
		traits[req.UAtrk3] = model.TraitValue{
			Value: req.UAtrv3,
			Type:  req.UAtrt3,
		}
	}
	convertedReq.Traits = traits

	jsonReq, _ := json.Marshal(convertedReq)
	return jsonReq
}

func sendToEndpoint(req []byte) {
	url := "https://webhook.site/d2b36c5a-f9ba-4350-b644-0711a4ba502c"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println("Error sending request to endpoint:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

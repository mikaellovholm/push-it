package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

//go:embed templates/index.html
var indexHTML string

const placesAPIKey = "AIzaSyCUhREm6d3cBdIpAY3fjdXbq33sJNMI4qk"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(indexHTML))
}

func nearbyHandler(w http.ResponseWriter, r *http.Request) {
	lat, err1 := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, err2 := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	if err1 != nil || err2 != nil {
		http.Error(w, "lat and lng required", http.StatusBadRequest)
		return
	}

	radius := 1000.0
	if r, err := strconv.ParseFloat(r.URL.Query().Get("radius"), 64); err == nil && r >= 250 && r <= 2500 {
		radius = r
	}

	placeType := r.URL.Query().Get("type")
	if placeType != "cafe" {
		placeType = "restaurant"
	}

	body := map[string]any{
		"includedTypes":  []string{placeType},
		"maxResultCount": 20,
		"locationRestriction": map[string]any{
			"circle": map[string]any{
				"center": map[string]float64{"latitude": lat, "longitude": lng},
				"radius": radius,
			},
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://places.googleapis.com/v1/places:searchNearby", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", placesAPIKey)
	req.Header.Set("X-Goog-FieldMask", "places.displayName,places.rating,places.userRatingCount,places.location,places.types,places.currentOpeningHours,places.priceLevel,places.id,places.primaryType")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("places API error: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/nearby", nearbyHandler)
	println("push-it running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

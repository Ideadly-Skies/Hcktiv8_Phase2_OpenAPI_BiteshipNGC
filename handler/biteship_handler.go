package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"w3/ngc/config"
	"w3/ngc/utils"

	"github.com/labstack/echo/v4"
)

func GetCouriers(c echo.Context, cfg *config.Config) error {
    newUrl := "https://api.biteship.com/v1/couriers"

    headers := map[string]string{
        "Authorization": "Bearer " + cfg.BiteshipAPIKey,
    }

    // Perform the GET request
    res, err := utils.RequestGET(newUrl, headers)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to fetch couriers"})
    }

    // Print the raw response for debugging (optional)
    fmt.Println(string(res))

    // Define a struct to match the expected JSON response
    type Courier struct {
        AvailableCollectionMethod        []string `json:"available_collection_method"`
        AvailableForCashOnDelivery       bool     `json:"available_for_cash_on_delivery"`
        AvailableForProofOfDelivery      bool     `json:"available_for_proof_of_delivery"`
        AvailableForInstantWaybillID     bool     `json:"available_for_instant_waybill_id"`
        CourierName                      string   `json:"courier_name"`
        CourierCode                      string   `json:"courier_code"`
        CourierServiceName               string   `json:"courier_service_name"`
        CourierServiceCode               string   `json:"courier_service_code"`
        Tier                             string   `json:"tier"`
        Description                      string   `json:"description"`
        ServiceType                      string   `json:"service_type"`
        ShippingType                     string   `json:"shipping_type"`
        ShipmentDurationRange            string   `json:"shipment_duration_range"`
        ShipmentDurationUnit             string   `json:"shipment_duration_unit"`
    }

    type BiteshipResponse struct {
        Success  bool       `json:"success"`
        Object   string     `json:"object"`
        Couriers []Courier  `json:"couriers"`
    }

    // Unmarshal the response into the struct
    var result BiteshipResponse
    err = json.Unmarshal(res, &result)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "error while unmarshalling couriers"})
    }

    // Return the couriers data
    return c.JSON(http.StatusOK, result.Couriers)
}

func CalculateShippingCost(c echo.Context, cfg *config.Config) error {
	newUrl := "https://api.biteship.com/v1/rates/costs"
	headers := map[string]string{
		"Authorization": "Bearer " + cfg.BiteshipAPIKey,
		"Content-Type":  "application/json",
	}

	origin := c.FormValue("origin")
	destination := c.FormValue("destination")
	weight := c.FormValue("weight")
	courier := c.FormValue("courier")

	payload := fmt.Sprintf(`{
		"origin": "%s",
		"destination": "%s",
		"weight": %s,
		"courier": "%s"
	}`, origin, destination, weight, courier)

	res, err := utils.RequestPOST(newUrl, headers, strings.NewReader(payload))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to calculate shipping cost"})
	}

	var result map[string]interface{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "error while unmarshalling shipping cost"})
	}

	return c.JSON(http.StatusOK, result)
}
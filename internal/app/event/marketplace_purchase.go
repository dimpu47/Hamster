package event

import (
    "encoding/json"
)

// Any time a user purchases, cancels, or changes their GitHub Marketplace plan.
type MarketplacePurchase struct {

}

func (e *MarketplacePurchase) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *MarketplacePurchase) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
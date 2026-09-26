package expenses

import (
	"context"
	"encoding/json"

	"time"

	"github.com/salbadr/smartspends/internal/ai"
	"github.com/salbadr/smartspends/internal/common"
)


func GetExpenses(payload string, aiClient ai.AiModel) (*common.ExpenseResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Ensures resources are freed when the function returns

	res, err := aiClient.GetResponse(ctx, payload)

	if err != nil {
		return nil, err
	}
	resp := common.ExpenseResponse{}

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

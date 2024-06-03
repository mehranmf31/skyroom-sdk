package skyroom

import (
	"encoding/json"
	"fmt"
)

type GetServicesResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Status     int    `json:"status"`
	UserLimit  int    `json:"user_limit"`
	VideoLimit int    `json:"video_limit"`
	TimeLimit  *int   `json:"time_limit"`
	TimeUsage  int    `json:"time_usage"`
	StartTime  int64  `json:"start_time"`
	StopTime   int64  `json:"stop_time"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (sky *Skyroom) GetServices() ([]GetServicesResponse, error) {
	result, err := sky.Post("getServices", make(map[string]string))
	if err != nil {
		fmt.Println(err)
	}

	var services []GetServicesResponse

	if err := json.Unmarshal(result, &services); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return services, nil
}

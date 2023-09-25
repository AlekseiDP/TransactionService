package api

// ResultModel модель для обертки ответа сервиса
type ResultModel struct {
	Data             any    `json:"data,omitempty"`
	ErrorCode        string `json:"errorCode,omitempty"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
	ErrorDisplay     string `json:"errorDisplay,omitempty"`
	ErrorDevelopment string `json:"errorDevelopment,omitempty"`
}

// PageableModel модель для обертки постраничного списка
type PageableModel struct {
	Items            any   `json:"items"`
	CurrentPageIndex int   `json:"currentPageIndex"`
	TotalCount       int64 `json:"totalCount"`
}

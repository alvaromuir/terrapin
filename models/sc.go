package models

// SCpingResponse returns API heartbeat status code
type SCpingResponse struct {
	Status int
}

// SCBookmarks is a collection of SC reporting bookmark folders
type SCBookmarks struct {
	Folders []*SCBookmarkFolderResult
}

// SCBookmarkFolderResult is a collection of SC reporting bookmark folder items
type SCBookmarkFolderResult struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	Bookmarks []*SCBookmarkResult
}

// SCBookmarkResult describes a SC reporting bookmark
type SCBookmarkResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Rsid string `json:"rsid"`
}

// SCRealTimeSettingResult describes the settings for a realtime reporting
type SCRealTimeSettingResult []struct {
	Rsid             string `json:"rsid"`
	SiteTitle        string `json:"site_title"`
	RealTimeSettings []struct {
		Name           string   `json:"name"`
		MinGranularity int      `json:"min_granularity"`
		UIReport       bool     `json:"ui_report"`
		Metric         string   `json:"metric"`
		Elements       []string `json:"elements"`
	} `json:"real_time_settings"`
}

// SCRealtimeMetricsResult descibes a realtime metrics response
type SCRealtimeMetricsResult struct {
	Report struct {
		Data []struct {
			Counts []string `json:"counts"`
			Day    int      `json:"day"`
			Hour   int      `json:"hour"`
			Minute int      `json:"minute"`
			Month  int      `json:"month"`
			Name   string   `json:"name"`
			Year   int      `json:"year"`
		} `json:"data"`
		Elements []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"elements"`
		Metrics []struct {
			Decimals int    `json:"decimals"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			Type     string `json:"type"`
		} `json:"metrics"`
		Period      string `json:"period"`
		ReportSuite struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"reportSuite"`
		Totals  []string `json:"totals"`
		Type    string   `json:"type"`
		Version string   `json:"version"`
	} `json:"report"`
}

// SCMetricsRequest describes a SC metrics data request
type SCMetricsRequest struct {
	ReportDescription struct {
		Source        string            `json:"source"`
		ReportSuiteID string            `json:"reportSuiteID"`
		Metrics       []*SCMetricsArray `json:"metrics"`
	}
}

// SCMetricsArray describes a collection of desired metrics
type SCMetricsArray struct {
	ID string `json:"id"`
}

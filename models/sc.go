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

// SCMetricsRequest describes a SC metrics data request
type SCMetricsRequest struct {
	ReportDescription *SCMetricsReportDescription `json:"reportDescription"`
}

// SCMetricsReportDescription descibes a SC data request
type SCMetricsReportDescription struct {
	Source        string           `json:"source"`
	ReportSuiteID string           `json:"reportSuiteID"`
	Metrics       []SCMetricsArray `json:"metrics"`
}

// SCMetricsArray describes a collection of desired metrics
type SCMetricsArray struct {
	ID string `json:"id"`
}

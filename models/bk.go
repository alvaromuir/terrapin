package models

// BKpingResponse returns API heartbeat status code
type BKpingResponse struct {
	Status int
}

// BKbuyerViewCategoryResult is a collection of OwnerViewCategory types
type BKbuyerViewCategoryResult struct {
	Count        int  `json:"count"`
	HasMore      bool `json:"hasMore"`
	Items        []*BKbuyerViewCategory
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

// BKbuyerViewCategory describes an owner view specific BK category object
type BKbuyerViewCategory struct {
	CategoryType            string        `json:"categoryType"`
	ID                      int           `json:"id"`
	IsForNavigationOnlyFlag bool          `json:"isForNavigationOnlyFlag"`
	IsLeafFlag              bool          `json:"isLeafFlag"`
	IsPublicFlag            bool          `json:"isPublicFlag"`
	Links                   []interface{} `json:"links"`
	Name                    string        `json:"name"`
	OwnershipType           string        `json:"ownershipType"`
	ParentCategory          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"parentCategory"`
	Partner struct {
		ID int `json:"id"`
	} `json:"partner"`
	PathFromRoot struct {
		Ids   []int    `json:"ids"`
		Names []string `json:"names"`
	} `json:"pathFromRoot"`
	PriceFloor float32 `json:"priceFloor"`
	Stats      struct {
		Reach int `json:"reach"`
	} `json:"stats"`
	Status   string `json:"status"`
	Vertical struct {
		Name string `json:"name"`
	} `json:"vertical"`
}

/* BKownerViewCategoryResult Not in use yet, @alvaromuir 04.10.2016 */

// BKownerViewCategoryResult is a collection of BuyerViewCategory types
type BKownerViewCategoryResult struct {
	Count        int  `json:"count"`
	HasMore      bool `json:"hasMore"`
	Items        []*BKownerViewCategory
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

// BKownerViewCategory describes a buyer view specific BK category object
type BKownerViewCategory struct {
	CategoryType                    string        `json:"categoryType"`
	Description                     string        `json:"description"`
	ID                              int           `json:"id"`
	IsCountableFlag                 bool          `json:"isCountableFlag"`
	IsExplicitPriceFloorFlag        bool          `json:"isExplicitPriceFloorFlag"`
	IsForNavigationOnlyFlag         bool          `json:"isForNavigationOnlyFlag"`
	IsIncludeForAnalyticsFlag       bool          `json:"isIncludeForAnalyticsFlag"`
	IsLeafFlag                      bool          `json:"isLeafFlag"`
	IsMutuallyExclusiveChildrenFlag bool          `json:"isMutuallyExclusiveChildrenFlag"`
	IsPublicFlag                    bool          `json:"isPublicFlag"`
	Links                           []interface{} `json:"links"`
	Name                            string        `json:"name"`
	NamespaceID                     int           `json:"namespaceId"`
	OwnershipType                   string        `json:"ownershipType"`
	ParentCategory                  struct {
		ID int `json:"id"`
	} `json:"parentCategory"`
	Partner struct {
		ID int `json:"id"`
	} `json:"partner"`
	PathFromRoot struct {
		Ids   []int    `json:"ids"`
		Names []string `json:"names"`
	} `json:"pathFromRoot"`
	PriceFloor float32 `json:"priceFloor"`
	SoftFloor  int     `json:"softFloor"`
	SortOrder  int     `json:"sortOrder"`
	Status     string  `json:"status"`
	Vertical   struct {
		Name string `json:"name"`
	} `json:"vertical"`
	VisibilityStatus string `json:"visibilityStatus"`
}

// BKsiteResult is a collection of Site types
type BKsiteResult struct {
	TotalCount int `json:"total_count"`
	Sites      []*BKsite
}

// BKsite describes a BK site object
type BKsite struct {
	AllowedBuyers []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"allowed_buyers"`
	AnalyticsOnly             string      `json:"analytics_only"` // REVIEW: should be boo
	BlockedCountries          []string    `json:"blocked_countries"`
	CreatedAt                 string      `json:"created_at"`                  // REVIEW: should be time
	UpdatedAt                 string      `json:"updated_at"`                  // REVIEW: should be time
	DataTransferBoostEnabled  string      `json:"data_transfer_boost_enabled"` // REVIEW: should be bool
	DataTransferBoostInterval int         `json:"data_transfer_boost_interval"`
	DataTransferLimit         int         `json:"data_transfer_limit"`
	GroupID                   int         `json:"group_id"`
	ID                        int         `json:"id"`
	Labels                    interface{} `json:"labels"`
	Name                      string      `json:"name"`
	Notes                     string      `json:"notes"`
	Partner                   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"partner"`
	PrivateData      string `json:"private_data"`
	TransactionScope string `json:"transaction_scope"`
	Type             int    `json:"type"`
}

package models

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

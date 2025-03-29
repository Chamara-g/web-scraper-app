package models

type SiteData struct {
	URL             string              `json:"url"`
	TITLE           string              `json:"title"`
	HTML_VERSION    string              `json:"html_version"`
	HEADING_LEVELS  map[string]int      `json:"heading_levels"`
	HAVE_LOGIN_FORM bool                `json:"have_login_form"`
	LINKS           map[string][]string `json:"links"`
}

// func NewSiteData(url, title string) SiteData {
// 	return SiteData{URL: url, Title: title}
// }
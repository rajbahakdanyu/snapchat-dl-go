package cmd

type SnapchatData struct {
	Props                 Props  `json:"props"`
	Page                  string `json:"page"`
	Query                 Query  `json:"query"`
	BuildID               string `json:"buildId"`
	AssetPrefix           string `json:"assetPrefix"`
	IsFallback            bool   `json:"isFallback"`
	IsExperimentalCompile bool   `json:"isExperimentalCompile"`
	Gip                   bool   `json:"gip"`
	AppGip                bool   `json:"appGip"`
	ScriptLoader          []any  `json:"scriptLoader"`
}
type ViewerInfo struct {
	Country       string `json:"country"`
	Locale        string `json:"locale"`
	IsGdprCountry bool   `json:"isGdprCountry"`
}
type Localization struct {
	Direction int `json:"direction"`
}
type OneLinkParams struct {
	OneLinkBaseURL        string   `json:"oneLinkBaseUrl"`
	PidKeys               []string `json:"pidKeys"`
	PidFallbackValue      string   `json:"pidFallbackValue"`
	CampaignKeys          []string `json:"campaignKeys"`
	CampaignFallbackValue string   `json:"campaignFallbackValue"`
	GoogleClickIDParam    string   `json:"googleClickIdParam"`
	DeepLinkURL           string   `json:"deepLinkUrl"`
	IosAppStoreURL        string   `json:"iosAppStoreUrl"`
	DesktopPageURL        string   `json:"desktopPageUrl"`
}
type PageLinks struct {
	OneLinkURL           string        `json:"oneLinkUrl"`
	SnapchatCanonicalURL string        `json:"snapchatCanonicalUrl"`
	CanonicalURL         string        `json:"canonicalUrl"`
	SnapcodeImageURL     string        `json:"snapcodeImageUrl"`
	DownloadURL          string        `json:"downloadUrl"`
	OneLinkParams        OneLinkParams `json:"oneLinkParams"`
	SnapchatDeepLinkURL  string        `json:"snapchatDeepLinkUrl"`
}
type PageDescription struct {
	Value string `json:"value"`
}
type PageMetadata struct {
	PageType        int             `json:"pageType"`
	PageTitle       string          `json:"pageTitle"`
	PageDescription PageDescription `json:"pageDescription"`
	ShareID         string          `json:"shareId"`
}
type PublicProfileInfo struct {
	Username               string `json:"username"`
	Title                  string `json:"title"`
	SnapcodeImageURL       string `json:"snapcodeImageUrl"`
	Badge                  int    `json:"badge"`
	CategoryStringID       string `json:"categoryStringId"`
	SubcategoryStringID    string `json:"subcategoryStringId"`
	SubscriberCount        string `json:"subscriberCount"`
	Bio                    string `json:"bio"`
	WebsiteURL             string `json:"websiteUrl"`
	ProfilePictureURL      string `json:"profilePictureUrl"`
	Address                string `json:"address"`
	HasCuratedHighlights   bool   `json:"hasCuratedHighlights"`
	HasSpotlightHighlights bool   `json:"hasSpotlightHighlights"`
	MutableName            string `json:"mutableName"`
	PublisherType          string `json:"publisherType"`
	SquareHeroImageURL     string `json:"squareHeroImageUrl"`
	PrimaryColor           string `json:"primaryColor"`
	HasStory               bool   `json:"hasStory"`
	RelatedAccountsInfo    []any  `json:"relatedAccountsInfo"`
}
type UserProfile struct {
	Case              string            `json:"$case"`
	PublicProfileInfo PublicProfileInfo `json:"publicProfileInfo"`
}
type SnapID struct {
	Value string `json:"value"`
}
type MediaPreviewURL struct {
	Value string `json:"value"`
}
type SnapUrls struct {
	MediaURL        string          `json:"mediaUrl"`
	MediaPreviewURL MediaPreviewURL `json:"mediaPreviewUrl"`
}
type TimestampInSec struct {
	Value string `json:"value"`
}
type SnapList struct {
	SnapIndex      int            `json:"snapIndex"`
	SnapID         SnapID         `json:"snapId"`
	SnapMediaType  int            `json:"snapMediaType"`
	SnapUrls       SnapUrls       `json:"snapUrls"`
	TimestampInSec TimestampInSec `json:"timestampInSec"`
}
type StoryID struct {
	Value string `json:"value"`
}
type ThumbnailURL struct {
	Value string `json:"value"`
}
type Story struct {
	StoryType    int          `json:"storyType"`
	SnapList     []SnapList   `json:"snapList"`
	StoryID      StoryID      `json:"storyId"`
	ThumbnailURL ThumbnailURL `json:"thumbnailUrl"`
	StoryTapID   string       `json:"storyTapId"`
}
type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type TwitterImage struct {
	URL  string `json:"url"`
	Size Size   `json:"size"`
}
type FacebookImage struct {
	URL  string `json:"url"`
	Size Size   `json:"size"`
}
type LinkPreview struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	CanonicalURL  string        `json:"canonicalUrl"`
	TwitterImage  TwitterImage  `json:"twitterImage"`
	FacebookImage FacebookImage `json:"facebookImage"`
}
type StoryTitle struct {
	Value string `json:"value"`
}
type HighlightID struct {
	Value string `json:"value"`
}
type CuratedHighlights struct {
	StoryType    int          `json:"storyType"`
	SnapList     []SnapList   `json:"snapList"`
	StoryID      StoryID      `json:"storyId"`
	StoryTitle   StoryTitle   `json:"storyTitle"`
	ThumbnailURL ThumbnailURL `json:"thumbnailUrl"`
	StoryTapID   string       `json:"storyTapId"`
	HighlightID  HighlightID  `json:"highlightId"`
}
type SpotlightHighlights struct {
	StoryType    int          `json:"storyType"`
	SnapList     []SnapList   `json:"snapList"`
	StoryID      StoryID      `json:"storyId"`
	StoryTitle   StoryTitle   `json:"storyTitle"`
	ThumbnailURL ThumbnailURL `json:"thumbnailUrl"`
	StoryTapID   string       `json:"storyTapId"`
	HighlightID  HighlightID  `json:"highlightId"`
}
type PersonCreator struct {
	Username string `json:"username"`
	URL      string `json:"url"`
	Name     string `json:"name"`
}
type Creator struct {
	Case          string        `json:"$case"`
	PersonCreator PersonCreator `json:"personCreator"`
}
type VideoMetadata struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ThumbnailURL string  `json:"thumbnailUrl"`
	UploadDateMs string  `json:"uploadDateMs"`
	ViewCount    string  `json:"viewCount"`
	ContentURL   string  `json:"contentUrl"`
	Creator      Creator `json:"creator"`
	DurationMs   string  `json:"durationMs"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	Keywords     []any   `json:"keywords"`
	ShareCount   string  `json:"shareCount"`
}
type ContextCards struct {
	ContextType   int    `json:"contextType"`
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	URL           string `json:"url"`
	ThumbnailURL  string `json:"thumbnailUrl"`
	SnapcodeURL   string `json:"snapcodeUrl"`
	ThumbnailType int    `json:"thumbnailType"`
	HasBadge      bool   `json:"hasBadge"`
}
type EngagementStats struct {
	ViewCount  string `json:"viewCount"`
	ShareCount string `json:"shareCount"`
}
type SpotlightStoryMetadata struct {
	VideoMetadata   VideoMetadata   `json:"videoMetadata"`
	Hashtags        []string        `json:"hashtags"`
	ContextCards    []ContextCards  `json:"contextCards"`
	EngagementStats EngagementStats `json:"engagementStats"`
	Deeplink        string          `json:"deeplink"`
	OneLinkParams   OneLinkParams   `json:"oneLinkParams"`
	Description     string          `json:"description"`
}
type Gaid struct {
	UAID          string `json:"UAId"`
	MeasurementID string `json:"measurementId"`
}
type PageProps struct {
	ViewerInfo                ViewerInfo               `json:"viewerInfo"`
	Localization              Localization             `json:"localization"`
	PageLinks                 PageLinks                `json:"pageLinks"`
	PageMetadata              PageMetadata             `json:"pageMetadata"`
	UserProfile               UserProfile              `json:"userProfile"`
	Story                     Story                    `json:"story"`
	LinkPreview               LinkPreview              `json:"linkPreview"`
	Lenses                    []any                    `json:"lenses"`
	CuratedHighlights         []CuratedHighlights      `json:"curatedHighlights"`
	SpotlightHighlights       []SpotlightHighlights    `json:"spotlightHighlights"`
	ShowSnapExpiredToast      bool                     `json:"showSnapExpiredToast"`
	LensCursor                string                   `json:"lensCursor"`
	CuratedHighlightsCursor   string                   `json:"curatedHighlightsCursor"`
	SpotlightHighlightsCursor string                   `json:"spotlightHighlightsCursor"`
	SpotlightStoryMetadata    []SpotlightStoryMetadata `json:"spotlightStoryMetadata"`
	Gaid                      Gaid                     `json:"gaid"`
}
type Props struct {
	PageProps  PageProps `json:"pageProps"`
	Status     int       `json:"status"`
	DeviceType string    `json:"deviceType"`
}
type Query struct {
	ProfileParams []string `json:"profileParams"`
}

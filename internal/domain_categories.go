package internal

// Simple categories represent common website types with predefined categories
var DomainCategories = map[string][]string{
	// Social Media
	"facebook.com":  {"social", "messaging", "advertising"},
	"instagram.com": {"social", "photo-sharing", "advertising"},
	"twitter.com":   {"social", "microblogging", "news"},
	"linkedin.com":  {"social", "professional", "jobs"},
	"tiktok.com":    {"social", "video", "entertainment"},
	"pinterest.com": {"social", "photo-sharing", "lifestyle"},
	"reddit.com":    {"social", "forum", "news"},

	// Technology
	"google.com":        {"search", "advertising", "technology"},
	"youtube.com":       {"video", "streaming", "social"},
	"microsoft.com":     {"technology", "software", "cloud"},
	"apple.com":         {"technology", "retail", "hardware"},
	"amazon.com":        {"ecommerce", "retail", "cloud"},
	"github.com":        {"technology", "development", "collaboration"},
	"stackoverflow.com": {"technology", "qa", "development"},

	// AI/ML
	"openai.com":     {"ai", "technology", "development"},
	"claude.ai":      {"ai", "chatbot", "productivity"},
	"anthropic.com":  {"ai", "technology", "research"},
	"huggingface.co": {"ai", "development", "machine-learning"},

	// News/Media
	"cnn.com":       {"news", "media", "politics"},
	"bbc.com":       {"news", "media", "broadcast"},
	"nytimes.com":   {"news", "media", "journalism"},
	"reuters.com":   {"news", "finance", "journalism"},
	"bloomberg.com": {"news", "finance", "business"},

	// Education
	"coursera.org":     {"education", "online-learning", "professional"},
	"udemy.com":        {"education", "online-learning", "technology"},
	"edx.org":          {"education", "online-learning", "academic"},
	"khan-academy.org": {"education", "online-learning", "non-profit"},

	// Entertainment
	"netflix.com": {"streaming", "entertainment", "video"},
	"spotify.com": {"streaming", "music", "entertainment"},
	"disney.com":  {"entertainment", "streaming", "media"},
	"twitch.tv":   {"streaming", "gaming", "entertainment"},

	// Business/Professional
	"salesforce.com": {"business", "crm", "cloud"},
	"zoom.us":        {"business", "communication", "video-conferencing"},
	"slack.com":      {"business", "communication", "collaboration"},
	"atlassian.com":  {"business", "software", "collaboration"},

	// Financial
	"paypal.com":     {"financial", "payment", "ecommerce"},
	"visa.com":       {"financial", "payment", "banking"},
	"mastercard.com": {"financial", "payment", "banking"},
	"stripe.com":     {"financial", "payment", "technology"},

	// Cloud Providers
	"aws.amazon.com":   {"cloud", "technology", "hosting"},
	"azure.com":        {"cloud", "technology", "hosting"},
	"cloud.google.com": {"cloud", "technology", "hosting"},

	// Productivity
	"office.com":  {"productivity", "software", "collaboration"},
	"dropbox.com": {"cloud-storage", "productivity", "collaboration"},
	"notion.so":   {"productivity", "collaboration", "notes"},

	// Shopping
	"ebay.com":    {"ecommerce", "retail", "auction"},
	"walmart.com": {"retail", "ecommerce", "shopping"},
	"etsy.com":    {"ecommerce", "handmade", "marketplace"},

	// Travel
	"booking.com": {"travel", "hospitality", "booking"},
	"airbnb.com":  {"travel", "hospitality", "marketplace"},
	"expedia.com": {"travel", "booking", "flights"},

	// Sports
	"espn.com": {"sports", "news", "entertainment"},
	"nba.com":  {"sports", "basketball", "entertainment"},
	"fifa.com": {"sports", "football", "organization"},
}

// MaliciousDomains represents known malicious domains
var MaliciousDomains = map[string]bool{
	"malware-example.com":         true,
	"phishing-example.net":        true,
	"spam-distribution.example":   true,
	"fake-bank-example.com":       true,
	"malicious-downloads.example": true,
	"credential-theft.example":    true,
	"ransomware-domain.example":   true,
	"botnet-cc.example":           true,
	"exploit-kit.example":         true,
	"scam-website.example":        true,
}

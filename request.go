package alexa

// constants

// built in intents
const (
	//HelpIntent is the Alexa built-in Help Intent
	HelpIntent = "AMAZON.HelpIntent"

	//CancelIntent is the Alexa built-in Cancel Intent
	CancelIntent = "AMAZON.CancelIntent"

	//StopIntent is the Alexa built-in Stop Intent
	StopIntent = "AMAZON.StopIntent"
)

// locales
const (
	// LocaleItalian is the locale for Italian
	LocaleItalian = "it-IT"

	// LocaleGerman is the locale for standard dialect German
	LocaleGerman = "de-DE"

	// LocaleAustralianEnglish is the locale for Australian English
	LocaleAustralianEnglish = "en-AU"

	//LocaleCanadianEnglish is the locale for Canadian English
	LocaleCanadianEnglish = "en-CA"

	//LocaleBritishEnglish is the locale for UK English
	LocaleBritishEnglish = "en-GB"

	//LocaleIndianEnglish is the locale for Indian English
	LocaleIndianEnglish = "en-IN"

	//LocaleAmericanEnglish is the locale for American English
	LocaleAmericanEnglish = "en-US"

	// LocaleJapanese is the locale for Japanese
	LocaleJapanese = "ja-JP"
)

func IsEnglish(locale string) bool {
	switch locale {
	case LocaleAmericanEnglish:
		return true
	case LocaleIndianEnglish:
		return true
	case LocaleBritishEnglish:
		return true
	case LocaleCanadianEnglish:
		return true
	case LocaleAustralianEnglish:
		return true
	default:
		return false
	}
}

// request

// Request is an Alexa skill request
// see https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#request-format
type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Body    ReqBody `json:"request"`
	Context Context `json:"context"`
}

// Session represents the Alexa skill session
type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

// Context represents the Alexa skill request context
type Context struct {
	System struct {
		APIAccessToken string `json:"apiAccessToken"`
		APIEndpoint string `json:"apiEndpoint"`
                Device         struct {
			DeviceID string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
	} `json:"System,omitempty"`
}

// ReqBody is the actual request information
type ReqBody struct {
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
	Intent    Intent `json:"intent,omitempty"`
	Reason    string `json:"reason,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
}

// Intent is the Alexa skill intent
type Intent struct {
	Name  string          `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

// Slot is an Alexa skill slot
type Slot struct {
	Name  			  string `json:"name"`
	Value 			  string `json:"value"`
	Resolutions  Resolutions `json:"resolutions"`
}

type Resolutions struct {
	ResolutionPerAuthority []struct{
		Values []struct{
			Value struct{
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}

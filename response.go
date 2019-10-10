package alexa

//NewSimpleResponse builds a session response
func NewSimpleResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "PlainText",
				Text: text,
			},
			Card: &Payload{
				Type:    "Simple",
				Title:   title,
				Content: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

// Response Types

// Response is the response back to the Alexa speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResBody                `json:"response"`
}

// ResBody is the actual body of the response
type ResBody struct {
	OutputSpeech     *Payload  `json:"outputSpeech,omitempty"`
	Card             *Payload  `json:"card,omitempty"`
	Reprompt         *Reprompt `json:"reprompt,omitempty"`
	Directives       []Directives `json:"directives,omitempty"`
	ShouldEndSession bool      `json:"shouldEndSession"`
}

// Reprompt is imformation
type Reprompt struct {
	OutputSpeech Payload `json:"outputSpeech,omitempty"`
}

// Directives is imformation
type Directives struct {
	Type          string         `json:"type,omitempty"`
	SlotToElicit  string         `json:"slotToElicit,omitempty"`
	UpdatedIntent *UpdatedIntent `json:"updatedIntent,omitempty"`
	PlayBehavior  string         `json:"playBehavior,omitempty"`
	AudioItem     struct {
		Stream struct {
			Token                string `json:"token,omitempty"`
			URL                  string `json:"url,omitempty"`
			OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
		} `json:"stream,omitempty"`
	} `json:"audioItem,omitempty"`
	Template interface{} `json:"template,omitempty"`
}

type BodyTemplate struct {
	Type string `json:"type"`
	Token string `json:"token,omitempty"`
	BackButton string `json:"backButton"`
	BackgroundImage *ImageTemplate `json:"backgroundImage,omitempty"`
	Image *ImageTemplate `json:"image,omitempty"`
	TextContext *TextContext `json:"textContent"`
}

type ListTemplate struct {
	Type string `json:"type"`
	Token string `json:"token,omitempty"`
	BackButton string `json:"backButton"`
	BackgroundImage *ImageTemplate `json:"backgroundImage,omitempty"`
	Title string `json:"title"`
	ListItems *[]ListItems `json:"listItems"`
}

type ListItems struct {
	Token string `json:"token,omitempty"`
	Image *ImageTemplate `json:"image,omitempty"`
	TextContent *TextContext `json:"textContent"`
}

type ImageTemplate struct {
	ContentDescription string `json:"contentDescription,omitempty"`
	Sources *[]ImageTemplateSources `json:"sources,omitempty"`
}

type ImageTemplateSources struct {
	Url string `json:"url"`
	Size string `json:"size,omitempty"`
	WidthPixels int `json:"widthPixels,omitempty"`
	HeightPixels int `json:"heightPixels,omitempty"`
}

type TextContext struct {
	PrimaryText *TextField `json:"primaryText,omitempty"`
	SecondaryText *TextField `json:"secondaryText,omitempty"`
	TertiaryText *TextField `json:"tertiaryText,omitempty"`
}

type TextField struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// UpdatedIntent is to update the Intent
type UpdatedIntent struct {
	Name               string                 `json:"name,omitempty"`
	ConfirmationStatus string                 `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}

// Image ...
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// Payload ...
type Payload struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Text    string `json:"text,omitempty"`
	SSML    string `json:"ssml,omitempty"`
	Content string `json:"content,omitempty"`
	Image   Image  `json:"image,omitempty"`
}

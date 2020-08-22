package alexa

import (
	"strings"
)

// String parse utility.
func ParseString(text string) string {
	text = strings.ToLower(text)
	text = strings.Replace(text, "&", "and", -1)
	text = strings.Replace(text, "+", "plus", -1)
	text = strings.Replace(text, "@", "at", -1)
	text = strings.Replace(text, "w/", "with", -1)
	text = strings.Replace(text, "in.", "inches", -1)
	text = strings.Replace(text, "s/h", "shipping and handling", -1)
	text = strings.Replace(text, " ac ", " after coupon ", -1)
	text = strings.Replace(text, "fs", "free shipping", -1)
	text = strings.Replace(text, "f/s", "free shipping", -1)
	text = strings.Replace(text, "-", "", -1)
	text = strings.Replace(text, "™", "", -1)
	text = strings.Replace(text, "  ", " ", -1)
	return text
}

// Inject new text in builder
func (builder *SSMLBuilder) Say(text string) {
	text = ParseString(text)
	builder.SSML = append(builder.SSML, SSML{text: text})
}

// Inject pauses in builder
func (builder *SSMLBuilder) Pause(pause string) {
	builder.SSML = append(builder.SSML, SSML{pause: pause})
}

// Build SSML response text
func (builder *SSMLBuilder) Build() string {
	var response string
	for index, ssml := range builder.SSML {
		if ssml.text != "" {
			response += ssml.text + " "
		} else if ssml.pause != "" && index != len(builder.SSML)-1 {
			response += "<break time='" + ssml.pause + "ms'/> "
		}
	}
	return "<speak>" + response + "</speak>"
}

//NewSSMLResponse needs builder for more advanced responses
func NewSSMLResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

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
	OutputSpeech     *Payload     `json:"outputSpeech,omitempty"`
	Card             *Payload     `json:"card,omitempty"`
	Reprompt         *Reprompt    `json:"reprompt,omitempty"`
	Directives       []Directives `json:"directives,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

// Reprompt is imformation
type Reprompt struct {
	OutputSpeech Payload `json:"outputSpeech,omitempty"`
}

// Directives is imformation
type Directives struct {
	Type          string         `json:"type,omitempty"`
	SlotToElicit  string         `json:"slotToElicit,omitempty"`
	UpdatedIntent *UpdatedIntent `json:"UpdatedIntent,omitempty"`
	PlayBehavior  string         `json:"playBehavior,omitempty"`
	AudioItem     struct {
		Stream struct {
			Token                string `json:"token,omitempty"`
			URL                  string `json:"url,omitempty"`
			OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
		} `json:"stream,omitempty"`
	} `json:"audioItem,omitempty"`
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

// SSML ...
type SSML struct {
	text  string
	pause string
}

// SSMLBuilder ...
type SSMLBuilder struct {
	SSML []SSML
}

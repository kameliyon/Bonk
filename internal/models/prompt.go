package models

// {
// 	"role": "system", 
// 	"content": "You are a helpful assistant." 
// },
// {
// 	"role": "user", 
// 	"content": [
// 		{
// 			"type": "text",
// 			"text": "Describe this picture:"
// 		},
// 		{
// 			"type": "image_url",
// 			"image_url": {
// 				"url": "<image URL>"
// 			}
// 		} 
//    ] 
// }

type Prompt struct {
	Role string `json:"role"`
	Content []ContentType `json:"content"`
}

type ContentType interface {
	GetType() string
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ImageContent struct {
	Type string `json:"type"`
	ImageUrl Image `json:"image_url"`
}

type Image struct {
	Url string `json:"url"`
}

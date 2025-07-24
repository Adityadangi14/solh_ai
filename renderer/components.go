package renderer

import (
	"encoding/json"
	"fmt"

	"github.com/Adityadangi14/solh_ai/appmodels"
	"github.com/Adityadangi14/solh_ai/db"
)

func ComponentRenderer(items []string) ([]map[string]any, error) {

	components := make([]map[string]any, 0)
	for _, item := range items {
		res, err := db.GetUrlObject(item)

		if err != nil {
			fmt.Println(err)
		}

		data := res.Data

		byt, err := json.Marshal(data)

		if err != nil {
			return nil, err
		}

		var contentModel appmodels.ContentModel

		err = json.Unmarshal(byt, &contentModel)

		if err != nil {
			return nil, err
		}

		comp := RenderTypeController(contentModel)

		components = append(components, comp)
	}

	return components, nil
}

func RenderTypeController(obj appmodels.ContentModel) map[string]any {
	switch obj.Get.Content[0].ContentType {
	case "blog":
		return RenderBlog(obj)

	case "audio":
		return RenderAudio(obj)
	default:
		return map[string]any{}
	}
}

func RenderBlog(model appmodels.ContentModel) map[string]any {
	str := `
	    {
  "type": "container",
  "decoration": {
    "color": "#FFFFFF",
    "borderRadius": {
      "topLeft": 16.0,
      "topRight": 16.0,
      "bottomLeft": 16.0,
      "bottomRight": 16.0
    },
    "boxShadow": [
      {
        "color": "#22000000",
        "blurRadius": 8.0,
        "offset": {"dx": 0, "dy": 2}
      }
    ]
  },
  "child": {
    "type": "padding",
    "padding": { "all": 12.0 },
    "child": {
      "type": "row",
      "crossAxisAlignment": "start",
      "children": [
        {
          "type": "image",
          "src": "%v",
          "imageType": "network",
          "width": 80.0,
          "height": 80.0,
          "fit": "cover"
        },
        {
          "type": "sizedBox",
          "width": 12.0
        },
        {
          "type": "expanded",
          "child": {
            "type": "column",
            "crossAxisAlignment": "start",
            "children": [
             {
            "type": "padding",
                "padding": { "right": 12.0 },
                "child": {
                "type": "text",
                "maxLines": 3,
                "data": "%v",
           
                "style": {
                  "fontSize": 14.0,
                  "fontWeight": "w700",
                  "color": "#222222"
                }
              }
              },
               {
                "type": "sizedBox",
                "height": 4.0
              },
              {
                "type": "text",
                "data": "Blog",
                "style": {
                  "fontSize": 12.0,
                  "fontWeight": "w500",
                  "color": "#777777"
                }
              }         
            ]
          }
        }
      ]
    }
  }
}`

	blogMap := map[string]any{
		"type":    "blog",
		"route":   "/blogs",
		"data":    model.Get.Content[0].URL,
		"content": fmt.Sprintf(str, model.Get.Content[0].Image, model.Get.Content[0].Title),
	}
	return blogMap
}

func RenderAudio(model appmodels.ContentModel) map[string]any {
	str := `
 {
        "type":  "container",
        "height":  80,
        "decoration":  {
            "color":  "#FFFFFF",
            "borderRadius":  {
                "topLeft":  16,
                "topRight":  16,
                "bottomLeft":  16,
                "bottomRight":  16
            },
            "boxShadow":  [
                {
                    "color":  "#22000000",
                    "blurRadius":  8,
                    "offset":  {
                        "dx":  0,
                        "dy":  2
                    }
                }
            ]
        },
        "child":  {
            "type":  "padding",
            "padding":  {
                "all":  12
            },
            "child":  {
                "type":  "row",
                "crossAxisAlignment":  "center",
                "children":  [
                   {
                        "type":  "sizedBox",
                        "width":  8
                    },
                    {
                        "type":  "container",
                        "width":  60,
                        "height":  60,
                        "clipBehavior":  "hardEdge",
                        "decoration":  {
                            "borderRadius":  {
                                "topLeft":  30,
                                "topRight":  30,
                                "bottomLeft":  30,
                                "bottomRight":  30
                            },
                            "color":  "#F0F0F0"
                        },
                        "child":  {
                            "type":  "padding",
                            "padding":  {
                                "all":  0
                            },
                            "child":  {
                                "type":  "image",
                                "src":  "%v",
                                "imageType":  "network",
                                "fit":  "cover"
                            }
                        }
                    },
                    {
                        "type":  "sizedBox",
                        "width":  12
                    },
                    {
                        "type":  "expanded",
                        "child":  {
                            "type":  "column",
                            "crossAxisAlignment":  "start",
                            "children":  [
                                {
                                    "type":  "padding",
                                    "padding":  {
                                        "right":  12
                                    },
                                    "child":  {
                                        "type":  "text",
                                        "maxLines":  2,
                                        "data":  "%v",
                                        "style":  {
                                            "fontSize":  14,
                                            "fontWeight":  "w700",
                                            "color":  "#222222"
                                        }
                                    }
                                },
                                {
                                    "type":  "sizedBox",
                                    "height":  4
                                },
                                {
                                    "type":  "text",
                                    "data":  "Audio",
                                    "style":  {
                                        "fontSize":  12,
                                        "fontWeight":  "w500",
                                        "color":  "#777777"
                                    }
                                }
                            ]
                        }
                    },
                    {
                        "type":  "sizedBox",
                        "width":  8
                    },
                    {
                        "type":  "icon",
                        "icon":  "play_circle_filled",
                        "size":  28,
                        "color":  "#1E88E5"
                    },
                    {
                        "type":  "sizedBox",
                        "width":  8
                    }
                ]
            }
        }
    }
  `

	audioMap := map[string]any{
		"type":    "audio",
		"route":   "/audio",
		"data":    model.Get.Content[0].URL,
		"content": fmt.Sprintf(str, model.Get.Content[0].Image, model.Get.Content[0].Title),
	}
	return audioMap
}
func RenderVideo() map[string]any {
	return map[string]any{}
}

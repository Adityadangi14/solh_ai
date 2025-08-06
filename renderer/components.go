package renderer

import (
	"encoding/json"
	"fmt"

	"github.com/Adityadangi14/solh_ai/appmodels"
	"github.com/Adityadangi14/solh_ai/db"
	"github.com/Adityadangi14/solh_ai/initializers"
)

func ComponentRenderer(items []string) ([]map[string]any, error) {

	components := make([]map[string]any, 0)
	for _, item := range items {
		res, err := db.GetUrlObject(item)
		initializers.AppLogger.Info("objects", "GetUrlObject", res)
		if err != nil {
			fmt.Println(err)
			initializers.AppLogger.Error("failded to get Url objects", "GetUrlObject", err)
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

		initializers.AppLogger.Info("contentModel", " contentModel", contentModel.Get.Content)

		comp := RenderTypeController(contentModel)

		components = append(components, comp)

		initializers.AppLogger.Info("components", "components", components)
	}

	return components, nil
}

func RenderTypeController(obj appmodels.ContentModel) map[string]any {
	switch obj.Get.Content[0].ContentType {
	case "Blog":
		return RenderBlog(obj)

	case "audio":
		return RenderAudio(obj)

	case "pdf":
		return RenderPdf(obj)

	case "Video":
		return RenderVideo(obj)

	default:
		return map[string]any{}
	}
}

func RenderPdf(model appmodels.ContentModel) map[string]any {
	str := `
	    {
  "type": "container",
  "height": 200,
  "width": 160,
  "decoration": {
    "color": "#FFFFFF",
    "borderRadius": {
      "topLeft": 10,
      "topRight": 10,
      "bottomLeft": 10,
      "bottomRight": 10
    },
    "boxShadow": [
      {
        "color": "#22000000",
        "blurRadius": 8,
        "offset": {"dx": 0, "dy": 2}
      }
    ]
  },
  "child": {
    "type": "row",
    "mainAxisSize": "min",
    "children": [
      {"type": "sizedBox", "width": 5},
      {
        "type": "sizedBox",
        "width": 150,
        "child": {
          "type": "column",
          "crossAxisAlignment": "start",
          "children": [
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "height": 90,
              "clipBehavior": "hardEdge",
              "decoration": {
                "color": "#F0F0F0",
                "borderRadius": {
                  "topLeft": 10,
                  "topRight": 10,
                  "bottomLeft": 10,
                  "bottomRight": 10
                }
              },
              "child": {
                "type": "image",
                "imageType": "network",
                "src":
                    "%v",
                "fit": "cover",
                "width": 150,
                "height": 200
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "decoration": {
                "color": "#2196F3",
                "borderRadius": {
                  "topLeft": 8,
                  "topRight": 8,
                  "bottomLeft": 8,
                  "bottomRight": 8
                }
              },
              "child": {
                "type": "row",
                "mainAxisSize": "min",
                "children": [
                  {"type": "sizedBox", "width": 5},
                  {
                    "type": "column",
                    "children": [
                      {"type": "sizedBox", "height": 5},
                      {
                        "type": "text",
                        "data": "PDF",
                        "style": {"fontSize": 12, "color": "#FFFFFF"}
                      },
                      {"type": "sizedBox", "height": 5}
                    ]
                  },
                  {"type": "sizedBox", "width": 5}
                ]
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "expanded",
              "child": {
                "type": "text",
                "data": "%v",
                "maxLines": 2,
                "overflow": "ellipsis",
                "style": {
                  "fontSize": 14,
                  "fontWeight": "w700",
                  "color": "#222222"
                }
              }
            },
            {"type": "sizedBox", "height": 10}
          ]
        }
      }
    ]
  }
}`

	pdfMap := map[string]any{
		"type":    "pdf",
		"route":   "/pdfViewerScreen",
		"data":    model.Get.Content[0].URL,
		"content": fmt.Sprintf(str, model.Get.Content[0].Image, model.Get.Content[0].Title),
	}
	return pdfMap
}

func RenderBlog(model appmodels.ContentModel) map[string]any {
	str := `
	    {
  "type": "container",
  "height": 200,
  "width": 160,
  "decoration": {
    "color": "#FFFFFF",
    "borderRadius": {
      "topLeft": 10,
      "topRight": 10,
      "bottomLeft": 10,
      "bottomRight": 10
    },
    "boxShadow": [
      {
        "color": "#22000000",
        "blurRadius": 8,
        "offset": {"dx": 0, "dy": 2}
      }
    ]
  },
  "child": {
    "type": "row",
    "mainAxisSize": "min",
    "children": [
      {"type": "sizedBox", "width": 5},
      {
        "type": "sizedBox",
        "width": 150,
        "child": {
          "type": "column",
          "crossAxisAlignment": "start",
          "children": [
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "height": 90,
              "clipBehavior": "hardEdge",
              "decoration": {
                "color": "#F0F0F0",
                "borderRadius": {
                  "topLeft": 10,
                  "topRight": 10,
                  "bottomLeft": 10,
                  "bottomRight": 10
                }
              },
              "child": {
                "type": "image",
                "imageType": "network",
                "src":
                    "%v",
                "fit": "cover",
                "width": 150,
                "height": 200
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "decoration": {
                "color": "#3d1bd1",
                "borderRadius": {
                  "topLeft": 8,
                  "topRight": 8,
                  "bottomLeft": 8,
                  "bottomRight": 8
                }
              },
              "child": {
                "type": "row",
                "mainAxisSize": "min",
                "children": [
                  {"type": "sizedBox", "width": 5},
                  {
                    "type": "column",
                    "children": [
                      {"type": "sizedBox", "height": 5},
                      {
                        "type": "text",
                        "data": "Blog",
                        "style": {"fontSize": 12, "color": "#FFFFFF"}
                      },
                      {"type": "sizedBox", "height": 5}
                    ]
                  },
                  {"type": "sizedBox", "width": 5}
                ]
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "expanded",
              "child": {
                "type": "text",
                "data": "%v",
                "maxLines": 2,
                "overflow": "ellipsis",
                "style": {
                  "fontSize": 14,
                  "fontWeight": "w700",
                  "color": "#222222"
                }
              }
            },
            {"type": "sizedBox", "height": 10}
          ]
        }
      }
    ]
  }
}`

	blogMap := map[string]any{
		"type":    "blog",
		"route":   "/blogScreen",
		"data":    model.Get.Content[0].URL,
		"content": fmt.Sprintf(str, model.Get.Content[0].Image, model.Get.Content[0].Title),
	}
	return blogMap
}

func RenderAudio(model appmodels.ContentModel) map[string]any {
	str := `
 {
  "type": "container",
  "height": 200,
  "width": 160,
  "decoration": {
    "color": "#FFFFFF",
    "borderRadius": {
      "topLeft": 10,
      "topRight": 10,
      "bottomLeft": 10,
      "bottomRight": 10
    },
    "boxShadow": [
      {
        "color": "#22000000",
        "blurRadius": 8,
        "offset": {"dx": 0, "dy": 2}
      }
    ]
  },
  "child": {
    "type": "row",
    "mainAxisSize": "min",
    "children": [
      {"type": "sizedBox", "width": 5},
      {
        "type": "sizedBox",
        "width": 150,
        "child": {
          "type": "column",
          "crossAxisAlignment": "start",
          "children": [
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "height": 90,
              "clipBehavior": "hardEdge",
              "decoration": {
                "color": "#F0F0F0",
                "borderRadius": {
                  "topLeft": 10,
                  "topRight": 10,
                  "bottomLeft": 10,
                  "bottomRight": 10
                }
              },
              "child": {
                "type": "image",
                "imageType": "network",
                "src":
                    "%v",
                "fit": "cover",
                "width": 150,
                "height": 200
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "decoration": {
                "color": "#a2eb50",
                "borderRadius": {
                  "topLeft": 8,
                  "topRight": 8,
                  "bottomLeft": 8,
                  "bottomRight": 8
                }
              },
              "child": {
                "type": "row",
                "mainAxisSize": "min",
                "children": [
                  {"type": "sizedBox", "width": 5},
                  {
                    "type": "column",
                    "children": [
                      {"type": "sizedBox", "height": 5},
                      {
                        "type": "text",
                        "data": "Audio",
                        "style": {"fontSize": 12, "color": "#FFFFFF"}
                      },
                      {"type": "sizedBox", "height": 5}
                    ]
                  },
                  {"type": "sizedBox", "width": 5}
                ]
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "expanded",
              "child": {
                "type": "text",
                "data": "%v",
                "maxLines": 2,
                "overflow": "ellipsis",
                "style": {
                  "fontSize": 14,
                  "fontWeight": "w700",
                  "color": "#222222"
                }
              }
            },
            {"type": "sizedBox", "height": 10}
          ]
        }
      }
    ]
  }
}
  `

	audioMap := map[string]any{
		"type":    "audio",
		"route":   "/audioPlayerScreen",
		"data":    model.Get.Content[0].URL,
		"content": fmt.Sprintf(str, model.Get.Content[0].Image, model.Get.Content[0].Title),
	}
	return audioMap
}

func RenderVideo(model appmodels.ContentModel) map[string]any {
	str := `
 {
  "type": "container",
  "height": 200,
  "width": 160,
  "decoration": {
    "color": "#FFFFFF",
    "borderRadius": {
      "topLeft": 10,
      "topRight": 10,
      "bottomLeft": 10,
      "bottomRight": 10
    },
    "boxShadow": [
      {
        "color": "#22000000",
        "blurRadius": 8,
        "offset": {"dx": 0, "dy": 2}
      }
    ]
  },
  "child": {
    "type": "row",
    "mainAxisSize": "min",
    "children": [
      {"type": "sizedBox", "width": 5},
      {
        "type": "sizedBox",
        "width": 150,
        "child": {
          "type": "column",
          "crossAxisAlignment": "start",
          "children": [
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "height": 90,
              "clipBehavior": "hardEdge",
              "decoration": {
                "color": "#F0F0F0",
                "borderRadius": {
                  "topLeft": 10,
                  "topRight": 10,
                  "bottomLeft": 10,
                  "bottomRight": 10
                }
              },
              "child": {
                "type": "image",
                "imageType": "network",
                "src":
                    "%v",
                "fit": "cover",
                "width": 150,
                "height": 200
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "container",
              "decoration": {
                "color": "#f56642",
                "borderRadius": {
                  "topLeft": 8,
                  "topRight": 8,
                  "bottomLeft": 8,
                  "bottomRight": 8
                }
              },
              "child": {
                "type": "row",
                "mainAxisSize": "min",
                "children": [
                  {"type": "sizedBox", "width": 5},
                  {
                    "type": "column",
                    "children": [
                      {"type": "sizedBox", "height": 5},
                      {
                        "type": "text",
                        "data": "Video",
                        "style": {"fontSize": 12, "color": "#FFFFFF"}
                      },
                      {"type": "sizedBox", "height": 5}
                    ]
                  },
                  {"type": "sizedBox", "width": 5}
                ]
              }
            },
            {"type": "sizedBox", "height": 10},
            {
              "type": "expanded",
              "child": {
                "type": "text",
                "data": "%v",
                "maxLines": 2,
                "overflow": "ellipsis",
                "style": {
                  "fontSize": 14,
                  "fontWeight": "w700",
                  "color": "#222222"
                }
              }
            },
            {"type": "sizedBox", "height": 10}
          ]
        }
      }
    ]
  }
}
  `

	videoMap := map[string]any{
		"type":    "Video",
		"route":   "/videoPlayerScreen",
		"data":    model.Get.Content[0].URL,
		"content": fmt.Sprintf(str, model.Get.Content[0].Image, model.Get.Content[0].Title),
	}
	return videoMap
}

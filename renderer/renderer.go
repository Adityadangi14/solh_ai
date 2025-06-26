package renderer

import "regexp"

func Render(str string) (map[string]any, error) {

	renderedObject := make(map[string]any)

	re := regexp.MustCompile(`https?://[^\s]+`)

	urls := re.FindAllString(str, -1)

	cleanedText := re.ReplaceAllString(str, "")

	renderedObject["text"] = cleanedText

	recmom, err := ComponentRenderer(urls)

	if err != nil {
		return nil, err
	}

	renderedObject["recommendation"] = recmom

	return renderedObject, nil
}

package renderer

import (
	"regexp"
	"strings"

	"github.com/Adityadangi14/solh_ai/initializers"
)

func Render(str string) (map[string]any, error) {

	renderedObject := make(map[string]any)

	re := regexp.MustCompile(`https?://[^\s]+`)

	urls := re.FindAllString(str, -1)

	initializers.AppLogger.Info("Urls", "urls", strings.Join(urls, ","))

	cleanedText := re.ReplaceAllString(str, "")

	renderedObject["text"] = cleanedText

	recmom, err := ComponentRenderer(urls)

	if err != nil {
		return nil, err
	}

	renderedObject["recommendation"] = recmom

	return renderedObject, nil
}

package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

type Update struct {
	Doc map[string]interface{} `json:"doc"`
}

// Save updates the document in the index.
func (h Hit) Save(ctx context.Context) error {
	j, err := json.Marshal(Update{Doc: h.Source})
	if err != nil {
		return err
	}

	reader := strings.NewReader(string(j))

	req := opensearchapi.UpdateRequest{
		Index:      h.Index,
		DocumentID: h.ID,
		Body:       reader,
	}

	resp, err := req.Do(ctx, client)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("search engine status %d, response: %s", resp.StatusCode, body)
	}

	return nil
}

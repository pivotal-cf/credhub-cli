package credhub

import (
	"net/http"
	"net/url"
)

// DeleteByName will delete all versions of a credential by name
func (ch *CredHub) DeleteByName(name string) error {
	query := url.Values{}
	query.Set("name", name)
	resp, err := ch.Request(http.MethodDelete, "/api/v1/data", query, nil, true)

	if err == nil {
		defer resp.Body.Close()
	}

	return err
}

package cmd

import (
	"fmt"

	"github.com/anchore/syft"

	"github.com/khulnasoft-labs/changelogger/changelogger/release"
	"github.com/khulnasoft-labs/changelogger/changelogger/release/format"
	"github.com/khulnasoft-labs/changelogger/changelogger/release/format/json"
	"github.com/khulnasoft-labs/changelogger/changelogger/release/format/markdown"
)

type presentationTask func(description release.Description) (presenter.Presenter, error)

func selectPresenter(f format.Format) (presentationTask, error) {
	switch f {
	case format.MarkdownFormat:
		return presentMarkdown, nil
	case format.JSONFormat:
		return presentJSON, nil
	default:
		return nil, fmt.Errorf("unsupported output format: %+v", f)
	}
}

func presentMarkdown(description release.Description) (presenter.Presenter, error) {
	return markdown.NewMarkdownPresenter(markdown.Config{
		Description: description,
		Title:       appConfig.Title,
	})
}

func presentJSON(description release.Description) (presenter.Presenter, error) {
	return json.NewJSONPresenter(description)
}

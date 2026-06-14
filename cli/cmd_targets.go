package cli

import (
	"github.com/spf13/cobra"
)

func (a *App) targetsCmd() *cobra.Command {
	var goal string
	cmd := &cobra.Command{
		Use:   "targets",
		Short: "List UN SDG targets",
		RunE: func(cmd *cobra.Command, _ []string) error {
			limit := a.effectiveLimit(0)
			targets, err := a.client.Targets(cmd.Context(), goal, limit)
			if err != nil {
				return mapFetchErr(err)
			}
			return a.renderOrEmpty(targets, len(targets))
		},
	}
	cmd.Flags().StringVar(&goal, "goal", "", "filter by goal number (1-17)")
	return cmd
}

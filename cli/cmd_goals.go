package cli

import (
	"github.com/spf13/cobra"
)

func (a *App) goalsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goals",
		Short: "List all 17 UN Sustainable Development Goals",
		RunE: func(cmd *cobra.Command, _ []string) error {
			limit := a.effectiveLimit(0)
			goals, err := a.client.Goals(cmd.Context(), limit)
			if err != nil {
				return mapFetchErr(err)
			}
			return a.renderOrEmpty(goals, len(goals))
		},
	}
	return cmd
}

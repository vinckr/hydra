/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/hydra/cmd/cli"
)

// clientsDeleteCmd represents the delete command
func NewClientsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id> [<id>...]",
		Short: "Delete an OAuth 2.0 Client",
		Long: `This command deletes one or more OAuth 2.0 Clients by their respective IDs.

Example:
  hydra clients delete client-1 client-2 client-3`,
		Run: cli.NewHandler().Clients.DeleteClient,
	}
}

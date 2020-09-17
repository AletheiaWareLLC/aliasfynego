/*
 * Copyright 2020 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package aliasfynego

import (
	"encoding/base64"
	"fmt"
	"fyne.io/fyne/dialog"
	"github.com/AletheiaWareLLC/aliasfynego/ui"
	"github.com/AletheiaWareLLC/aliasgo"
	"github.com/AletheiaWareLLC/bcclientgo"
	"github.com/AletheiaWareLLC/bcfynego"
)

type AliasFyne struct {
	bcfynego.BCFyne
}

func (f *AliasFyne) NewList(client *bcclientgo.BCClient) *ui.AliasList {
	return ui.NewAliasList(func(id string, alias *aliasgo.Alias) {
		f.ShowAlias(client, id, alias)
	})
}

func (f *AliasFyne) ShowAlias(client *bcclientgo.BCClient, id string, alias *aliasgo.Alias) {
	info := fmt.Sprintf("Alias: %s\nPublicKey: %s\n", alias.Alias, base64.RawURLEncoding.EncodeToString(alias.PublicKey))
	if f.Dialog != nil {
		f.Dialog.Hide()
	}
	f.Dialog = dialog.NewInformation("Alias", info, f.Window)
	f.Dialog.Show()
}

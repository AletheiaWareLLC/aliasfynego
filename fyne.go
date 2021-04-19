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
	"aletheiaware.com/aliasgo"
	"aletheiaware.com/bcclientgo"
	"aletheiaware.com/bcfynego"
	"aletheiaware.com/bcfynego/ui"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AliasFyne interface {
	bcfynego.BCFyne
	ShowAlias(bcclientgo.BCClient, string, uint64, *aliasgo.Alias)
	ShowHelp(bcclientgo.BCClient)
}

type aliasFyne struct {
	bcfynego.BCFyne
}

func NewAliasFyne(a fyne.App, w fyne.Window) AliasFyne {
	return &aliasFyne{
		BCFyne: bcfynego.NewBCFyne(a, w),
	}
}

func (f *aliasFyne) ShowAlias(client bcclientgo.BCClient, id string, timestamp uint64, alias *aliasgo.Alias) {
	timestampScroller := container.NewHScroll(ui.NewTimestampLabel(timestamp))
	aliasScroller := container.NewHScroll(ui.NewAliasLabel(alias.Alias))
	bytesScroller := container.NewVScroll(ui.NewKeyLabel(alias.PublicKey))
	bytesScroller.SetMinSize(fyne.NewSize(0, 10*theme.TextSize())) // Show at least 10 lines
	formatScroller := container.NewVScroll(widget.NewLabel(alias.PublicFormat.String()))

	form := widget.NewForm(
		widget.NewFormItem(
			"Timestamp",
			timestampScroller,
		),
		widget.NewFormItem(
			"Alias",
			aliasScroller,
		),
		widget.NewFormItem(
			"Public Key",
			bytesScroller,
		),
		widget.NewFormItem(
			"Public Key Format",
			formatScroller,
		),
	)
	dialog := dialog.NewCustom("Alias", "OK", form, f.Window())
	dialog.Show()
	dialog.Resize(ui.DialogSize)
}

func (f *aliasFyne) ShowHelp(client bcclientgo.BCClient) {
	// TODO
	f.ShowError(fmt.Errorf("Not yet implemented: %s", "AliasFyne.ShowHelp"))
}

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
	"aletheiaware.com/bcgo"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AliasFyne struct {
	bcfynego.BCFyne
}

func NewAliasFyne(a fyne.App, w fyne.Window) *AliasFyne {
	return &AliasFyne{
		BCFyne: *bcfynego.NewBCFyne(a, w),
	}
}

func (f *AliasFyne) ShowAlias(client *bcclientgo.BCClient, id string, timestamp uint64, alias *aliasgo.Alias) {
	aliasScroller := container.NewHScroll(ui.NewAliasView(func() string { return alias.Alias }))
	publicKeyScroller := container.NewVScroll(ui.NewKeyView(func() []byte { return alias.PublicKey }))
	publicKeyScroller.SetMinSize(fyne.NewSize(0, 10*theme.TextSize())) // Show at least 10 lines

	form := widget.NewForm(
		widget.NewFormItem(
			"Timestamp",
			widget.NewLabel(bcgo.TimestampToString(timestamp)),
		),
		widget.NewFormItem(
			"Alias",
			aliasScroller,
		),
		widget.NewFormItem(
			"Public Key",
			publicKeyScroller,
		),
	)
	if d := f.Dialog; d != nil {
		d.Hide()
	}
	f.Dialog = dialog.NewCustom("Alias", "OK", form, f.Window)
	f.Dialog.Resize(ui.DialogSize)
	f.Dialog.Show()
}

func (f *AliasFyne) ShowHelp(client *bcclientgo.BCClient) {
	// TODO
	f.ShowError(fmt.Errorf("Not yet implemented: %s", "AliasFyne.ShowHelp"))
}

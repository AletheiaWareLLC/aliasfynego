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
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/aliasfynego/ui"
	"github.com/AletheiaWareLLC/aliasgo"
	"github.com/AletheiaWareLLC/bcclientgo"
	"github.com/AletheiaWareLLC/bcfynego"
)

type AliasFyne struct {
	bcfynego.BCFyne
}

func NewAliasFyne(a fyne.App, w fyne.Window) *AliasFyne {
	return &AliasFyne{
		BCFyne: *bcfynego.NewBCFyne(a, w),
	}
}

func (f *AliasFyne) NewList(client *bcclientgo.BCClient) *ui.AliasList {
	return ui.NewAliasList(func(id string, alias *aliasgo.Alias) {
		f.ShowAlias(client, id, alias)
	})
}

func (f *AliasFyne) ShowAlias(client *bcclientgo.BCClient, id string, alias *aliasgo.Alias) {
	publicKeyBase64 := base64.RawURLEncoding.EncodeToString(alias.PublicKey)
	var publicKeyRunes []rune
	for i, r := range []rune(publicKeyBase64) {
		if i > 0 && i%64 == 0 {
			publicKeyRunes = append(publicKeyRunes, '\n')
		}
		publicKeyRunes = append(publicKeyRunes, r)
	}

	aliasScroller := widget.NewHScrollContainer(widget.NewLabel(alias.Alias))
	publicKeyScroller := widget.NewHScrollContainer(widget.NewLabelWithStyle(string(publicKeyRunes), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}))
	publicKeyScroller.SetMinSize(fyne.NewSize(10*theme.TextSize(), 0))

	form := widget.NewForm(
		widget.NewFormItem(
			"Alias",
			aliasScroller,
		),
		widget.NewFormItem(
			"Public Key",
			publicKeyScroller,
		),
	)
	if f.Dialog != nil {
		f.Dialog.Hide()
	}
	f.Dialog = dialog.NewCustom("Alias", "OK", form, f.Window)
	f.Dialog.Show()
}

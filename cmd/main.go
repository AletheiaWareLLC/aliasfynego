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

package main

import (
	"aletheiaware.com/aliasfynego"
	"aletheiaware.com/aliasfynego/ui"
	"aletheiaware.com/aliasgo"
	"aletheiaware.com/bcclientgo"
	bcui "aletheiaware.com/bcfynego/ui"
	"aletheiaware.com/bcfynego/ui/data"
	"aletheiaware.com/bcgo"
	"flag"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"log"
)

var peer = flag.String("peer", "", "Alias peer")

func main() {
	// Parse command line flags
	flag.Parse()

	// Set log flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Create Application
	a := app.NewWithID("com.aletheiaware.alias")

	// Create Window
	w := a.NewWindow("Alias")

	// Create BC client
	c := bcclientgo.NewBCClient(bcgo.SplitRemoveEmpty(*peer, ",")...)

	// Create Alias Fyne
	f := aliasfynego.NewAliasFyne(a, w)

	// Create a scrollable list of registered aliases
	l := ui.NewAliasList(func(id string, timestamp uint64, alias *aliasgo.Alias) {
		go f.ShowAlias(c, id, timestamp, alias)
	})

	refreshList := func() {
		// Get local cache
		cache, err := c.GetCache()
		if err != nil {
			f.ShowError(err)
			return
		}
		// Get global network
		network, err := c.GetNetwork()
		if err != nil {
			f.ShowError(err)
			return
		}
		l.Update(cache, network)
	}

	// Populate list in goroutine
	go refreshList()

	t := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			go refreshList()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(data.NewPrimaryThemedResource(data.AccountIcon), func() {
			go f.ShowAccount(c)
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			go f.ShowHelp(c)
		}),
	)
	b := widget.NewButton("Register", func() {
		go func() {
			node, err := f.GetNode(c)
			if err != nil {
				f.ShowError(err)
				return
			}
			// Show Progress Dialog
			progress := dialog.NewProgress("Registering", "Registering "+node.Alias, f.Window)
			progress.Show()
			defer progress.Hide()
			listener := &bcui.ProgressMiningListener{Func: progress.SetValue}

			// Register Alias
			if err := aliasgo.Register(node, listener); err != nil {
				f.ShowError(err)
				return
			}
		}()
	})

	// Set window content, resize window, center window, show window, and run application
	w.SetContent(container.NewBorder(t, b, nil, nil, l))
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.ShowAndRun()
}

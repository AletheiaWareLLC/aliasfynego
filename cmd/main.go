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
	"flag"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/aliasfynego"
	"github.com/AletheiaWareLLC/aliasfynego/ui"
	"github.com/AletheiaWareLLC/aliasgo"
	"github.com/AletheiaWareLLC/bcclientgo"
	"github.com/AletheiaWareLLC/bcfynego"
	bcuidata "github.com/AletheiaWareLLC/bcfynego/ui/data"
	"github.com/AletheiaWareLLC/bcgo"
	"log"
)

var peer = flag.String("peer", "", "Alias peer")

func main() {
	// Create application
	a := app.New()

	// Create window
	w := a.NewWindow("Alias")
	w.SetMaster()

	peers := bcgo.SplitRemoveEmpty(*peer, ",")
	if len(peers) == 0 {
		peers = append(peers,
			bcgo.GetBCHost(), // Add BC host as peer
		)
	}

	// Create BC client
	c := &bcclientgo.BCClient{
		Peers: peers,
	}

	// Create Alias Fyne
	f := &aliasfynego.AliasFyne{
		BCFyne: bcfynego.BCFyne{
			App:    a,
			Window: w,
		},
	}

	// Create a scrollable list of registered aliases
	aliasList := f.NewList(c)
	go refresh(f, c, aliasList)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			log.Println("Refresh List")
			go refresh(f, c, aliasList)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(bcuidata.NewPrimaryThemedResource(bcuidata.AccountIcon), func() {
			log.Println("Account Info")
			go f.ShowAccount(c)
		}),
	)

	// Set window content, resize window, center window, show window, and run application
	w.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, aliasList))
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func refresh(fyne *aliasfynego.AliasFyne, client *bcclientgo.BCClient, list *ui.AliasList) {
	// Get local cache
	cache, err := client.GetCache()
	if err != nil {
		fyne.ShowError(err)
		return
	}
	// Get global network
	network, err := client.GetNetwork()
	if err != nil {
		fyne.ShowError(err)
		return
	}
	// Open Alias channel
	aliases := aliasgo.OpenAliasChannel()
	if err := aliases.Refresh(cache, network); err != nil {
		log.Println(err)
	}
	if err := aliasgo.IterateAliases(aliases, cache, network, list.Update); err != nil {
		fyne.ShowError(err)
		return
	}
	list.Refresh()
}

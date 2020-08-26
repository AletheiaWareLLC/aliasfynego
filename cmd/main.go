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
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/aliasgo"
	"github.com/AletheiaWareLLC/bcclientgo"
	"github.com/AletheiaWareLLC/bcfynego"
	"github.com/AletheiaWareLLC/bcgo"
	"log"
)

func main() {
	// Create application
	a := app.New()

	// Create window
	w := a.NewWindow("Alias")
	w.SetMaster()

	// Create BC client
	c := &bcfynego.BCFyneClient{
		BCClient: bcclientgo.BCClient{
			Peers: []string{
				bcgo.BC_HOST_TEST, // Add BC test node as peer
			},
		},
		App:    a,
		Window: w,
	}

	// Create a button to show the current node
	nodeButton := widget.NewButton("Node", func() {
		go c.ShowNode()
	})

	// Create a scrollable list of registered aliases
	aliasList := widget.NewVBox()
	aliasScroll := widget.NewScrollContainer(aliasList)

	refresh := func() {
		// Get local cache
		cache, err := c.GetCache()
		if err != nil {
			c.ShowError(err)
			return
		}
		// Get global network
		network, err := c.GetNetwork()
		if err != nil {
			c.ShowError(err)
			return
		}
		// Open Alias channel
		aliases := aliasgo.OpenAliasChannel()
		if err := aliases.Refresh(cache, network); err != nil {
			log.Println(err)
		}
		aliasList.Children = []fyne.CanvasObject{}
		// Iterate Aliases and populate list
		if err := aliasgo.IterateAliases(aliases, cache, network, func(alias *aliasgo.Alias) error {
			aliasList.Children = append(aliasList.Children, widget.NewLabel(alias.Alias))
			return nil
		}); err != nil {
			c.ShowError(err)
			return
		}
		// Trigger redraw
		aliasList.Refresh()
	}
	go refresh()

	// Create button to refresh list
	refreshButton := widget.NewButton("Refresh", func() {
		go refresh()
	})

	// Set window content, resize window, center window, show window, and run application
	w.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(refreshButton, nodeButton, nil, nil), refreshButton, nodeButton, aliasScroll))
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.ShowAndRun()
}

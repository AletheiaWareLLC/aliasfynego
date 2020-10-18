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

package ui

import (
	"encoding/base64"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/aliasgo"
	"github.com/AletheiaWareLLC/bcgo"
	"log"
	"sort"
)

type AliasList struct {
	widget.List
	ids        []string
	aliases    map[string]*aliasgo.Alias
	timestamps map[string]uint64
}

func NewAliasList(callback func(id string, timestamp uint64, alias *aliasgo.Alias)) *AliasList {
	l := &AliasList{
		aliases:    make(map[string]*aliasgo.Alias),
		timestamps: make(map[string]uint64),
		List: widget.List{
			CreateItem: func() fyne.CanvasObject {
				return &widget.Label{}
			},
		},
	}
	l.Length = func() int {
		return len(l.ids)
	}
	l.UpdateItem = func(index int, item fyne.CanvasObject) {
		a, ok := l.aliases[l.ids[index]]
		if ok {
			item.(*widget.Label).SetText(a.Alias)
		}
	}
	l.OnSelectionChanged = func(index int) {
		id := l.ids[index]
		if a, ok := l.aliases[id]; ok {
			callback(id, l.timestamps[id], a)
		}
	}
	l.ExtendBaseWidget(l)
	return l
}

func (l *AliasList) Add(entry *bcgo.BlockEntry, alias *aliasgo.Alias) error {
	id := base64.RawURLEncoding.EncodeToString(entry.RecordHash)
	if _, ok := l.aliases[id]; !ok {
		l.aliases[id] = alias
		l.timestamps[id] = entry.Record.Timestamp
		l.ids = append(l.ids, id)
		sort.Slice(l.ids, func(i, j int) bool {
			return l.timestamps[l.ids[i]] < l.timestamps[l.ids[j]]
		})
	}
	return nil
}

func (l *AliasList) Update(cache bcgo.Cache, network bcgo.Network) error {
	// Open Alias channel
	aliases := aliasgo.OpenAliasChannel()
	if err := aliases.Refresh(cache, network); err != nil {
		log.Println(err)
	}
	if err := aliasgo.IterateAliases(aliases, cache, network, l.Add); err != nil {
		return err
	}
	l.Refresh()
	return nil
}

// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

const CHANNEL_SEARCH_DEFAULT_LIMIT = 50

type ChannelSearch struct {
	Term                    string   `json:"term"`
	ExcludeDefaultChannels  bool     `json:"exclude_default_channels"`
	NotAssociatedToGroup    string   `json:"not_associated_to_group"`
	TeamIds                 []string `json:"team_ids"`
	GroupConstrained        bool     `json:"group_constrained"`
	ExcludeGroupConstrained bool     `json:"exclude_group_constrained"`
	Public                  bool     `json:"public"`
	Private                 bool     `json:"private"`
	IncludeDeleted          bool     `json:"include_deleted"`
	Deleted                 bool     `json:"deleted"`
	Page                    *int     `json:"page,omitempty"`
	PerPage                 *int     `json:"per_page,omitempty"`
	SearchColumns           []string `json:"search_columns"`
}

// ToJson convert a Channel to a json string
func (c *ChannelSearch) ToJson() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ChannelSearchFromJson will decode the input and return a Channel
func ChannelSearchFromJson(data io.Reader) *ChannelSearch {
	var cs *ChannelSearch
	json.NewDecoder(data).Decode(&cs)
	return cs
}


// SanitizeSearchColumns sanitizes the SearchColumns array and matches any entries to our columns
func SanitizeSearchColumns(columns []string) []string {
	var searchColumns []string

	for _, column := range columns {
		if column == "DisplayName" {
			searchColumns = append(searchColumns, "DisplayName")
		}
		if column == "Name" {
			searchColumns = append(searchColumns, "Name")
		}
		if column == "Purpose" {
			searchColumns = append(searchColumns, "Purpose")
		}
	}
	return searchColumns
}
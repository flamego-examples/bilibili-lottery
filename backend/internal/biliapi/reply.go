package biliapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ReplyType string

const (
	ReplyTypeVideo   = "1"
	ReplyTypeDynamic = "17"
)

type ListReplyOptions struct {
	Oid  string
	Type ReplyType
	Next int
}

type ListReplyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Cursor struct {
			IsBegin     bool   `json:"is_begin"`
			Prev        int    `json:"prev"`
			Next        int    `json:"next"`
			IsEnd       bool   `json:"is_end"`
			AllCount    int    `json:"all_count"`
			Mode        int    `json:"mode"`
			SupportMode []int  `json:"support_mode"`
			Name        string `json:"name"`
		} `json:"cursor"`
		Replies []struct {
			Rpid      int64  `json:"rpid"`
			Oid       int64  `json:"oid"`
			Type      int    `json:"type"`
			Mid       int    `json:"mid"`
			Root      int    `json:"root"`
			Parent    int    `json:"parent"`
			Dialog    int    `json:"dialog"`
			Count     int    `json:"count"`
			Rcount    int    `json:"rcount"`
			State     int    `json:"state"`
			Fansgrade int    `json:"fansgrade"`
			Attr      int    `json:"attr"`
			Ctime     int    `json:"ctime"`
			RpidStr   string `json:"rpid_str"`
			RootStr   string `json:"root_str"`
			ParentStr string `json:"parent_str"`
			Like      int    `json:"like"`
			Action    int    `json:"action"`
			Member    struct {
				Mid            string `json:"mid"`
				Uname          string `json:"uname"`
				Sex            string `json:"sex"`
				Sign           string `json:"sign"`
				Avatar         string `json:"avatar"`
				Rank           string `json:"rank"`
				FaceNftNew     int    `json:"face_nft_new"`
				IsSeniorMember int    `json:"is_senior_member"`
				LevelInfo      struct {
					CurrentLevel int `json:"current_level"`
					CurrentMin   int `json:"current_min"`
					CurrentExp   int `json:"current_exp"`
					NextExp      int `json:"next_exp"`
				} `json:"level_info"`
			} `json:"member"`
			Content struct {
				Message string `json:"message"`
				JumpUrl struct {
				} `json:"jump_url"`
				MaxLine     int `json:"max_line"`
				AtNameToMid struct {
					Field1 int `json:"！"`
				} `json:"at_name_to_mid,omitempty"`
			} `json:"content"`
		} `json:"replies"`
	} `json:"data"`
}

func ListReply(ctx context.Context, options ListReplyOptions) (*ListReplyResponse, error) {
	u := fmt.Sprintf("https://api.bilibili.com/x/v2/reply/main?callback=&jsonp=&next=%d&type=%s&oid=%s&mode=3&plat=1", options.Next, options.Type, options.Oid)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}
	req = req.WithContext(ctx)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	defer func() { _ = resp.Body.Close() }()

	var respJSON ListReplyResponse
	if err := json.NewDecoder(resp.Body).Decode(&respJSON); err != nil {
		return nil, errors.Wrap(err, "unmarshal response JSON")
	}
	return &respJSON, nil
}

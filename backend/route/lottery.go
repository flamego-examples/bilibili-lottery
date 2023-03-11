package route

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"

	"github.com/flamego/binding"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/flamego-examples/bilibili-lottery/context"
	"github.com/flamego-examples/bilibili-lottery/form"
	"github.com/flamego-examples/bilibili-lottery/internal/biliapi"
)

type LotteryHandler struct{}

func NewLotteryHandler() *LotteryHandler {
	return &LotteryHandler{}
}

func (*LotteryHandler) Lottery(ctx context.Context, form form.Lottery, errs binding.Errors) error {
	if len(errs) > 0 {
		return ctx.Error(errs[0].Err)
	}

	u, err := url.Parse(form.URL)
	if err != nil {
		return ctx.Error(errors.New("invalid url"))
	}

	allowedHosts := []string{
		"t.bilibili.com",
		"www.bilibili.com",
		"bilibili.com",
	}
	if !lo.Contains(allowedHosts, u.Host) {
		return ctx.Error(errors.Errorf("unexpected URL host: %q", u.Host))
	}

	var replyType biliapi.ReplyType
	var replyOid string
	switch {
	case u.Host == "t.bilibili.com":
		replyType = biliapi.ReplyTypeDynamic
		replyOid = strings.Trim(u.Path, "/")
	case strings.Contains(u.Path, "/opus/"):
		replyType = biliapi.ReplyTypeDynamic
		replyOid = strings.TrimPrefix(u.Path, "/opus/")
	case strings.Contains(u.Path, "/video/"):
		bvid := strings.TrimPrefix(u.Path, "/video/")
		bvid = strings.Trim(bvid, "/")
		videoInfo, err := biliapi.GetVideoInfo(ctx.Request().Context(), bvid)
		if err != nil {
			return ctx.Error(err)
		}
		replyType = biliapi.ReplyTypeVideo
		replyOid = strconv.Itoa(videoInfo.Data.Aid)
	}
	if replyOid == "" {
		return ctx.Error(errors.New("empty replyOID"))
	}

	type winner struct {
		UID       string `json:"uid"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatarURL"`
	}
	winners := make([]winner, 0)

	var next int
	for {
		replies, err := biliapi.ListReply(ctx.Request().Context(), biliapi.ListReplyOptions{
			Oid:  replyOid,
			Type: replyType,
			Next: next,
		})
		if err != nil {
			return ctx.Error(err)
		}

		for _, reply := range replies.Data.Replies {
			winners = append(winners, winner{
				UID:       reply.Member.Mid,
				Name:      reply.Member.Uname,
				AvatarURL: reply.Member.Avatar,
			})
		}
		if replies.Data.Cursor.IsEnd {
			break
		}
		next = replies.Data.Cursor.Next
	}

	winners = lo.UniqBy(winners, func(winner winner) string {
		return winner.UID
	})

	rand.Shuffle(len(winners), func(i, j int) {
		winners[i], winners[j] = winners[j], winners[i]
	})

	winnerCount := form.WinnerCount
	if len(winners) < winnerCount {
		winnerCount = len(winners)
	}
	winners = winners[:winnerCount]

	return ctx.Success(winners)
}

package client

import (
	"context"
	"strconv"

	"github.com/dnsimple/dnsimple-go/dnsimple"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	Spec       Spec
	Dnsimple   *dnsimple.Client
	DnsimpleID string
}

func (c *Client) ID() string {
	return "dnsimple"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {

	tc := dnsimple.StaticTokenHTTPClient(ctx, s.Token)
	cl := dnsimple.NewClient(tc)

	if s.BaseURL != "" {
		cl.BaseURL = s.BaseURL
	}

	if s.UserAgent != "" {
		cl.SetUserAgent(s.UserAgent)
	}

	res, err := cl.Identity.Whoami(ctx)
	if err != nil {
		return Client{}, err
	}

	accountID := strconv.FormatInt(res.Data.Account.ID, 10)

	c := Client{
		logger:     logger,
		Spec:       *s,
		Dnsimple:   cl,
		DnsimpleID: accountID,
	}

	return c, nil
}

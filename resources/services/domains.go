package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/dnsimple/dnsimple-go/dnsimple"
	"github.com/razin99/cq-source-dnsimple/client"
)

func DomainsTable() *schema.Table {
	return &schema.Table{
		Name:      "dnsimple_domains",
		Resolver:  fetchDomains,
		Transform: transformers.TransformWithStruct(&dnsimple.Domain{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	page := 1

	fetchParams := &dnsimple.DomainListOptions{
		ListOptions: dnsimple.ListOptions{
			Page: &page,
		},
	}

	for {
		queryOk, err := cl.Dnsimple.Domains.ListDomains(ctx, cl.DnsimpleID, fetchParams)
		if err != nil {
			return err
		}

		for _, domain := range queryOk.Data {
			res <- domain
		}

		if queryOk.Pagination.TotalPages == page {
			break
		} else {
			page++
		}

	}
	return nil
}

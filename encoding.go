package dagobert

import (
	"fmt"

	"github.com/jina-ai/client-go/docarray"
	"github.com/jina-ai/client-go/jina"
	"github.com/samber/mo"
)

func (c *Client) Encode(docs []*Document) ([]*docarray.DocumentProto, error) {
	requests := make(chan *jina.DataRequestProto)
	defer close(requests)

	go func() {
		req := &jina.DataRequestProto{
			Data: &jina.DataRequestProto_DataContentProto{
				Documents: &jina.DataRequestProto_DataContentProto_Docs{
					Docs: &docarray.DocumentArrayProto{
						Docs: documentsToArray(docs),
					},
				},
			},
		}

		requests <- req
	}()

	return mo.
		NewFuture(func(resolve func([]*docarray.DocumentProto), reject func(error)) {
			onDone := func(result *jina.DataRequestProto) {
				docs := result.GetData().GetDocs().GetDocs()
				resolve(docs)
			}
			onError := func(result *jina.DataRequestProto) {
				err := NewDagobertError(result.Header.GetStatus())
				reject(fmt.Errorf("💥 failed to encode: %w", err))
			}

			err := c.client.POST(requests, onDone, onError, nil)
			if err != nil {
				reject(err)
			}
		}).
		Either().
		Swap().
		Unpack()
}

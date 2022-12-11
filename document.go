package dagobert

import (
	"strconv"

	"github.com/jina-ai/client-go/docarray"
	"github.com/samber/lo"
)

const (
	TextDocumentType = iota
	URIDocumentType
	ReaderDocumentType
)

type Document struct {
	tYpe int

	text string
	uri  string
	blob []byte
}

func NewTextDocument(text string) *Document {
	return &Document{
		tYpe: TextDocumentType,
		text: text,
	}
}

func NewURIDocument(uri string) *Document {
	return &Document{
		tYpe: URIDocumentType,
		uri:  uri,
	}
}

func NewBlobDocument(blob []byte) *Document {
	return &Document{
		tYpe: ReaderDocumentType,
		blob: blob,
	}
}

func (d *Document) toDocumentProto(id string) *docarray.DocumentProto {
	switch d.tYpe {
	case TextDocumentType:
		return &docarray.DocumentProto{
			Id: id,
			Content: &docarray.DocumentProto_Text{
				Text: d.text,
			},
		}
	case URIDocumentType:
		return &docarray.DocumentProto{
			Id:  id,
			Uri: d.uri,
		}
	case ReaderDocumentType:
		return &docarray.DocumentProto{
			Id: id,
			Content: &docarray.DocumentProto_Blob{
				Blob: d.blob,
			},
		}
	}

	panic("unexpected document type")
}

func documentsToArray(documents []*Document) []*docarray.DocumentProto {
	return lo.Map(documents, func(doc *Document, index int) *docarray.DocumentProto {
		return doc.toDocumentProto(strconv.FormatInt(int64(index), 10))
	})
}

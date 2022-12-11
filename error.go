package dagobert

import (
	"fmt"

	"github.com/jina-ai/client-go/jina"
)

type DagobertError struct {
	code        jina.StatusProto_StatusCode
	description string
	exception   *jina.StatusProto_ExceptionProto
}

func NewDagobertError(status *jina.StatusProto) error {
	return &DagobertError{
		code:        status.GetCode(),
		description: status.GetDescription(),
		exception:   status.GetException(),
	}
}

func (e DagobertError) Error() string {
	return fmt.Sprintf("jina: error code %d (%s). Message: %s", e.code, e.description, e.exception)
}

package gapi

import "context"

type Metadata struct {
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}
	return mtdt
}

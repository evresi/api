package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func (s *Server) routePOI(r chi.Router) {
	r.Get("/{poiID}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "poiID")

		if id, err := uuid.Parse(idParam); err == nil {
			row := s.db.QueryRow(context.Background(), "SELECT id, name, thumbnail, description FROM poi WHERE id = $1", id.String())

			poi, err := parsePOIFromRow(row)
			if err != nil {
				w.WriteHeader(500)
				return
			}

			bytes, err := json.Marshal(poi)
			if err != nil {
				w.WriteHeader(500)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", fmt.Sprint(len(bytes)))
			w.Write(bytes)
		} else {
			w.WriteHeader(400)
		}
	})
}

func parsePOIFromRow(row pgx.Row) (POI, error) {
	var idPg pgtype.UUID
	var name string
	var thumbnailPg pgtype.UUID
	var description string
	if err := row.Scan(&idPg, &name, &thumbnailPg, &description); err != nil {
		return POI{}, err
	}

	var thumbnail *uuid.UUID
	if thumbnailPg.Status == pgtype.Present {
		tn, err := uuid.FromBytes(thumbnailPg.Bytes[:])
		if err == nil {
			thumbnail = &tn
		} else {
			return POI{}, err
		}
	}

	var id *uuid.UUID
	if idPg.Status == pgtype.Present {
		i, err := uuid.FromBytes(idPg.Bytes[:])
		if err == nil {
			id = &i
		} else {
			return POI{}, err
		}
	}

	return POI{ID: id, Name: name, Thumbnail: thumbnail, Description: description}, nil
}

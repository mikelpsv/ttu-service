package models

import (
	"database/sql"
	"fmt"
)

type TrackService struct {
	Db *sql.DB
}

type Track struct {
	Id  int64  `json:"id"`
	Uid string `json:"uid"`
}

func (ts *TrackService) CreateTrack(namespace string, trackUid string) (int64, error) {
	nServ := NamespaceService{ts.Db}
	ns, err := nServ.FindNamespaceByName(namespace)
	if err != nil {
		return 0, err
	}

	sqlTrack := fmt.Sprintf("INSERT INTO storage%d SET track", ns.Id, trackUid)
	res, err := ts.Db.Exec(sqlTrack)
	if err != nil {
		return 0, err
	}

	afRows, err := res.RowsAffected()
	if err != nil {
		return afRows, err
	}
	return afRows, nil
}

func (ts *TrackService) GetTrack(namespace string, trackUid string) (*Track, error) {
	nServ := NamespaceService{ts.Db}
	ns, err := nServ.FindNamespaceByName(namespace)
	if err != nil {
		return nil, err
	}

	sqlTrack := fmt.Sprintf("SELECT id, uid FROM storage%d WHERE track_uid = %s", ns.Id, trackUid)
	rows, err := ts.Db.Query(sqlTrack)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	track := new(Track)
	if rows.Next() {
		err = rows.Scan(&track.Id, &track.Uid)
		if err != nil {
			return nil, err
		}
	}
	return track, nil
}

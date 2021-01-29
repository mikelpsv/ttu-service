package models

import (
	"database/sql"
	"fmt"
)

type NamespaceService struct {
	Db *sql.DB
}

type Namespace struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (ns *NamespaceService) FindNamespaceByName(namespaceName string) (*Namespace, error) {
	sqlNs := fmt.Sprintf("SELECT id, name FROM namespaces WHERE name = %s", namespaceName)
	rows, err := ns.Db.Query(sqlNs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		ns := new(Namespace)
		err := rows.Scan(&ns.Id, &ns.Name)
		if err != nil {
			return nil, err
		}
		return ns, err
	}
	return nil, sql.ErrNoRows
}

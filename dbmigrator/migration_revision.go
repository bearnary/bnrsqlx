package bnrdbmigrator

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type MigrationRevision interface {
	Revision() int64
	ShouldInstall(installedRevision int64) bool
	StartMigrate(db *gorm.DB, path string, getBinData func(path string) ([]byte, error)) error
}

type RevisionBase struct {
	RevisionVal int64
	FileName    string
}

func (s *RevisionBase) Revision() int64 {
	return s.RevisionVal
}

func (s *RevisionBase) ShouldInstall(installedRevision int64) bool {
	return s.RevisionVal > installedRevision
}

func (s *RevisionBase) StartMigrate(db *gorm.DB, path string, getBinData func(path string) ([]byte, error)) error {
	data, err := getBinData(fmt.Sprintf("%s%s", path, s.FileName))
	if err != nil {
		return err
	}

	cmds := strings.Split(string(data), ";\n")

	for _, cmd := range cmds {
		cmd = strings.Trim(cmd, "\n")
		if cmd == "" {
			continue
		}
		err := db.Exec(cmd).Error
		if err != nil {
			return err
		}
	}
	return nil
}

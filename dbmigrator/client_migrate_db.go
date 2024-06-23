package bnrdbmigrator

import (
	"time"

	"github.com/kazekim/kzklogger"
	"gorm.io/gorm"
)

func (c *defaultClient) MigrateDB(revisions []MigrationRevision) error {
	log := kzklogger.New("AutoMigration")
	log.Info("start")
	defer func() {
		log.Info("completed")
	}()

	err := c.mustHaveVersionTable()
	if err != nil {
		return err
	}

	updatingRevision := time.Now().UnixNano()
	log.Infof("revision %d ", updatingRevision)
	var tx *gorm.DB
	defer func() {
		if tx != nil {
			tx.Commit()
		}
	}()

	for tx, err = c.canRecordUpdatingRevision(updatingRevision); err != nil; {
		tx, err = c.canRecordUpdatingRevision(updatingRevision)
		if err != nil {
			log.WithError(err).Println("can't record updating revision")
		}
	}

	installedRevision, err := c.mustGetInstalledRevision()
	if err != nil {
		return err
	}
	log.Infof("installed revision is %d, checking for more update", installedRevision)

	for _, s := range revisions {
		if !s.ShouldInstall(installedRevision) {
			log.Infof("revision %3d: skipped", s.Revision())
			continue
		}

		if err := s.StartMigrate(c.db, c.cfg.SQLPath, c.getBinData); err != nil {
			log.Errorf("revision %3d: panic", s.Revision())
			return err
		} else {
			err = c.mustRecordInstalledRevision(s.Revision())
			if err != nil {
				return err
			}
			log.Infof("revision %3d: installed", s.Revision())
		}
	}
	return nil
}

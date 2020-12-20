package db

import (
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/ztrue/tracerr"
)

// Where's DeleteFileHash? DeleteFileMeta handles that for us
func (db *HarmonyDB) AddFileHash(fileID string, hash []byte) error {
	return tracerr.Wrap(db.queries.AddHash(ctx, queries.AddHashParams{
		Hash:   hash,
		FileID: fileID,
	}))
}

func (db *HarmonyDB) GetFileIDByHash(hash []byte) (string, error) {
	data, err := db.queries.GetFileIDByHash(ctx, hash)
	err = tracerr.Wrap(err)
	return data, err
}

func (db *HarmonyDB) SetFileMetadata(fileID string, contentType, name string, size int32) error {
	return tracerr.Wrap(db.queries.AddFileMetadata(ctx, queries.AddFileMetadataParams{
		FileID:      fileID,
		ContentType: contentType,
		Name:        name,
		Size:        size,
	}))
}

func (db *HarmonyDB) GetFileMetadata(fileID string) (queries.GetFileMetadataRow, error) {
	data, err := db.queries.GetFileMetadata(ctx, fileID)
	err = tracerr.Wrap(err)
	return data, err
}

func (db *HarmonyDB) DeleteFileMeta(fileID string) error {
	return tracerr.Wrap(db.queries.DeleteFileMetadata(ctx, fileID))
}

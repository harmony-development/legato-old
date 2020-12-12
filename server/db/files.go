package db

import "github.com/harmony-development/legato/server/db/queries"

// Where's DeleteFileHash? DeleteFileMeta handles that for us
func (db *HarmonyDB) AddFileHash(fileID string, hash []byte) error {
	return db.queries.AddHash(ctx, queries.AddHashParams{
		Hash:   hash,
		FileID: fileID,
	})
}

func (db *HarmonyDB) GetFileIDByHash(hash []byte) (string, error) {
	return db.queries.GetFileIDByHash(ctx, hash)
}

func (db *HarmonyDB) SetFileMetadata(fileID string, contentType, name string, size int32) error {
	return db.queries.AddFileMetadata(ctx, queries.AddFileMetadataParams{
		FileID:      fileID,
		ContentType: contentType,
		Name:        name,
		Size:        size,
	})
}

func (db *HarmonyDB) GetFileMetadata(fileID string) (queries.GetFileMetadataRow, error) {
	return db.queries.GetFileMetadata(ctx, fileID)
}

func (db *HarmonyDB) DeleteFileMeta(fileID string) error {
	return db.queries.DeleteFileMetadata(ctx, fileID)
}

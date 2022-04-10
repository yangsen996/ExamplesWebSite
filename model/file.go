package model

type FileMeta struct {
	FileShal     string
	FileName     string
	FileSize     int64
	LocationPath string
	UploadAt     string
}

var fileMetaMap map[string]FileMeta

func init() {
	fileMetaMap = make(map[string]FileMeta)
}

func AddFileMeta(fmeta FileMeta) {
	fileMetaMap[fmeta.FileShal] = fmeta
}

func GetFileMeta(fileShal string) interface{} {
	v, ok := fileMetaMap[fileShal]
	if !ok {
		return nil
	}
	return v
}

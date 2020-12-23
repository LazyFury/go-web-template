package upload

// NewEchoUploader EchoUploader
func NewEchoUploader() *Uploader {
	return &Uploader{
		BaseDir:      "./static/upload",
		UploadMethod: DefaultUpload,
		GetFile:      DefaultGetFile,
	}
}

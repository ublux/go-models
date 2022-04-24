package models

type StoredFile struct {
	Id              string `bson:"_id" json:"id"`
	FileName        string `bson:"fileName" json:"fileName"`
	FileSizeInBytes int32  `bson:"fileSizeInBytes" json:"fileSizeInBytes"`
	FolderName      string `bson:"folderName" json:"folderName"`
	IdAccount       string `bson:"idAccount" json:"idAccount"`
	Md5Hash         string `bson:"md5Hash" json:"md5Hash"`
	Url             string `bson:"url" json:"url"`
}

// Implementing interface IUbluxSubDocument
// Implementing interface IUbluxDocumentId
func (x StoredFile) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x StoredFile) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxSubDocument

// BUILDER from bson map:
func BuildStoredFile(m map[string]interface{}, x *StoredFile) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["fileName"]; ok && val != nil {
		x.FileName = val.(string)
	}
	if val, ok := m["fileSizeInBytes"]; ok && val != nil {
		x.FileSizeInBytes = val.(int32)
	}
	if val, ok := m["folderName"]; ok && val != nil {
		x.FolderName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["md5Hash"]; ok && val != nil {
		x.Md5Hash = val.(string)
	}
	if val, ok := m["url"]; ok && val != nil {
		x.Url = val.(string)
	}
}

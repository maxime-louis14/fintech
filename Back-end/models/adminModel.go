package models

type AdminForm struct {
	Titre string `json:"titre"`
	File  []byte `json:"file"` // Utilisez []byte si vous stockez les données binaires du fichier
	// Ajoutez d'autres champs si nécessaire
}

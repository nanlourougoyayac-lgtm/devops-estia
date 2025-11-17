package main

import "net/http"

// func getCat(req *http.Request) (int, any) {
// 	Logger.Info("Cat not found")
// 	return http.StatusNotFound, "Cat not found"
// }
func getCat(req *http.Request) (int, any) {
	// Récupérer l'ID depuis l'URL : /cats/{id}
	catID := req.PathValue("id")
	Logger.Infof("Getting cat with ID: %s", catID)

	// Chercher le chat dans la base
	cat, exists := catsDatabase[catID]
	if !exists {
		Logger.Warnf("Cat '%s' not found", catID)
		return http.StatusNotFound, "Cat not found"
	}

	// Retourner le chat trouvé
	return http.StatusOK, cat
}

func deleteCat(req *http.Request) (int, any) {
	// Récupérer l'ID depuis l'URL : /cats/{id}
	catID := req.PathValue("id")
	Logger.Infof("Deleting cat with ID: %s", catID)

	// Vérifier si le chat existe
	_, exists := catsDatabase[catID]
	if !exists {
		Logger.Warnf("Cat '%s' not found", catID)
		return http.StatusNotFound, "Cat not found"
	}

	// Suppression du chat
	delete(catsDatabase, catID)

	Logger.Infof("Cat '%s' successfully deleted", catID)
	return http.StatusOK, "Cat deleted"
}


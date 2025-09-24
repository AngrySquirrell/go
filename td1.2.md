On applique les concepts vus pour améliorer notre code :

1. Remplacer map[int]map[string]string par map[int]\*Contact.
   - On stocke des pointers pour pouvoir modifier les contacts facilement.
2. Transformer les fonctions (ajouter, supprimer...) en méthodes attachées à la struct Contact (quand celà a du sens)
3. Utiliser un constructeur NewContact pour valider les données à la création.

Résultats : Un code plus sûr, plus lisible et mieux organisé.

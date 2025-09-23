Mini-CRM CLI

> Un système de gestion de contacts minimaliste en ligne de commande développé en Go

## 🎯 Principe

Application CLI simple et efficace pour gérer vos contacts. Deux modes d'utilisation :

- **Interactif** : Menu principal avec navigation
- **Direct** : Ajout rapide via flags de ligne de commande

## 🚀 Utilisation

### Mode Interactif

```bash
go run main.go
```

### Mode Direct (Flags)

```bash
go run main.go --ajouter --name="Jean Dupont" --mail="jean@exemple.com"
```

## ⚡ Fonctionnalités

- ✅ Ajouter un contact (validation des données à la création)
- ✅ Lister tous les contacts
- ✅ Supprimer un contact par ID
- ✅ Mettre à jour un contact (validation des données)
- ✅ Interface console nettoyée automatiquement
- ✅ Structure Contact améliorée (pointeurs, méthodes)
- ✅ Code plus sûr et lisible grâce à l'utilisation de méthodes et d'un constructeur

## 🛠️ Concepts Go utilisés

- `for`, `switch`, `map[int]*Contact`, méthodes, constructeur, gestion d'erreur, `strconv`, `os.Stdin`, `bufio`, etc.

## 📄 Fichier

- `main.go` : code source principal
- `README.md` : ce fichier
- `td1.md` : consignes initiales
- `td2.md` : consignes d'amélioration

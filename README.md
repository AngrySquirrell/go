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

- ✅ Ajouter un contact
- ✅ Lister tous les contacts
- ✅ Supprimer un contact par ID
- ✅ Mettre à jour un contact
- ✅ Interface console nettoyée automatiquement

## Repo

Chaques étapes de developpement du CLI ont été étagé via les branches.
La main contient uniquement les énoncés et le readme.

- `td1` correspond aux exigences de `td1.md`
- `td2` correspond aux exigences de `td2.md`

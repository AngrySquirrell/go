# Système de Notifications et Logging (Go)

Ce projet est une application en ligne de commande écrite en Go qui simule l'envoi de notifications via Email, SMS et Push, avec archivage des envois réussis.

## Fonctionnalités

- **Notifications colorées** : chaque type de notification s'affiche dans une couleur différente (Email : bleu, SMS : vert, Push : cyan, Erreur : rouge).
- **Validation SMS** : les numéros non valides sont signalés en rouge et ne sont pas archivés.
- **Archivage** : chaque notification envoyée avec succès est enregistrée avec son message et un timestamp.
- **Historique** : affichage de l'historique des notifications archivées à la fin du programme.

## Utilisation

1. Lancez le programme avec `go run main.go`.
2. Les notifications sont simulées et affichées dans le terminal avec leur couleur.
3. À la fin, l'historique des notifications archivées s'affiche avec la date et l'heure d'envoi.

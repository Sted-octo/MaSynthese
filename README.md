# Octoptimist

## Why 
Présenter une synthèse d'activité avec une répartition done/todo et un calcul du TACE maximal ateignable si d'aventure tous l'intercontrat à venir se tranformait en mission facturable.

## How
En lien avec l'API d'Octopod, Octoptimist récupère les informations de pointage d'un Octo et les présente de manière synthétique avec les ruptures qui vont bien, et effectue les cumuls souhaités.

## What
Application Web, avec des pages rendues côté serveur, écrite en Go

## Title and Logo
![Logo Octopimist](/static/octoptimist.svg?raw=true "Logo Octopimist")

Jeu de mot entre OCTO et Optimiste d'une par, car l'une des principale fonctionnalitées est de présenter le TACE optimiste.
Autre jeu de mot, pti "mist" = Petit brouillard (représenté par les vagelettes dans le logo) : Car c'est parfois difficile d'avancer et d'avoir une vision claire de son activité future dans un monde d'incertitude.

## Définitions 

### TACE
Taux d'activité facturé congès exclus.
Sur une période donnée : TACE = Total des jours facturés / (TOTAL des jours ouvrés - cumul des congés) x 100 

### TACE FYxx
TACE calculé par Octopod, sans retraitement. Sur la période fiscale.
### TACE Période
TACE calculé calculé par Octoptimist sur la période choisie à l'écran, qui peut être différente de la période fiscale (de Septembre à Août).
Quand le résultat est identique au TACE FYxx, cette information n'est pas affichée.
### TACE FYxx Optimist
TACE calculé par Octoptimist. 
Sur une période fiscale, Octoptimist complète les jours à venir (après aujourd'hui) non saisie dans Octopod par de l'intercontrat.
L'intercontrat à venir est considéré comme une activité facturable, et donc fera grossir le TACE.
Il montre ainsi le TACE maximal ateignable sur une période.
### TACE FYxx I.G. inclus
TACE calculé calculé par Octoptimist sur la période fiscale.
Les activités d'interêt général (I.G.) sont considérées comme du facturable, et donc feront grossir le TACE.
Quand le résultat est identique au TACE FYxx ou au Tace Période, cette information n'est pas affichée.

## Clean Archi
```
------------------------------------------------------------
| -----------------------------------------                |
| | -----------------------               |                |
| | | ----------          | presenters    | ui             |
| | | | domain | usecases |               |                |
| | | ----------          | dataproviders | infrastructure |
| | -----------------------               |                |
| -----------------------------------------                |
------------------------------------------------------------
```

## Env Variables

 * PORT : port d'écoute de l'application, valeur par défaut 9090
 * CLIENT_ID : ID de l'application du point de vue de l'API Octopod
 * CLIENT_SECRET : Mot de passe de l'application du point de vue de l'API Octopod
 * REDIRECT_URL : URL qui sera transmise à l'API Octopod durant la phase d'autentification, afin qu'Octopod puisse effectuer sa redirection une fois l'autentification correcte. Dois être accessible via internet.
 * OCTOPOD_DOMAIN : Base de l'URL d'acccès à Octopod.
 * DEBUG : (true/false) permet d'activer le log détaillé.

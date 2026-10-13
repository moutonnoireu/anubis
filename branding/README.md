# Personnalisation Medee (sans rebuild)

Variables d'environnement :

| Variable          | Défaut  | Effet                                            |
| :---------------- | :------ | :----------------------------------------------- |
| `BRAND_NAME`      | `Medee` | Nom affiché dans le pied de page                 |
| `BRAND_LINK`      | (vide)  | URL liée au nom (vide = pas de lien)             |
| `CHALLENGE_TITLE` | (locale)| Titre de la page de vérification                |
| `ERROR_TITLE`     | (locale)| Titre des pages d'erreur                        |
| `OVERLAY_DIR`     | (vide)  | Dossier d'assets prioritaires (images/CSS)       |

Exemple :

```sh
BRAND_NAME="Medee" BRAND_LINK="https://intranet.example.com" \
CHALLENGE_TITLE="Vérification de sécurité en cours" \
OVERLAY_DIR="./branding" ./anubis --target http://localhost:3000
```

## Dossier overlay

Tout fichier présent ici remplace l'asset embarqué de même chemin,
le reste retombe sur les assets embarqués :

```text
branding/
└── static
    ├── css
    │   └── brand.css         # remplace le theme embarque (exemple fourni)
    └── img
        ├── happy.webp        # vérification réussie
        ├── pensive.webp      # vérification en cours
        └── reject.webp       # erreur / blocage
```

Images : WebP 384×384. Les icônes par défaut du fork (géométrie plate
générée en code, sans artwork tiers) sont dans `web/static/img/`.

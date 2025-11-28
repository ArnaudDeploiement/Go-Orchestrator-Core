# 🚀 Orchestrator Go — Pitch

Une couche d’orchestration **multi‑agent** écrite en pur Go, sans framework, pour piloter des agents IA comme un chef d’orchestre. Elle s’appuie sur trois briques simples :
- 🧠 **State (sac à dos)** : le contexte partagé que chaque étape lit/écrit.
- 🗺️ **DAG (plan clair)** : un enchaînement start → intent → router → agent → output, avec branches/parallèle sans boucle infinie.
- 🧱 **Nœuds / Agents / Tools** : chaque étape fait une micro‑tâche et sait appeler ses tools métiers.

Objectif : une orchestration lisible, testable, extensible, prête pour la prod (timeouts, limites, obs, CI, Docker dans la roadmap).

Lancement rapide :
```bash
go run ./...
curl -XPOST localhost:8080/ia -d '{"content":"Bonjour"}'
```

Roadmap courte :
- Stabiliser HTTP/LLM et erreurs (ARN-38).
- Socle State + DAG + nœuds de base + branchement HTTP (ARN-39/40/41).
- Interface Tool + parallélisation (ARN-42).
- Observabilité, config, CI (ARN-43/44).
- Packaging Docker/Makefile, README final (ARN-46/47).

# OpsNexus Common (`opsnexus-common`)

[![Release](https://img.shields.io/badge/release-v0.5.0-blue.svg)](https://github.com/OpsNexusHQ/opsnexus-common/releases/tag/v0.5.0)
[![Go Version](https://img.shields.io/badge/go-1.25+-00ADD8.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Shared Go data models, structures, and protocol contracts used across the **OpsNexus** ecosystem (shared between `opsnexus-agent` and `opsnexus-backend`).

---

## 📦 Package Contents

```text
opsnexus-common/
├── go.mod
├── README.md
└── models/
    ├── agent.go           # Agent registration and health structs
    ├── registration.go    # Heartbeat and payload models
    └── telemetry.go       # System snapshot, CPU, Memory, Disk, Network models
```

---

## ⚙️ How to Import in Go Projects

```go
import (
    "github.com/OpsNexusHQ/opsnexus-common/models"
)

func main() {
    snapshot := models.SystemSnapshot{
        // ...
    }
}
```

Add dependency in your Go project's `go.mod`:

```bash
go get github.com/OpsNexusHQ/opsnexus-common@v0.5.0
```

---

## 📄 License

Part of the [OpsNexus](https://github.com/OpsNexusHQ) ecosystem. Licensed under the MIT License.

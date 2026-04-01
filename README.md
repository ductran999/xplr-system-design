# System Design Notes

> A personal learning project to explore and understand **System Design concepts** through structured notes, examples, and hands-on experimentation.

This repository serves as a **playground for studying backend architecture**, covering everything from **how systems communicate** to **how large-scale systems are designed, scaled, and kept resilient**.

---

```sh
📁 system-design-playground/
│
├── 📁 1-client-server-communication/       # HOW COMPUTERS TALK (Network Protocols)
│   ├── 📁 01-rest-api/                     # Traditional Request-Response (HTTP/JSON)
│   ├── 📁 02-short-polling/                # Client repeatedly polls (Spamming)
│   ├── 📁 03-long-polling/                 # Server holds connection waiting for data
│   ├── 📁 04-sse-server-sent-events/       # Server pushes one-way data (Streaming)
│   ├── 📁 05-websocket/                    # Full-duplex two-way channel (Real-time)
│   ├── 📁 06-grpc/                         # High-speed internal calls (Protobuf/HTTP2)
│   ├── 📁 07-graphql/                      # Flexible data querying (Query Language)
│   ├── 📁 08-webhooks/                     # Reverse events from third parties (Callback)
│   └── 📁 09-webrtc/                       # Peer-to-peer communication (P2P Video/Data)
│
├── 📁 2-system-architecture/               # HOW SERVERS ARE ORGANIZED (Topology/Deployment)
│   ├── 📁 01-monolithic/                   # Everything in one unit, one shared DB
│   ├── 📁 02-service-based/                # Split into services but STILL share one DB
│   ├── 📁 03-microservices/                # Fully decoupled, EACH SERVICE has its own DB
│   ├── 📁 04-event-driven-eda/             # Communication via Message Broker (Kafka/RabbitMQ)
│   ├── 📁 05-space-based/                  # Remove DB, ultra-fast processing on RAM grid
│   ├── 📁 06-service-mesh-planes/          # Control Plane (brain) & Data Plane (execution)
│   ├── 📁 07-bff-backend-for-frontend/     # Separate APIs for each client (Web/Mobile)
│   └── 📁 08-saga-pattern/                 # Distributed transaction management
│
├── 📁 3-application-architecture/          # HOW CODE IS ORGANIZED INTERNALLY (Software Arch)
│   ├── 📁 01-ui-patterns/                  # UI Layer Architectures (Frontend/Mobile)
│   │   ├── 📁 mvc-pattern/                 # Model - View - Controller (Classic)
│   │   ├── 📁 mvp-pattern/                 # Model - View - Presenter (Strict)
│   │   ├── 📁 mvvm-pattern/                # Data Binding (Modern Web/App)
│   │   └── 📁 viper-pattern/               # 5-layer decomposition for large projects
│   ├── 📁 02-layered-architecture/         # Controller - Service - Repository (Standard)
│   ├── 📁 03-pipeline-architecture/        # Sequential data processing (Pipe & Filter)
│   ├── 📁 04-microkernel-architecture/     # Core system + plug-ins/extensions
│   ├── 📁 05-clean-architecture/           # Clean arch: hexagonal with more circles
│   ├── 📁 06-hexagonal-architecture/       # Onion/Hexagonal: logic fully independent of DB
│   └── 📁 07-cqrs/                         # Separate write and read paths completely
│
├── 📁 4-resilience-patterns/               # HOW SYSTEMS SURVIVE (Stability - *Reference*)
│   ├── 📁 01-retry-strategy/               # Smart retries (Exponential Backoff)
│   ├── 📁 02-circuit-breaker/              # Prevent cascading failures (Safety switch)
│   ├── 📁 03-singleflight/                 # Deduplicate identical requests (Optimization)
│   ├── 📁 04-bulkhead/                     # Resource isolation (Ship compartment model)
│   ├── 📁 05-rate-limiting/                # Limit request rate (Throttling)
│   └── 📁 06-timeout-pattern/              # Set waiting time limits (Deadline)
│
└── 📁 5-data-architecture/                 # HOW DATA IS MANAGED (Database Strategy)
    ├── 📁 01-database-replication/         # Master - Slave (replication for faster reads)
    ├── 📁 02-database-sharding/            # Split data across multiple servers (Sharding)
    ├── 📁 03-cap-theorem/                  # Consistency - Availability - Partition tradeoff
    └── 📁 04-event-sourcing/               # Store all changes as an immutable event log
```
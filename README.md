# Distributed Tracing by Datadog with integrated logging on Golang
## Flow-Diagram
```mermaid
flowchart TB
    subgraph "Application Services"
        Service1["Service 1\n(with Tracing Library)"]
        Service2["Service 2\n(with Tracing Library)"]
        Service3["Service 3\n(with Tracing Library)"]
    end

    subgraph "Instrumentation Layer"
        TraceLib1["Datadog Tracer SDK\n(APM Library)"]
        TraceLib2["Datadog Tracer SDK\n(APM Library)"]
        TraceLib3["Datadog Tracer SDK\n(APM Library)"]
    end

    subgraph "Local Collection"
        Agent["Datadog Agent\n(Trace Collection)"]
    end

    subgraph "Datadog Backend"
        Intake["Trace Intake"]
        Storage["Trace Storage"]
        
        subgraph "Processing Pipeline"
            Sampling["Trace Sampling"]
            Aggregation["Trace Aggregation"]
            Analytics["Trace Analytics Engine"]
        end
        
        subgraph "Visualization Layer"
            Dashboard["Datadog APM Dashboard"]
            ServiceMap["Service Map"]
            TraceSearch["Trace Search & Analytics"]
            FlameGraph["Flame Graphs"]
        end
    end

    Service1 --> TraceLib1
    Service2 --> TraceLib2
    Service3 --> TraceLib3
    
    TraceLib1 --> Agent
    TraceLib2 --> Agent
    TraceLib3 --> Agent
    
    Agent --> Intake
    
    Intake --> Sampling
    Sampling --> Aggregation
    Aggregation --> Analytics
    Aggregation --> Storage
    
    Analytics --> Dashboard
    Storage --> ServiceMap
    Storage --> TraceSearch
    Storage --> FlameGraph
    
    %% Request Flow
    Service1 -- "Request/Response\n(with trace context)" --> Service2
    Service2 -- "Request/Response\n(with trace context)" --> Service3
    
    %% Add some styling
    classDef services fill:#d0e0ff,stroke:#3333dd,stroke-width:2px
    classDef traceLib fill:#ffe0b0,stroke:#ff9900,stroke-width:1px
    classDef agent fill:#d0ffe0,stroke:#00cc66,stroke-width:2px
    classDef backend fill:#ffe0d0,stroke:#ff6600,stroke-width:1px
    classDef viz fill:#f0d0ff,stroke:#9900cc,stroke-width:1px
    
    class Service1,Service2,Service3 services
    class TraceLib1,TraceLib2,TraceLib3 traceLib
    class Agent agent
    class Intake,Storage,Sampling,Aggregation,Analytics backend
    class Dashboard,ServiceMap,TraceSearch,FlameGraph viz
```

## Credits
This flow-diagram was generated with assistance from Claude, an AI assistant by Anthropic.
Code generated on: March 20, 2025

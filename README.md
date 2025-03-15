# Bilan Rekar Proto

This repository contains Protocol Buffer definitions for Bilan Rekar microservices.

## Project Structure

```
proto/
├── common/              # Common definitions shared across services
│   └── v1/
│       ├── error.proto  # Common error definitions
│       └── types.proto  # Common types (Money, Address, etc.)
├── user/                # User service
├── order/              # Order service
├── product/            # Product service
└── payment/            # Payment service
```

## Development

### Prerequisites

- [Protocol Buffer Compiler](https://github.com/protocolbuffers/protobuf)
- [Buf](https://buf.build)
- Go 1.21 or later

### Generate Code

To generate code from proto files:

```bash
buf generate
```

### Adding New Service

1. Create a new directory under `proto/` for your service
2. Add your `.proto` files in the `v1/` subdirectory
3. Run `buf generate` to generate code
4. Commit both `.proto` files and generated code

### Best Practices

1. Use versioning (v1, v2, etc.) for all services
2. Keep messages and services focused and single-purpose
3. Reuse common types from `common/v1/types.proto`
4. Follow Protocol Buffer style guide
5. Add proper documentation for all messages and services 
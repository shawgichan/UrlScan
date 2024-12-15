# UrlScan 🚧 (Work in Progress)

## Project Status: Active Development 🛠️

UrlScan is a web application for scanning multiple URLs, integrating PowerDNS, with advanced logging, metrics, and categorization capabilities.

## 🎯 Project Goals

- Develop a robust multi-URL scanning service
- Integrate PowerDNS for DNS resolution
- Implement comprehensive logging and monitoring
- Create a flexible URL categorization system

## 🏗️ Architecture (Draft)
```
UrlScan/
│
├── main.go         # Primary application launcher
│
├── internal/           # Internal packages
│   ├── handler/        # Request handlers
│   └── └── handler.go  # URL scanning logic
│
├── docker/             # Containerization
│   └── Dockerfile      
│
├── configs/            # Configuration files
│
└── README.md           # Project documentation
```

## 🔍 Detailed Feature Specifications

### 1. Multi-URL Scanning
- Support GET endpoint with multiple URL parameters
- Batch processing of URL scans
- Individual and aggregate results

### 2. URL Categorization
- Manual category mapping
  - Use of key-value store (map) for domain categories
- Initial focus on Top-Level Domain (TLD) categorization
- Future plans for subdomain-level categorization

### 3. PowerDNS Integration
- Containerized PowerDNS setup
- Supervisor to manage multiple services
- Reliable DNS resolution for URL scanning

### 4. Logging
- Zap logger implementation
- Structured logging
- Configurable log levels
- Performance-optimized logging

### 5. Metrics
- Prometheus metrics integration
- Tracking of:
  - Scan request counts
  - Response times
  - Error rates
  - DNS resolution performance

### 6. Testing Strategy
- Unit tests for individual components
- Integration tests
- Mock implementations
- Coverage reporting

### 7. Performance
- Benchmarking of key components
- Performance profiling
- Optimization analysis

## 💻 Development Roadmap

### Phase 1: Core Functionality
- [x] Initial project setup
- [x] PowerDNS Docker integration
- [ ] Multi-URL scanning endpoint
- [ ] Basic URL categorization
- [ ] Zap logger implementation

### Phase 2: Advanced Features
- [ ] Comprehensive categorization
- [ ] Prometheus metrics
- [ ] Extensive test coverage
- [ ] Performance benchmarks

### Phase 3: Refinement
- [ ] Advanced categorization strategies
- [ ] Performance optimizations
- [ ] Production-readiness improvements

## 🛠️ Development Setup

### Prerequisites
- Go 1.21+
- Docker

## 🚧 Current Limitations
- Early development stage
- Basic categorization
- Limited DNS capabilities
- No authentication mechanism


### Development Principles
- Test-driven development
- Performance consciousness
- Clean, readable code
- Comprehensive documentation

## 📝 Immediate TODOs
- [ ] Implement multi-URL scanning
- [ ] Create category mapping system
- [ ] Set up Zap logger
- [ ] Integrate Prometheus metrics
- [ ] Develop comprehensive test suite
- [ ] Create performance benchmarks
- [ ] Dockerize with Supervisor

## 🔬 Future Investigations
- Machine learning-based categorization
- Advanced DNS resolution strategies
- Distributed scanning capabilities

## 🔒 License
[To be determined]

---

**Disclaimer**: Project under active development. Interface and features subject to significant changes.

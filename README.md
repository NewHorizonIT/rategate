# RateGate — Multi-Tenant API Gateway with Distributed Rate Limiting

## Overview

**RateGate** is a high-performance, multi-tenant API Gateway that enforces **distributed rate limiting** using Redis.

It is designed to simulate real-world infrastructure systems like Stripe or GitHub API gateways, focusing on:

- Scalability
- Concurrency safety
- Distributed coordination
- Performance under load

This project demonstrates how to build a production-grade backend system beyond basic CRUD applications.

---

## Problem Statement

In modern systems, APIs must:

- Handle high traffic from multiple clients (tenants)
- Enforce fair usage (rate limiting)
- Remain consistent across multiple instances
- Avoid performance bottlenecks

Naive implementations often fail due to:

- Race conditions
- Inconsistent limits across instances
- Poor performance under load

**RateGate solves these problems with a distributed, Redis-backed design.**


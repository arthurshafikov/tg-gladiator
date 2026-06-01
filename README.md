# TG Gladiator

A Telegram RPG bot built with Go featuring turn-based battles, gear, a shop system, hero progression and events.

This is a production-style pet project with layered architecture, tests, Docker-based environment setup, and scalable update processing.

## What This Project Demonstrates

- Strong Go backend development (clean structure, service layer, database integration).
- Domain and business logic design (heroes, items, battles, progression).
- External API integration (Telegram Bot API).
- Concurrency and async processing (update handling, worker pool, per-user request synchronization).
- Unit and integration testing practices.
- Real-world localization, configuration, and containerized deployment workflow.

## Core Features

- Hero creation and management.
- PvE battles across multiple locations.
- Equipment and in-game shop system.
- XP/level progression with stat bonus distribution.
- Telegram interaction flow handling (commands, callback queries, user scenarios).
- Message localization (EN/RU).

## Tech Stack

- Go
- PostgreSQL
- Telegram Bot API
- Docker, Docker Compose
- GolangCI-Lint
- Unit and integration tests

## Architecture

The project is organized by layers:

- cmd: application entrypoints.
- internal/core: domain models, constants, types, errors.
- internal/services: business logic.
- internal/repository: data access (PostgreSQL, mocks).
- internal/transport/telegram: incoming Telegram update handling.
- internal/events and internal/scheduler: events and background jobs.
- migrations and data.sql: schema and seed data.

This structure makes the codebase easier to maintain, test, and scale.

## Demo

YouTube demo video:

- https://www.youtube.com/watch?v=fgbaNs3_fUA

What is shown in the demo:

- Hero creation flow
- Turn-based combat interaction
- Item/shop interaction in Telegram UI

## Quick Start

Run the project:

	make up

Stop the project:

	make down

## Testing and Quality

Run unit tests:

	make test

Run integration tests:

	make integration-tests

Run linters:

	make lint

## Potential Next Steps

- Metrics and observability (Prometheus/Grafana).
- CI pipeline (automated tests, linting, build).
- Load testing for update processing.
- Expanded gameplay systems (guilds, PvP, tournaments).

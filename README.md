# MaistroApp: Automotive Workshop Management System

MaistroApp is a high-performance, monolithic management solution specifically designed for automotive workshops. Built with Go and Astro, the application is engineered to operate efficiently on legacy hardware and low-resource environments, specifically optimized for systems with 4GB of RAM.

The project name, "Maistro," pays tribute to the traditional Mexican term for a master mechanic, combining decades of practical expertise with modern software engineering.

## Technical Architecture

The application follows a minimalist, single-binary architectural pattern to ensure the lowest possible memory footprint and simplified deployment:

* Backend: Go (Golang) utilizing the standard library and raw SQL (database/sql) to eliminate the overhead of heavy ORMs.
* Database: SQLite 3 with Write-Ahead Logging (WAL) mode enabled for high-concurrency performance in a serverless format.
* Frontend: Astro (Static Site Generation) to serve pre-rendered HTML/CSS, significantly reducing client-side JavaScript execution.
* Distribution: Single-binary deployment using the Go embed package to bundle frontend assets directly into the executable.

## Core Functionality

* Service Lifecycle Management: Real-time tracking of vehicle status from initial diagnosis to final delivery.
* Comprehensive Vehicle History: Centralized database to track repairs, spare parts, and maintenance history by VIN or license plate.
* Automated Budgeting: Calculation tools to separate labor costs from part expenses, providing clear profit margin analysis.
* WhatsApp Integration: Automated generation of service updates and budget approvals via pre-formatted messaging links.
* Memory Optimization: Aggressive memory management and connection pooling to ensure system stability on 4GB RAM hardware.

## Installation and Build Process

### Prerequisites
* Go 1.21 or higher
* Node.js (Required only for the frontend build stage)

### Build Steps

1. Build the Frontend Assets:
   ```bash
   cd frontend
   pnpm install
   pnpm run build
   ```
2. Compile the Monolith:
   Navigate back to the root directory and compile the Go binary 
   ```bash
   go build -ldflags"-s -w" -o maistroapp 
   ```
3. Execution:
   Run the generated binary. The application will initialize the SQLite database automatically   on the first launch.
   ```bash
   ./maistroapp 
   ```

## Project Structure

* `/frontend`: Astro source code and UI components.
* `/internal/repository`: Raw SQL logic and SQLite interaction.
* `/internal/handlers`: API endpoints and business logic.
* `/internal/models`: Data structures and schema definitions.
* `main.go`: Entry point and static file embedding logic.

## Performance Configuration

To maintain a minimal footprint, the application is configured with:

* **PRAGMA busy_timeout = 5000**: To handle concurrent write attempts without locking errors.
* **MaxOpenConns(1)**: To prevent database file descriptors from exhausting system memory in limited environments.
* **go:embed**: To eliminate the need for an external web server like Nginx or Apache, serving assets directly from memory.

## Development and Deployment

The application is distributed as a single binary. To ensure optimal performance on 4GB RAM systems, it is recommended to:

1. Use the `-ldflags="-s -w"` flag during the Go build process to reduce binary size.
2. Store images in a local directory and save only the file paths in SQLite to prevent database bloat.
3. Enable Zswap or a similar compressed swap mechanism on the host operating system.

## License

MIT

# Groupie Trackers

Groupie Trackers is a small Go web application that consumes the public [Groupie Trackers API](https://groupietrackers.herokuapp.com/api) to display bands, their members, concert locations, and upcoming dates through a clean HTML interface.

## Features
- Fetches artists, locations, dates, and relation data directly from the Groupie Trackers API.
- Presents a responsive gallery of artists with names, images, debut info, and member counts.
- Provides a rich detail page per artist including locations, dates, and a grouped concert schedule.
- Serves static assets (CSS, images) alongside rendered Go HTML templates.
- Gracefully handles common HTTP errors with an application-specific error page.

## Requirements
- Go 1.20 or newer (module currently targets Go 1.24.5).
- Internet connection (the server pulls fresh data from the external Groupie Trackers API at startup).

## Getting Started
```bash
git clone <repository-url>
cd groupie-tracker/backend
go run .
```

Once the server starts, visit [http://localhost:8080](http://localhost:8080) in your browser. The terminal output includes the working directory and confirmation that the listener is running.

### Building a Binary
```bash
cd groupie-tracker/backend
go build -o groupie-tracker
./groupie-tracker
```

## Project Structure
```
assets/      Static files served with the app (CSS, images, error assets)
backend/     Entry point, HTTP handlers, template initialization, error helpers
models/      Go types plus API client helpers that populate global data
templates/   HTML templates for the index, artist details, and error views
```

Key runtime flow:
1. `backend/main.go` calls `models.Load*` helpers to hydrate shared data from the external API.
2. Handlers in `backend/handlers.go` render `templates/*.html` using this in-memory state.
3. Static assets are exposed under `/assets/` via `http.FileServer`.

## Development Notes
- Run `go fmt ./...` before committing changes to keep formatting consistent.
- The API client lives in `models/api.go`; update the constants there if the remote endpoints ever move.
- Template rendering uses relative paths (`../templates` and `../assets`), so run commands from the `backend` directory when executing binaries locally.
- If the remote API is unreachable, the server will exit at startup with the error surfaced in the logs.

## Future Enhancements
- Add client-side search or filtering for the artist list.
- Cache API responses locally or add scheduled refresh support.
- Expand the error handling pages with more context or retry guidance.

## License
This project currently does not declare a license. Add a LICENSE file if you intend to distribute it.

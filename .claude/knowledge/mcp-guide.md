# MCP Servers Guide

Model Context Protocol (MCP) servers extend Claude Code's capabilities. This project is pre-configured with several MCP servers in `.claude/settings.json`.

## Context Management

Long sessions degrade in quality as the context window fills. Several mechanisms counter this:

- **`autoCompact`** (enabled in settings.json) — Claude Code automatically summarizes older conversation history when approaching context limits, preserving the active working context
- **`memory` MCP** — persist facts, decisions, and preferences to a knowledge graph that survives between sessions; agents retrieve only what's relevant rather than re-establishing everything from scratch
- **`sequential-thinking` MCP** — structures complex reasoning into explicit steps, avoiding the wasted context of unstructured exploration
- **Scoped agents** — each agent loads only its domain's rules and knowledge, not the entire project context
- **`.claude/memory/`** — file-based memory for durable project facts; see `MEMORY.md` for the index

## Installed MCP Servers

### context7 — Library Documentation
**Purpose**: Get up-to-date documentation and code examples for any library.

**When to use**:
- "How do I use pgx/v5 batch queries?"
- "What's the React Query v5 syntax for infinite queries?"
- "Show me Caddy TLS configuration examples"

**Usage**: Claude automatically uses this when it needs current library docs. You can also ask explicitly: "Use context7 to find the latest Chi middleware documentation."

**Setup**: Pre-configured, requires `npx` (included with Node.js).

---

### github — GitHub Integration
**Purpose**: Search code, manage issues and PRs, access repository information.

**When to use**:
- Searching for code examples across GitHub
- Creating issues from bug reports
- Reviewing PR status
- Accessing private repository information

**Setup**:
1. Create a GitHub Personal Access Token at github.com/settings/tokens
2. Set environment variable: `export GITHUB_TOKEN=ghp_xxxx`
3. Or add to your shell profile for persistence

---

### brave-search — Web Search
**Purpose**: Search the web for current information, documentation, and solutions.

**When to use**:
- Finding solutions to errors not in documentation
- Researching new libraries or tools
- Getting current information (post-training cutoff)

**Setup**:
1. Get API key at brave.com/search/api/
2. Set environment variable: `export BRAVE_API_KEY=BSA_xxxx`

---

### postgres — Direct Database Access
**Purpose**: Query your PostgreSQL database directly for debugging and exploration.

**When to use**:
- Debugging data issues
- Exploring schema during development
- Running ad-hoc queries without psql

**Setup**:
1. Set environment variable: `export DATABASE_URL=postgresql://user:pass@localhost:5432/myapp`
2. Ensure database is accessible from your machine

**Security note**: Only use with development databases. Never point at production.

---

## Installing MCP Servers

All configured MCP servers use `npx` and install automatically on first use. No manual installation needed.

To verify MCP servers are working:
```bash
claude mcp list  # Show configured servers
```

## Adding New MCP Servers

To add a new MCP server, edit `.claude/settings.json`:

```json
{
  "mcpServers": {
    "my-new-server": {
      "command": "npx",
      "args": ["-y", "@package/server-name"],
      "env": {
        "API_KEY": "${MY_API_KEY}"
      },
      "description": "What this server does"
    }
  }
}
```

## Popular MCP Servers for SaaS Development

| Server | Package | Use Case |
|--------|---------|----------|
| Filesystem | `@modelcontextprotocol/server-filesystem` | Enhanced file operations |
| Slack | `@modelcontextprotocol/server-slack` | Send notifications |
| Stripe | Community servers | Payment API integration |
| Resend | Community servers | Email API integration |
| Linear | Community servers | Issue tracking |
| Notion | Community servers | Documentation sync |

Browse more at: modelcontextprotocol.io/servers

---

### memory — Persistent Knowledge Graph
**Purpose**: Store and retrieve facts, decisions, and context across sessions without consuming the conversation window.

**When to use**:
- Remembering architectural decisions made in previous sessions
- Storing user preferences Claude should always apply
- Keeping track of which features are in progress

**Setup**: Pre-configured, no API key needed.

**Usage**: Ask Claude to remember/recall/forget things. It manages the graph automatically.

---

### sequential-thinking — Structured Reasoning
**Purpose**: Break complex tasks into explicit reasoning steps before acting.

**When to use**:
- Debugging non-obvious issues
- Planning a multi-file refactor
- Any task where jumping straight to code risks going in the wrong direction

**Setup**: Pre-configured, no API key needed.

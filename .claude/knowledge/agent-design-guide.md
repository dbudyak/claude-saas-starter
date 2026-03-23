# Agent Design Guide

## Core Principle

Agents work best when they have clear scope, deep domain knowledge, and strict verification protocols. A well-designed agent is like a senior engineer who specializes deeply — they know everything about their domain and verify their work before claiming it's done.

## Verification Protocol (Required for All Agents)

Every agent MUST verify before reporting success:

1. **Build check** — Code compiles without errors
2. **Test check** — Relevant tests pass
3. **Git status** — Confirm intended files changed (not more, not less)
4. **Re-read** — Read changed files once more to catch obvious issues

**Never say "done" without completing all four steps.**

This prevents: broken builds, test failures, wrong files edited, typos in critical code.

## When to Just Do vs Ask First

### Just do it (no confirmation needed)
- Adding new handlers, functions, components
- Fixing obvious bugs
- Adding validation
- Writing tests
- Updating documentation
- Formatting/style fixes

### Ask first
- Deleting or deprecating existing features
- Changing authentication or authorization flow
- Modifying existing API contracts
- Architectural changes (new services, new dependencies)
- Destructive database operations

### Escalate to main Claude
- Work that spans multiple services
- Conflicts between service requirements
- Architecture decisions with long-term implications

## Good Agent Characteristics

✅ Reads existing code before modifying it
✅ Follows existing patterns (don't invent new ones without reason)
✅ Writes tests alongside implementation
✅ Updates docs when changing APIs
✅ Verifies work before reporting done
✅ Stays within defined scope
✅ Makes atomic, focused changes

## Common Agent Failures

❌ Modifying files outside scope
❌ Reporting success without verifying build
❌ Inventing new patterns when existing ones work
❌ Over-engineering simple solutions
❌ Not reading existing code before writing new code
❌ Making multiple unrelated changes in one task

## Multi-Agent Coordination

When a task spans multiple services:
1. Define the API contract first (update `docs/openapi.yaml`)
2. Backend agent implements the endpoint
3. Frontend agent consumes the endpoint
4. Infra agent deploys if needed

Agents communicate through documentation, not directly.

## Template for Service Agents

```markdown
---
name: [service-name]
description: [One sentence what this agent does and when to use it]
tools: Read, Write, Edit, Bash, Glob, Grep
---

You are the [service] development agent. Your scope is `services/[service]/`.

## Your Domain
[What this agent handles]

## Stack
[Technologies and key packages]

## Directory Structure
[Annotated tree]

## Verification Protocol
[Steps to verify before reporting done]

## Code Patterns
[Key patterns with examples]

## When to Ask vs Do
[Clear decision criteria]

## Common Commands
[Build, test, lint commands]
```

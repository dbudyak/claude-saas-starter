# Memory Index

This file is the index for project memory. Claude reads it at the start of each session to restore context that would otherwise be lost between conversations.

Memory files live in this directory. Each file covers one topic. This index links to them with a one-line description.

## How to use

Ask Claude to remember something: *"Remember that we're using soft deletes for user records"*
Ask Claude to recall something: *"What did we decide about the email provider?"*
Ask Claude to forget something: *"Forget the note about the old auth approach"*

## Memory files

<!-- Claude maintains this list. Example format:
- [decisions/auth-approach.md](decisions/auth-approach.md) — chose Resend over SES for simplicity
- [user/preferences.md](user/preferences.md) — prefers minimal code, dislikes boilerplate
-->

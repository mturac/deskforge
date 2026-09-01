# deskforge

**Your agent agency, in your repo.**

Grok Bot–shaped multi-agent desk you can `init` locally and paste into [Grok Bot](https://docs.x.ai/docs/guides/grok-bot) — without host lock-in.

## Why

People already pay for persistent multi-agent desks (Grok Bot, Construct, Skydive, Cowork, …). deskforge copies the **job** (narrow bots + group loop + human gates), cuts hosted lock-in, and makes the **git ledger** the product.

## Install

```bash
go install github.com/mturac/deskforge@latest
```

Or clone and build:

```bash
git clone https://github.com/mturac/deskforge
cd deskforge && go build -o deskforge .
```

## First value (<10 minutes)

```bash
mkdir my-desk && cd my-desk
git init
deskforge init
deskforge agents
```

1. Create six Grok Bots: **Mainer, Pulse, Scout, Judge, Builder, Critic**
2. Paste each `agents/*.md` fenced Description into the matching bot
3. Put them in one group with a shared computer that can `git` this folder
4. Paste the kickoff from `HOWTO.md`

You now have a runnable desk loop — not a markdown souvenir.

## What init writes

| Path | Purpose |
| --- | --- |
| `agents/` | Grok Bot Description contracts (paste-ready) |
| `desks/` | DEMAND, PREFLIGHT, LAUNCH |
| `ARCHITECTURE.md` | hierarchy + control loop |
| `APPROVAL.md` | autonomous vs human-gated |
| `WATCHLIST.md` | Pulse sources |
| `ideas/`, `vault/`, `tickets/` | ledger |

## Non-goals

- Not a coding-agent IDE workspace
- Not a Firecracker/micro-VM runtime
- Not “Grok Bot but free” branding

## License

MIT

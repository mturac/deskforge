# HOWTO — Grok Bot desk in your repo

## What you get

`deskforge init` writes a git-native agency:

- 6 agent contracts under `agents/` (paste into Grok Bot Descriptions)
- desks: DEMAND, PREFLIGHT, LAUNCH
- WATCHLIST, APPROVAL, ARCHITECTURE
- empty `ideas/`, `vault/signals/`, `vault/demands/`, `tickets/`

## Install

```bash
go install github.com/mturac/deskforge@latest
# or from this repo:
go build -o deskforge .
```

## Init

```bash
mkdir my-desk && cd my-desk
git init
deskforge init
deskforge agents
```

## Wire Grok Bot

1. Create 6 bots: Mainer, Pulse, Scout, Judge, Builder, Critic
2. Open each `agents/<NAME>.md` and paste the fenced Description block into that bot’s Description
3. Put the bots in one group chat with a shared computer that can `git` this repo
4. Paste the kickoff below

### Group kickoff

```text
@Mainer Run this desk as a private full-service OSS agency.
Repo: <this git remote>

Lanes ON: household + challenger.
Challenger twins of loved closed products are IN SCOPE (including GrokBot-shaped open desks). Do not soft-kill to protect vendors.

Autopilot ON: Pulse → Scout → Judge → brief.
Autopilot OFF: status=build, public, Product Hunt, tweets, spend — unless I said build <slug> or unlock public.

Loop:
1) @Pulse WATCHLIST + PH/GH radar → vault/signals/ OR nothing_today
2) @Scout household wedge OR challenger copy/cut/add memo
3) @Judge score; refuse brief without PREFLIGHT pass
4) @Mainer ranked brief daily OR nothing_today
5) Builder only after: build <slug>

Start: read ARCHITECTURE.md APPROVAL.md desks/*; next cycle; brief me.
```

## Your rhythm

- Read briefs
- Decide: `build <slug>` / `hold` / `kill`
- Ask for usable kits (agent pack updates, skills, init) — agents should ship them, not only scores

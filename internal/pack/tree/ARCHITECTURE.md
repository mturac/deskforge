# deskforge architecture

Pattern: narrow agents, tickets, heartbeats, producer ≠ auditor, silence is valid, irreversible = human.

Job: **your agent agency, in your repo** — GrokBot-shaped desk without host lock-in.

---

## Hierarchy

```
Human (brand, money, public, secrets)
── Mainer                         # routes + files DEMANDS from chat
    ── Pulse                      # WATCHLIST + launch radar
    ── Scout                      # ideas + PREFLIGHT
    ── Judge                      # scores + may issue DEMANDS
    ── Builder                    # promoted only
    ── Critic                     # launch gate
```

---

## Control loop

```text
DEMAND (steer) → Pulse radar / Scout hunt → PREFLIGHT → Judge score → human brief
        ↑                                         |
        └──── reject codes + next hunt ───────────┘
```

- Judge 0–100 is **not** approval.
- PREFLIGHT PASS is required before human brief.
- Open p0 DEMAND blocks fake `nothing_today`.

---

## Scores

Bands: 0–39 archive · 40–69 incubate · 70–84 brief · 85–100 recommend-build (human gate).

Only human `build <slug>` unlocks Builder.

---

## File map

| Path | Role |
| --- | --- |
| `desks/DEMAND.md` | steer bus |
| `desks/PREFLIGHT.md` | brief gate |
| `desks/LAUNCH.md` | PotD / trending bar |
| `APPROVAL.md` | what agents do alone vs human |
| `WATCHLIST.md` | Pulse sources |
| `agents/*` | Grok Bot Description contracts |
| `ideas/` | ranked bets |
| `vault/signals/` | Pulse output |
| `vault/demands/` | open hunts |
| `tickets/` | work orders |

---

## Limits

- Cap ~6 agents
- Empty reports forbidden when radar was required
- Do not soft-kill twins to protect closed vendors

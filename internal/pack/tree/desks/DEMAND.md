# DEMAND bus — steer the agency

Humans or Judge can **order a hunt**. Scout/Pulse must answer. Folding everything into one mega-slug is **not** an answer unless the demand says so.

## File a demand

Path: `vault/demands/D-NNN-<slug>.md`

```yaml
id: D-NNN
title: <one line>
issuer: human|mainer|judge|cursor
status: open|answered|expired|killed
priority: p0|p1|p2
expires: YYYY-MM-DD
kind: twin|household|skill|plugin|radar|kit|other
must_produce: idea|signal|skill_pr|agent_pack|nothing_with_reason
```

Body:

1. **Bring:** what to find or ship
2. **Not:** what is forbidden
3. **Done when:** concrete artifact path
4. **Window:** why now

## Rules

1. Open `p0` demands block `nothing_today` for Scout.
2. Judge may open a demand instead of soft-briefing trash.
3. Human one-liners in chat → Mainer files the demand same cycle.
4. Answer with `answered:` path or `blocked:` reason + next hunt.
5. Kind `kit` means ship something a human can run or paste (agents, skill, CLI) — not only a score.

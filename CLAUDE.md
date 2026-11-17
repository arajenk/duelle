# duelle

## context
- Read SPEC.md (repo root) before planning anything. It's the source of truth.
- frontend/DESIGN.md (if it exists) covers UI only. Read it for frontend work.
- Anything SPEC.md already decides is final unless something's actually broken.
- Features get their own spec in specs/<feature>.md. Plan from that file.
- When a feature ships, add a line or two about it to SPEC.md.
- Write tests from SPEC.md or specs/, not from the implementation.

## workflow
- New feature: draft specs/<feature>.md from SPEC.md (what + decisions, short), with me leading. I approve it before any planning.
- Then plan first from the spec (how: files, functions, order). Wait for my approval. Never implement straight from "implement X".
- Work one slice at a time. A feature is split into slices that fit the diff size below.
- After each slice, tick it off in the spec and add any decisions I made.
- Every plan lists the decisions I need to make, with options and tradeoffs.
- For anything non-trivial and real design decisions, don't pick for me. Make me pick.
- Small stuff like naming or file placement, just choose.
- After each change, show evidence: test output, the commands you ran, and code snippets.
- Tell me if you deviated from the plan, and where.

## who writes what
- Specs tag items `(me)` when I'm hand-coding them. Untagged = you write it.
- On `(me)` items: explain the approach in words and name the functions I'll need. No code, no skeletons. Review my implementation, use hints first.
- After your slices, ask me 1-2 quick questions about the code.

## diff size
- Backend: one function or ~50 lines per change. Max 150.
  If more, explain why and wait for approval.
- Tests: max 200 lines. If more, explain why.
- Frontend can go bigger, but never touch backend files during
  frontend work without asking.

## comments
- One short high-level comment above a function or block when the what or why
  isn't obvious. No line-by-line comments or docstring-style blocks. Lowercase.

## help me learn
- Flag unfamiliar patterns or libraries with a one-line explanation.
- When something breaks, give me a hint first so I can figure out the reason myself.
  Only fix it if I say "just fix it".

## always protect
- Server is authoritative. Client never gets a word it hasn't solved.

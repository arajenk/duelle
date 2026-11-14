# duelle

## context
- Read SPEC.md (repo root) before planning anything. It's the source of truth.
- frontend/DESIGN.md (if it exists) covers UI only. Read it for frontend work.
- Anything SPEC.md already decides is final unless something's actually broken.
- Features get their own spec in specs/<feature>.md. Plan from that file.
- When a feature ships, add a line or two about it to SPEC.md.
- Write tests from SPEC.md or specs/, not from the implementation

## workflow
- Plan first, then wait for my approval. Never implement straight from "implement X".
- Work one slice at a time. A feature is split into slices that fit the diff size below.
- Every plan lists the decisions I need to make, with options and tradeoffs.
- For anything non-trivial and real design decisions, don't pick for me. Make me pick.
- Small stuff like naming or file placement, just choose.
- After each change, show evidence: test output, the commands you ran, and code snippets.
- Tell me if you deviated from the plan, and where.

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
<!-- 1-3 things that would break this project if done wrong. Flag any change that touches them. -->
- 
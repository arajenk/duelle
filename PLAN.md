# duelle v0 plan

Working plan for getting v0 (see DESIGN.md) built. DESIGN.md says *what* the game is; this file says *what to build next, in what order, and what's still undecided*.

## How to use this

- **Claude Desktop** = study partner. Paste in one step at a time. It explains Go concepts with a tiny standalone (non-duelle) example first, I write the code, and I paste back small diffs for review. It doesn't write the core code for me.
- **Claude Code** = comes in after big milestones to update this plan, review what landed, and, later on, when I decide to let AI write code.
- **"Decided"** = choices I've already made. Don't re-argue them unless something's actually broken.
- **"Open"** = my call. Desktop should ask how I'd do it before suggesting anything.
- Each step should be one small diff that `go build` / `go test ./...` passes after.
- When a milestone is done: tick it, write one line under "Log" at the bottom, commit.

## Status

| # | Milestone | State |
|---|-----------|-------|
| 0 | Scoring + validate + terminal loop | done (score.go, main.go) |
| 1 | Room types + registry (`room.go`) | **next**, design decided |
| 2 | Word lists | not started, open |
| 3 | PROTOCOL.md (messages + phases) | not started, open |
| 4 | Websocket server + connections | partly decided |
| 5 | Match loop (trap → guess → round end → best of 5) | partly decided |
| 6 | v0 React frontend | not started (OK to delegate to AI) |
| – | Cleanup from earlier review | whenever |

Suggested order: 1 → 3 → 2 → 4 → 5 → 6. Word lists can slot in anywhere; they only need to exist by milestone 5. The order is mine to change.

---

## 1. Room types + registry (`room.go`)

**Goal:** the data model for a 1v1 room plus a registry to create, join and remove rooms. No networking yet; it's all testable with `go test`.

**Decided**
- Rooms are in memory only. A server restart loses games, and that's fine for v0.
- `Player`: nickname, token, this round's answer, this round's guesses (cleared each round), points (= rounds won), finish time as a `time.Duration` since round start.
- `Room`: 6-digit code, round number (`int`), round start time, players as `[2]*Player`. Pointers so changes stick, an empty slot is `nil` (fails loudly), and the same Player can be reattached on reconnect.
- Registry: `map[code]*Room` guarded by a `sync.Mutex`. It's only three quick operations, so a goroutine would be overkill.
- Create: generate a random **unused** 6-digit code.
- Join: lookup + full check + fill slot 2, all in **one** locked section (avoids check-then-act races). Returns errors: not found / full.
- Remove: delete the room at match end.
- Tokens: random from `crypto/rand` (not JWT).
- Everything stays in `package main` for now.
- The room goroutine is **not** started here. That's milestone 5.

**Go concepts to learn first** (tiny standalone examples)
- structs, methods, pointer receivers, `[2]*T` arrays
- maps + `sync.Mutex`, `defer mu.Unlock()`
- `crypto/rand` (random bytes → hex string; random 6-digit number)
- `errors.New` / sentinel errors, `errors.Is`
- `time.Time` vs `time.Duration`

**Small steps**
1. `Player` and `Room` structs. Builds.
2. `Registry` struct (map + mutex) + a constructor that makes the map.
3. Token generator + a test that two tokens differ and have the expected length.
4. Code generator (6 digits, keeps retrying until unused). Think about where the lock needs to be held.
5. `Create(nickname)` → room + host player. Test it.
6. `Join(code, nickname)` with the not-found / full errors. Tests: good join, bad code, full room.
7. `Remove(code)`. Test it.

**Open (small)**
- What does Create/Join return: room, player, token, or some combo?
- Should code generation live inside Create's lock or be its own thing?

**Done when:** `go test ./...` covers create / join / full / bad code / remove, and all pass.

---

## 2. Word lists

**Goal:** real lists to replace the hardcoded 10 words in main.go.

**Decided (DESIGN.md)**
- Built from ENABLE + frequency data, cleaned by hand. Not generated.
- Real, known words, but trappy ones (_ATCH clusters, double letters, JAZZY/KNACK). No obscure stuff.
- Earlier review note: keep **valid guesses** and **trap/answer candidates** as separate lists.

**Open**
- File format (one word per line? something else?) and where the files live.
- Ship them inside the binary or read them from disk at startup?
- What data structure for "is this a valid guess" lookups, and why?
- How the 3 trap choices get picked (repeats allowed within a match?).
- Is the cleaning process a script, by hand, or both?

**Go concepts:** `os.ReadFile` / `bufio.Scanner` on files, `embed` (`//go:embed`), maps as sets (`map[string]struct{}`), `math/rand/v2`.

**Done when:** `isValid` and answer picking use the real lists, and there are tests for a known-valid word, a known-invalid word, and lowercase handling.

---

## 3. PROTOCOL.md

**Goal:** write down every message between client and server, plus the phases a room goes through, *before* writing the websocket code. Words only, no Go.

**Decided**
- Backend sends `correct` / `present` / `absent`. The frontend picks colors.
- The client never receives a word it hasn't solved (the reveal happens at round end).
- Token sent on connect/reconnect ties the connection to a player.
- Server does all timing (`time.Now` on the server).

**Open (basically all of it)**
- Message format (JSON? what's the envelope, e.g. a `type` field?).
- Phases: what states does a room move through from created → match over?
- Client → server messages (create, join, pick trap, guess, …?).
- Server → client messages (room joined, trap choices, guess result, opponent progress, round over + reveal, match over, error, …?).
- What does a player see about the opponent mid-round? (DESIGN.md spectating rule: colors not letters.)
- What happens to an invalid guess: rejected without costing a guess?

**Done when:** the PROTOCOL.md walkthrough of one full best-of-5 match works on paper, with every message named.

---

## 4. Websocket server + connections

**Goal:** a browser (or a test client) can connect, create/join a room, and messages flow both ways.

**Decided**
- One **reader goroutine per connection**: reads messages and forwards them to the room over a channel.
- One **writer goroutine per player**: the only thing that writes to that socket, reading from an outbound channel (single-writer rule, keeps order).
- The outbound channel is **buffered**. If it's full, treat the player as disconnected (one slow client can't stall the room).
- Token on connect/reconnect identifies the player.

**Open**
- Which websocket library (or the standard library alone?), and why.
- HTTP routes: what endpoints exist (one `/ws`? separate create/join over HTTP?).
- What the room's incoming channel carries (raw bytes? parsed message + which player sent it?).
- What happens on disconnect in v0 (full afk/disconnect rules are v1).
- How to test it without a browser.

**Go concepts:** `net/http` handlers, goroutines, channels (buffered vs unbuffered), `select` with `default` (non-blocking send), `encoding/json`, `context` for shutdown.

**Done when:** two terminal/test clients can join the same room by code and echo messages through the room.

---

## 5. Match loop

**Goal:** the actual game. Trap pick → both solve → round result → repeat until best of 5 is decided.

**Decided**
- **One goroutine per room**, running the match top to bottom, sequentially. It starts once both players are in.
- Each round: each player gets 3 random words, has 15s to pick one for the opponent, and one is auto-picked if they're AFK.
- 5 letters, 6 guesses, scoring via `Score`.
- Round winner: fewest guesses, then fastest finish time. Both fail = draw.
- Best of 5 by rounds won. Tied match in v0 = draw (sudden death comes later).
- Reveal the trap word at round end.
- When the match ends, the room goroutine exits and removes itself from the registry.

**Open**
- How the loop waits for "both players picked" or "15s passed" at the same time.
- How it waits for both players to finish solving while still handling messages from each.
- Where round-winner logic lives. A plain function you can unit-test without goroutines would be the easy one to test.
- What "first to 3" vs "play all 5" means for when the match ends early.
- DESIGN.md wants modes as an interface. Introduce that now, or after duel works? (Points mode is v1.)

**Go concepts:** `select`, `time.After` / `time.NewTimer`, closing channels, interfaces (if doing modes now), testing with fake/injected time.

**Done when:** a full match plays through with test clients, and round-winner + match-winner logic has unit tests.

---

## 6. v0 frontend (React + Vite + Tailwind)

**Goal:** v0 screens only, no polish. "Ugly is fine."

**Decided:** React from day one. This part can be delegated to AI (come back to Claude Code for it).

**Open:** the screen list (home → create/join → trap pick → board → round result → match result?), and where `frontend/` lives in the repo.

**Done when:** two browsers can play a full best-of-5.

---

## Cleanup backlog (from the 2026-10-29 review)

- [ ] `Color` type (enum-style consts) instead of raw strings in `Score`
- [ ] Table-driven tests in score_test.go + more cases (all green, all grey, CRANE/EERIE)
- [ ] Consistent case handling (lowercase everywhere?)
- [ ] `go fmt` main.go
- [ ] Separate guess list vs answer list (folds into milestone 2)

---

## Not in v0 (don't get pulled in)

3+ players / points mode, sudden death, spectate, skip votes, accounts, Postgres, lobby lock, rate limiting, solo unlimited, theme polish. All in DESIGN.md under v1 / later.

---

## Log

- 2026-10-29: scoring + terminal loop done; room design decided.
- 2026-11-08: Room/Player types + registry design decided.
- 2026-11-13: PLAN.md written.

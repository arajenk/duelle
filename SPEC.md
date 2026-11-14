# duelle

word game where you trap your friends. you pick the word your opponent has to solve, they pick yours, fewest guesses wins.

skill wins over time, but anyone can steal a round with a good trap. every feature should protect this.

## core

- 5 letters, 6 guesses
- scoring is two passes (greens first, then yellows from leftover letters) so duplicates work right
- server holds the answers and does all scoring/timing, client never gets a word it hasn't solved

## trap words

- each round you get 3 random words and pick one for your opponent
- 15s to pick, auto picks if you're afk
- 3+ players: shuffle into a circle each round, you trap the person after you
- 3+ players: show "trapped by <name>" during the round
- reveal at the end of the round ("arajen trapped you with BATCH")

## theme

- medieval duel vibe - crimson, bronze-gold, stone

## rooms

- host makes a room, gets a pin or link, friends join with the link or pin + nickname (no account needed)
- 6 digit code so it's easy to read out on a call/to friends
- rate limit join attempts + host can lock the lobby
- host plays too, if they leave host goes to the next person
- best of 3 / 5 / 8, default 5

## accounts

- login with discord or google (no passwords to deal with)
- guests can still play rooms
- guests can upgrade to an account and keep their stats

## modes

- **2 players = duel** — round goes to fewest guesses, then fastest time. most rounds wins
- **3+ = points** — 6 pts for 1 guess down to 1 for 6, 0 for fail, +1 to first solver

## ties

- tied for 1st = sudden death, only the tied people play, same rules
- if everyone fails sudden death, fewest total guesses wins
- ties below 1st just share the spot

## spectating

- once you solve or run out you can spectate, `< name >` switcher at the bottom
- people still playing only see opponents' colors, not letters

## afk stuff

- skip vote shows up when 1–2 people are left. majority votes, they get 30s, then fail
- 1v1: can only skip if the other person hasn't typed for ~60s
- disconnect 30s+ = fail that round, can rejoin next round
- no round timer unless people start stalling

## solo

- unlimited: random words, streak + best streak, stats if logged in
- "challenge a friend" link after a round to send them the same word

## stack

- go backend (writing this myself, practice for torpor)
- websockets, postgres
- react + vite + tailwind frontend
- move to torpor when possible, figure out backend hosting when that bridge needs to be crossed
- modes as an interface so adding new ones doesn't mean rewriting stuff
- backend uses correct / present / absent, frontend decides colors
- PROTOCOL.md = client/server messages + room phases
- frontend/DESIGN.md = actual ui (skribbl-style sidebar on desktop, compact strip on phone)

## keep it original

- own name, own colors, own layout, own word list
- word list built from ENABLE + frequency data, cleaned by hand. not generated
- real words people know, but trappy ones (clusters like _ATCH, double letters, weird combos like JAZZY/KNACK)
- no super obscure stuff (if it doesn't feel english it's out), tune the frequency cutoff by playtesting
- later: solver ranks words by how hard they are to solve

## money

- free, no ads. cosmetics/supporter tier only if real players show up

## v0

- [x] score + validate
- [ ] word lists
- [ ] 1v1 only, room code
- [ ] trap words
- [ ] best of 5, fewest guesses wins
- [ ] nickname only, no accounts
- [ ] frontend in react from day one, only v0 screens, no polish

## v1

- [ ] rooms + lobby
- [ ] duel + points, best of 3/5/8
- [ ] sudden death
- [ ] spectate
- [ ] skip votes + disconnects
- [ ] unlimited
- [ ] accounts (optional)
- [ ] frontend

## later

- sudden death skull animation
- bots (easy → hard), ladder with a boss bot
- bots have two dials: how well they solve and how mean they trap
- game review (grade each guess)
- ghosts: every run saves guess + colors + time, race your own best / friends / randoms
- challenge a friend link = they race your ghost
- async challenges
- friends: recent players first, then friend codes
- modifiers (host toggles): random opener (same opener for everyone when compared, random in solo)
- daily rework: community trap — after the daily you pick tomorrow's word from 3, most picked wins
- if nobody votes, duellebot sets it. hide vote count until it's big
- progression: xp, medieval titles (peasant → squire → knight → lord → king), cosmetic unlocks only, nothing pay to win
- share cards after matches ("arajen trapped me with BATCH 💀")
- discord activity
- match history, achievements, suggestion box

## maybe

- gauntlet (not sold on it)
- random matchmaking if lobbies are ever busy
- ranked if it somehow blows up

## not doing

- other word lengths
- unlimited guesses
- blind mode
- hard mode
- plain daily
- draft mode
- mirror mode (that's just head-to-head wordle)
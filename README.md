# ♟️ Life Decision Tracker

A simple web application to record and evaluate the decisions we make in our daily lives. Perfect for personal logging, journaling, and self-reflection.

**Inspiration:** This project is inspired by the chess move evaluation system (like the one used on Chess.com). Every decision you make in life can be evaluated just like a move on a chessboard, ranging from a *Brilliant* move to a fatal *Blunder*.

## Key Features

- **Chess-Style Evaluation System:** Manually categorize your decisions based on chess evaluation levels. Example: *Feeling extremely lazy, but forcing yourself to finish a project until production -> Category: **Great**.*
- **Passwordless Access (Magic URL):** No need to register an account or remember passwords. The application uses a *Magic URL* (a unique, randomized URL) as the "key" to access your personal board. Just generate the URL and keep the link safe!
- **Anti-Spam (Rate Limiter):** Equipped with an IP-based rate limiter (1 board generation per hour) to prevent mass URL creation and spam.

## Tech Stack

This project is built using:
- **Backend:** [Go (Golang)](https://golang.org/) 
- **Framework:** [Gin Web Framework](https://gin-gonic.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** PostgreSQL
- **Frontend:** HTML Templates (Server-Side Rendering) + Vanilla CSS

## Decision Categories

The system will automatically seed the database with the following categories upon its first run:

1. 🌟 **Brilliant:** A rare decision with a massive, unexpectedly positive impact.
2. ❗ **Great:** A highly solid move that provides a significant advantage.
3. ⭐ **Best:** The most optimal choice available in that specific situation.
4. 👍 **Excellent:** A very good decision that strongly supports your main goals.
5. ✔️ **Good:** A correct step, although a slightly better option might have existed.
6. 📖 **Book:** Standard decisions, daily routines, or obligations that are expected to be done.
7. ❔ **Inaccuracy:** A suboptimal decision resulting in a slight waste of time or resources.
8. ❓ **Mistake:** A wrong choice with a noticeable negative impact.
9. ❌ **Miss:** Missing a critical opportunity or ignoring a top priority.
10. ⛔ **Blunder:** A fatal mistake that ruins progress or violates core principles.

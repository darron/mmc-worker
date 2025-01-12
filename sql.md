# SQL Commands through Wrangler

## Load the local SQLite for local development:

```bash
npx wrangler d1 execute massmurdercanada-dev --local --file=./schema.sql
```

## Upload to remote d1:

NOTE: Remove transaction if you dumped it straight from sqlite3.

```bash
npx wrangler d1 execute massmurdercanada-dev --remote --file=./schema.sql
```

## Run a remote SQL quiery against d1:

```bash
npx wrangler d1 execute massmurdercanada-dev --remote --command="SELECT * FROM news_stories"
```
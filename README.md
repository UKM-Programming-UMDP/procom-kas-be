# Procom Kas Backend

### Requirements

- Golang v1.22.0

### Development Setup

---

### Clone This Project

### Install

```bash
go mod tidy
```

### Database

This project using PostgreSQL as database. There is two options for you to host it:
1. Using docker, simply host it by download an official postgres image, and run it. But the mines is it consume a lot of RAM (you need at least 16 GB of RAM)
2. Using WSL2, this is quite complicate for the setup, but it's cost a lot less ram (you can run in on laptop with 8 GB RAM). Google for tutorial, or contact backend developer.

#### Step by Step Using WSL2
1. Install Ubuntu to your wsl
```bash
wsl --install
wsl --set-default Ubuntu
```
Note: after install, you need to input username and password, that's for root user of your linux

2. Install postgresql
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

3. Start postgresql
```bash
sudo systemctl status postgresql
sudo systemctl start postgresql
# make sure it's active when check the status
```

4. Create postgres user
```bash
sudo -u postgres psql
ALTER USER postgres PASSWORD 'secret';
CREATE DATABASE procom_kas;
```
Note: replace 'secret' with your preferences, this will be put to your .env

5. Update Your .env files with username and password that you create, since the username is 'postgres', you only need update the password

### Environment variables

- Copy sample.env and rename it to .env
- Update the values with your values

# Development

### Create a new branch

To solve an issue from Backlog, always create a new branch to the related repository. The format is `{type}/{your-issue}`. Example:
`feature/add-middleware`, `fix/cannot-create-user`

---

### Commit changes

Commit at the smallest meaningful changes possible. Give clear and concise commit message:

1.  What has been created/modified
2.  What will be affected

This will help you and other team member to debug later when there’s a failure. Example commit message:

> add: middleware for auth user

> edit: create new user response

> fix: cannot create new user

---

### Run on your local

After making modifications, make sure you can run it on your local. Pay attention to warnings and errors and fix that before continue to next step.

```bash
go run main.go
```
Or you can use `make dev` command if you install makefile software.

---

### Push to your branch

Push your working modification to **your own branch**.

---

### Make a Pull Request

- Make a pull request into development branch
- Sync your branch by running `git pull origin development`
- Make sure your code still working and remove any conflict
- Create the PR

Give clear PR title message to **each** branch, let the other developer know what you’re trying to merge. Example PR title:

> Feature/Route to create a financial request

> Fix/Cannot create a new user

---

### Main Branches

**Development**

- Every development and testing must be on this branch.
- This branch will using development database.

**Staging**

- Every issues from development, will be cherry-picked into this branch along with versioning.
- THis branch will using production database for final testing.

**Master**

- Full release into production.

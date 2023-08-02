Search
=========

This is a guide to help you set up and run the Search project on a Linux server, specifically Debian or Ubuntu.

### Step 1: Clone the repository

First, navigate to the directory where you want to clone the repository. Then run on terminal:


```bash
git clone https://github.com/fadelpamungkas/fulltext-search.git
cd search
```

### Step 2: Build the project (unnecessary unless you need to change the config)

Install the required Go dependencies:

```bash
go get
```

Build the project:

```bash
go build
```

### Step 3: Set up MeiliSearch (if not installed)

To add the MeiliSearch repository and install MeiliSearch, run the following commands:

```bash
echo "deb [trusted=yes] https://apt.fury.io/meilisearch/ /" | sudo tee /etc/apt/sources.list.d/fury.list
sudo apt update && sudo apt install meilisearch
```

Start MeiliSearch:

- Launch meilisearch:

```bash
meilisearch
```

- Launch meilisearch on background (recommended):

bash

```bash
sudo systemctl enable meilisearch
sudo systemctl start meilisearch
```

MeiliSearch should now be running at http://127.0.0.1:7700.

More information to run meilisearch on production:
https://www.meilisearch.com/docs/learn/cookbooks/running_production


### Step 4: Set up PostgreSQL (if not installed)

Install PostgreSQL

```bash
sudo apt-get install postgresql postgresql-contrib
```

Create a new PostgreSQL user and database:

```bash
sudo -u postgres createuser --interactive --pwprompt
sudo -u postgres createdb -O <username> job
```

Replace `<username>` with the PostgreSQL user you just created.

Create the necessary tables by running the following command:

```bash
psql -U <username> -d job -f job.sql
```

Replace `<username>` with the PostgreSQL user you just created.


### Step 5: Run the application

Finally, run the binary:

```bash
./main
```

The web server should now be running at http://127.0.0.1:8000.

## Default configuration

The default configuration for MeiliSearch is:

```md
Host:   "http://127.0.0.1:7700",
APIKey: "LQ2DYe1qqAWh9Kvihp7Lar86TTf6GLXuc-fpWbhVzGA",
Index:  "job",
```

The default configuration for PostgreSQL is:

```md
Username: "user_new",
Password: "userpass",
Hostname: "localhost",
Port:     "5432",
DBName:   "job",
```

To change the default configs, go to /app/main.go file

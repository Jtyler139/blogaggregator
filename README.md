<h1>Gator</h1>
<p>Blog Aggregator</p>

<h2>Dependencies</h2>
<p>Blog aggregatror requires the use of:</p>
<li>Postgres</li>
<li>Go</li>

<h2>How to install</h2>
Download the project using command go install https://github.com/Jtyler139/blogaggregator@latest

Create a config file in your home directory called ~/.gatorconfig.json
Copy and paste this inside file
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "username_goes_here"
}

make sure to change current user name to your username

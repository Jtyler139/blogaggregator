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

<h2>List of commands</h2>
<p>to run the blogaggregator app simply type blogaggregator followed by one of the command names</p>
<ol>register <i>name</i> creates a new user and logs them in</ol>
<ol>login: <i>name</i>: allows you to log in to specified user</ol>
<ol>reset: resets blogaggregator app</ol>
<ol>users: lists all users</ol>
<ol>addfeed <i>name, url</i>: allows you to add a new with the specified name and url</ol>
<ol>feeds: list current user's feeds</ol>
<ol>follow: <i>feed_url</i>: allows current user to follow an existing feed with specified url</ol>
<ol>following: lists current users follows</ol>
<ol>unfollow <i>feed_url</i>: unfollows feed with specifed url</ol>
<ol>agg <i>time_between_reqs</i>: collects data from all feeds every specifed time duration</ol>
<ol>browse <i>entry limit</i>: allows user to browse specified number of posts</ol>

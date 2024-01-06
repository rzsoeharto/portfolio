Welcome to the source code for my portfolio/blog project.

This project's tech stack is ReactJS, Gin-Gonic (Golang), and Postgres as the database.
Alongside React I also use the CSS framework Tailwind to accelerate styling.

The authentication in this project is made from scratch, logging in requires a unique username and password (although this is only meant for personal use).
While the persistent login mechanism employs JWT as a way of re-authentication.
Refreshed tokens are blacklisted; blacklisted tokens are kept track with a simple database solution.

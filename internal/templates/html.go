package templates

// This is just a collection of HTML templates.
//
// They could be stored in their own separate asset files,
// but that feels like more work for now, and I'm pretty OK
// at rawdogging the HTML.
//
// I've done simple stuff like this enough times.

const Home = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>goppert's ohq</title>
		<script
			src="https://unpkg.com/htmx.org@2.0.3"
			integrity="sha384-0895/pl2MU10Hqc6jd4RvrthNlDiE9U1tWmX7WRESftEDRosgxNsQG/Ze9YMRzHq"
			crossorigin="anonymous"></script>
		<style>
			body {
				max-width: 800px;
				margin: 0 auto;
				padding: 2vh;
			}

			header {
				margin: 0 0 3vh 0;
				padding: 0;
			}

			header p {
				font-weight: bold;
				font-size: 2rem;
				margin: 0;
				padding: 0;
			}

			#queue-disp,
			#queue-add {
				margin-bottom: 2rem;
			}

			footer {
				display: flex;
				justify-content: flex-start;
			}

			footer div {
				margin: 0 0 0 0.5rem;
			}

			footer a {
				text-decoration: none;
				color: black;
			}
		</style>
	</head>
	<body>
		<div id="top">
			<header>
				<p>ohq</p>
			</header>

			<div id="queue-disp">
				<ul>
					{{range .Users}}
					<li>{{.ID}}</li>
					{{end}}
				</ul>
			</div>

			<div id="queue-add">
				<form
					hx-post="/add"
					hx-target="#queue-disp"
					hx-swap="innerHTML"
					hx-trigger="submit">
					<input
						type="text"
						name="qid"
						id="qid"
						autocomplete="off" />
				</form>
			</div>
		</div>
		<div id="bottom">
			<footer>
				<div>© 2024 michael goppert</div>
				<div>|</div>
				<div>
					<a href="https://www.github.com/migopp/ohq">github.com/migopp/ohq</a>
				</div>
			</footer>
		</div>
	</body>
</html>
`

const QueueDisplay = `
<ul>
	{{range .Users}}
	<li>{{.ID}}</li>
	{{end}}
</ul>
`

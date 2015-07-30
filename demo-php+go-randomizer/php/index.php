<html>
	<head><title>The Randomizer</title></head>
	<body style="font-family: helvetica;">
		<div style="margin: auto; text-align: center;">
			<br><br>
			<h1>Randomizer App</h1>
			<br><br>
			<?php

			// Lookup API link location and consume all JSON found there
			$env = explode("//", getenv('RAND_URI'));
			$link = "http://" . $env[1];
			$ctx = stream_context_create(array( 
			    'http' => array( 
			        'timeout' => 1 
			        )
			    )
			);
			$json = file_get_contents($link, 0, $ctx);
			$json_data = json_decode($json, true);

			// Determine if API data was found.
			// If not, display message notifying user to connect to the service.
			if ($json_data["API_FQN"] != "") {
				echo "<p>The job you have reached is:</p>";
				echo "<h3>" . getenv('CNTM_JOB_FQN') . "<h3>";
				echo "<h2>Instance: " . getenv('CNTM_INSTANCE_UUID') . "</h2>";
				echo "<br><br>";
				echo "<p>The API consumed by the job is:</p>";
				echo "<h3>" . $json_data["API_FQN"] . "<h3>";
				echo "<h2>Instance: " . $json_data["API_Instance"] . "</h2>";
				echo "<br><br>";
				echo "<p>The API provided a randomly generated number (0-100) of:</p>";
				echo "<h1>" . $json_data["Number"] . "</h1>";
			} else {
				echo "<p>The randomizer API could not be found. Please ensure service is present.</p>";
			}

			?>
			<button onclick="reloadPage()">Reload page</button>
			<script>function reloadPage() { location.reload(); }</script>
		</div>
	</body>
</html>
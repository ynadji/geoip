<?php
  function geoip($ip) {
    $url = "http://geo:1234?ip=$ip";
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true); // hide trailing boolean
    $response = curl_exec($curl);
    curl_close($curl);
    echo $response;
  }
?>

<html>
  <head>
    <title>GIP GIP</title>
  </head>
  <body>
    <form action="index.php" method="get">
      <input id="ip" type="text" name="ip" value="1.2.3.4"></input>
      <input id="submit" type="submit" value="FINDIT" />
    </form>
    <div id="geo">Location for <?php echo $_GET["ip"]; ?> is: <?php geoip($_GET["ip"]); ?></div>
  </body>
</html>

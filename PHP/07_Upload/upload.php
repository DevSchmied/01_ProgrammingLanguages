<?php
    if(!empty($_FILES)) {
        for($i = 0; $i < count($_FILES['userbild']['name']); $i++) {
            if (
                $_FILES['userbild']['error'][$i] == 0 &&
                strpos($_FILES['userbild']['type'][$i], "image") == 0 &&
                $_FILES['userbild']['size'][$i] > 0 &&
                $_FILES['userbild']['size'][$i] < 15000000 &&
                move_uploaded_file(
                    $_FILES['userbild']['tmp_name'][$i],
                    'img/' . $_FILES['userbild']['name'][$i]
                )
            ) { ?>
                <img width="150" src="img/<?= $_FILES['userbild']['name'][$i] ?>"/>
<?php        }
        }
    }
?>

<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dateiupload</title>
</head>
    <body>
        <form action="upload.php" method="post" enctype="multipart/form-data">
            <input type="file" name="userbild[]" multiple="multiple"/>
            <input type="submit" value="hochladen"/>
        </form>
    </body>
</html>
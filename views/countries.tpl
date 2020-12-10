<!DOCTYPE html>

<html>

<head>
    <title>Countries</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>

<body>
    <h1>Countries 🎉🎉🎉</h1>
    <h3>
        <a href="/insert" onclick="return confirm('Are you sure you want to insert ?');">
        insert countries
      </a>
    </h3>
    <table border="1">
        <thead>
            <th>ID</th>
            <th>Code</th>
            <th>Name</th>
        </thead>
        <tbody>
            {{ range $val:=.listCountries }}
            <tr>
                <td>{{ $val.Id }}</td>
                <td>{{ $val.CountryCode }}</td>
                <td>{{ $val.CountryName }}</td>
            </tr>
            {{ end }}
        </tbody>
    </table>
</body>

</html>
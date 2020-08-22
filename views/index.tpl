<!DOCTYPE html>

<html>
  <head>
    <title>GormSSP</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>

    <!-- datatables inport-->
    <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.10.21/css/jquery.dataTables.css">
    <script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.10.21/js/jquery.dataTables.js"></script>
  </head>

  <body>
    <header>
      <h1>Welcome to GormSSP example</h1>
      <p>This is a example of work of GormSSP</p>
    </header>
    
    <div>
      <table id="table_id" class="display">
        <thead>
          <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Email</th>
              <th>Age</th>
          </tr>
        </thead>
      </table>
    </div>

  </body>
  <script>
  $(document).ready(function() {
    $('#table_id').DataTable( {
        processing: true,
        serverSide: true,
        ajax: "/simple/paginate",
        pageLength: 5,
        columns: [
          { data: "id" },
          { data: "name" },
          { data: "email" },
          { data: "age" },
        ]
    } );
  } );
  </script>
</html>
